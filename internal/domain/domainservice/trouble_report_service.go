package domainservice

import (
	"domain_service_study/internal/domain/model"
)

// TroubleReportService は TroubleReport の生成窓口。
// アプリケーション層は model.NewTroubleReport を直接呼ばず、このサービスの Create を使う。
type TroubleReportService struct {
	composer DefinitionComposer
}

func NewTroubleReportService(composer DefinitionComposer) TroubleReportService {
	return TroubleReportService{composer: composer}
}

func (s TroubleReportService) Definition(
	troubleType model.TroubleType,
	impactLevel model.ImpactLevel,
	answers model.Answers,
) (model.ReportDefinition, error) {
	return s.composer.Compose(troubleType, impactLevel, answers)
}

func (s TroubleReportService) Create(
	troubleType model.TroubleType,
	impactLevel model.ImpactLevel,
	answers model.Answers,
) (*model.TroubleReport, error) {
	def, err := s.composer.Compose(troubleType, impactLevel, answers)
	if err != nil {
		return nil, err
	}
	if err := def.Validate(answers); err != nil {
		return nil, err
	}
	return model.NewTroubleReport(troubleType, impactLevel, def, answers), nil
}
