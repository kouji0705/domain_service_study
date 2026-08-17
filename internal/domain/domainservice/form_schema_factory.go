package domainservice

import (
	"domain_service_study/internal/domain/casetype"
	"domain_service_study/internal/domain/model"
)

// typeFormSchemaModule はトラブル種類ごとの質問定義を同じ手順で合成するための差替え点。
type typeFormSchemaModule interface {
	NewBaseFormSchema() model.FormSchema
	NewAnswerDependentFormSchema(answers model.Answers) model.FormSchema
}

// FormSchemaFactory は条件に応じて FormSchema を生成する Factory。
type FormSchemaFactory struct {
	common  casetype.CommonModule
	pc      casetype.PCModule
	network casetype.NetworkModule
	impact  casetype.ImpactModule
}

// NewFormSchemaFactory は FormSchemaFactory を生成する。
func NewFormSchemaFactory() FormSchemaFactory {
	return FormSchemaFactory{}
}

// NewFormSchema は共通・種類別・影響度別・回答依存の質問項目定義を合成して生成する。
// トラブル種類または影響度が未知の場合はエラーを返す。
func (f FormSchemaFactory) NewFormSchema(
	troubleType model.TroubleType,
	impactLevel model.ImpactLevel,
	answers model.Answers,
) (model.FormSchema, error) {
	if !casetype.IsKnownTroubleType(troubleType) {
		return model.FormSchema{}, &casetype.UnknownTroubleTypeError{Value: troubleType.String()}
	}
	if !casetype.IsKnownImpactLevel(impactLevel) {
		return model.FormSchema{}, &casetype.UnknownImpactLevelError{Value: impactLevel.String()}
	}

	typeModule, err := f.typeModule(troubleType)
	if err != nil {
		return model.FormSchema{}, err
	}

	schema := f.common.NewFormSchema()
	schema, err = schema.Combine(typeModule.NewBaseFormSchema())
	if err != nil {
		return model.FormSchema{}, err
	}
	schema, err = schema.Combine(f.impact.NewFormSchema(impactLevel))
	if err != nil {
		return model.FormSchema{}, err
	}
	schema, err = schema.Combine(typeModule.NewAnswerDependentFormSchema(answers))
	if err != nil {
		return model.FormSchema{}, err
	}
	return f.impact.ApplyRequiredOverrides(schema, impactLevel)
}

func (f FormSchemaFactory) typeModule(troubleType model.TroubleType) (typeFormSchemaModule, error) {
	switch {
	case casetype.IsPC(troubleType):
		return f.pc, nil
	case casetype.IsNetwork(troubleType):
		return f.network, nil
	default:
		return nil, &casetype.UnknownTroubleTypeError{Value: troubleType.String()}
	}
}
