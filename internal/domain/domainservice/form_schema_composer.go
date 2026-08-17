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

// FormSchemaComposer は共通・種類別・影響度別・回答依存の質問項目定義を合成する。
type FormSchemaComposer struct {
	common  casetype.CommonModule
	pc      casetype.PCModule
	network casetype.NetworkModule
	impact  casetype.ImpactModule
}

// NewFormSchemaComposer は FormSchemaComposer を生成する。
func NewFormSchemaComposer() FormSchemaComposer {
	return FormSchemaComposer{}
}

// NewFormSchema は共通・種類別・影響度別・回答依存の質問項目定義を合成して生成する。
// トラブル種類または影響度が未知の場合はエラーを返す。
func (c FormSchemaComposer) NewFormSchema(
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

	typeModule, err := c.typeModule(troubleType)
	if err != nil {
		return model.FormSchema{}, err
	}

	schema := c.common.NewFormSchema()
	schema, err = schema.Combine(typeModule.NewBaseFormSchema())
	if err != nil {
		return model.FormSchema{}, err
	}
	schema, err = schema.Combine(c.impact.NewFormSchema(impactLevel))
	if err != nil {
		return model.FormSchema{}, err
	}
	schema, err = schema.Combine(typeModule.NewAnswerDependentFormSchema(answers))
	if err != nil {
		return model.FormSchema{}, err
	}
	return c.impact.ApplyRequiredOverrides(schema, impactLevel)
}

func (c FormSchemaComposer) typeModule(troubleType model.TroubleType) (typeFormSchemaModule, error) {
	switch {
	case casetype.IsPC(troubleType):
		return c.pc, nil
	case casetype.IsNetwork(troubleType):
		return c.network, nil
	default:
		return nil, &casetype.UnknownTroubleTypeError{Value: troubleType.String()}
	}
}
