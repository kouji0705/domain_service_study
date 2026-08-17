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

func (s TroubleReportService) Definition(draft model.TroubleReportDraft) (model.ReportDefinition, error) {
	return s.composer.Compose(draft.TroubleType(), draft.ImpactLevel(), draft.Answers())
}

func (s TroubleReportService) Create(draft model.TroubleReportDraft) (*model.TroubleReport, error) {
	def, err := s.composer.Compose(draft.TroubleType(), draft.ImpactLevel(), draft.Answers())
	if err != nil {
		return nil, err
	}
	if err := def.Validate(draft.Answers()); err != nil {
		return nil, err
	}
	return model.NewTroubleReport(draft.TroubleType(), draft.ImpactLevel(), def, draft.Answers()), nil
}
