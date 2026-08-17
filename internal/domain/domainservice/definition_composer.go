package domainservice

import (
	"domain_service_study/internal/domain/casetype"
	"domain_service_study/internal/domain/model"
)

// typeDefinitionModule はトラブル種類ごとの質問定義を同じ手順で合成するための差替え点。
type typeDefinitionModule interface {
	BaseDefinition() model.ReportDefinition
	AnswerDependentDefinition(answers model.Answers) model.ReportDefinition
}

// DefinitionComposer は共通・種類別・影響度別・回答依存の定義を合成する。
type DefinitionComposer struct {
	common  casetype.CommonModule
	pc      casetype.PCModule
	network casetype.NetworkModule
	impact  casetype.ImpactModule
}

func NewDefinitionComposer() DefinitionComposer {
	return DefinitionComposer{}
}

func (c DefinitionComposer) Compose(
	troubleType model.TroubleType,
	impactLevel model.ImpactLevel,
	answers model.Answers,
) (model.ReportDefinition, error) {
	if !casetype.IsKnownTroubleType(troubleType) {
		return model.ReportDefinition{}, &casetype.UnknownTroubleTypeError{Value: troubleType.String()}
	}
	if !casetype.IsKnownImpactLevel(impactLevel) {
		return model.ReportDefinition{}, &casetype.UnknownImpactLevelError{Value: impactLevel.String()}
	}

	typeModule, err := c.typeModule(troubleType)
	if err != nil {
		return model.ReportDefinition{}, err
	}

	def := c.common.Definition()
	def, err = def.Combine(typeModule.BaseDefinition())
	if err != nil {
		return model.ReportDefinition{}, err
	}
	def, err = def.Combine(c.impact.Definition(impactLevel))
	if err != nil {
		return model.ReportDefinition{}, err
	}
	def, err = def.Combine(typeModule.AnswerDependentDefinition(answers))
	if err != nil {
		return model.ReportDefinition{}, err
	}
	return c.impact.ApplyRequiredOverrides(def, impactLevel)
}

func (c DefinitionComposer) typeModule(troubleType model.TroubleType) (typeDefinitionModule, error) {
	switch {
	case casetype.IsPC(troubleType):
		return c.pc, nil
	case casetype.IsNetwork(troubleType):
		return c.network, nil
	default:
		return nil, &casetype.UnknownTroubleTypeError{Value: troubleType.String()}
	}
}
