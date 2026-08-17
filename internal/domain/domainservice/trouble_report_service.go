package domainservice

import (
	"domain_service_study/internal/domain/model"
)

// TroubleReportService は TroubleReport の生成窓口。
// アプリケーション層は model.NewTroubleReport を直接呼ばず、このサービスの NewTroubleReport を使う。
type TroubleReportService struct {
	factory FormSchemaFactory
}

// NewTroubleReportService は TroubleReportService を生成する。
func NewTroubleReportService(factory FormSchemaFactory) TroubleReportService {
	return TroubleReportService{factory: factory}
}

// NewFormSchema はトラブル種類・影響度・回答に基づいて、今回有効な質問項目定義を生成する。
func (s TroubleReportService) NewFormSchema(
	troubleType model.TroubleType,
	impactLevel model.ImpactLevel,
	answers model.Answers,
) (model.FormSchema, error) {
	return s.factory.NewFormSchema(troubleType, impactLevel, answers)
}

// NewTroubleReport は質問項目定義を生成し、必須回答を検証したうえで TroubleReport を生成する。
func (s TroubleReportService) NewTroubleReport(
	troubleType model.TroubleType,
	impactLevel model.ImpactLevel,
	answers model.Answers,
) (*model.TroubleReport, error) {
	schema, err := s.factory.NewFormSchema(troubleType, impactLevel, answers)
	if err != nil {
		return nil, err
	}
	if err := schema.Validate(answers); err != nil {
		return nil, err
	}
	return model.NewTroubleReport(troubleType, impactLevel, schema, answers), nil
}
