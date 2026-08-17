package application

import (
	"domain_service_study/internal/domain/casetype"
	"domain_service_study/internal/domain/domainservice"
	"domain_service_study/internal/domain/model"
)

// CreateTroubleReportRequest はアプリケーション層の入力DTO。
// Domain Service にはこの型を渡さず、ドメインの型へ変換してから渡す。
type CreateTroubleReportRequest struct {
	TroubleType string
	ImpactLevel string
	Answers     map[string]string
}

// CreateTroubleReportUseCase は IT トラブル報告書の作成ユースケース。
type CreateTroubleReportUseCase struct {
	service domainservice.TroubleReportService
}

// NewCreateTroubleReportUseCase は CreateTroubleReportUseCase を生成する。
func NewCreateTroubleReportUseCase(service domainservice.TroubleReportService) CreateTroubleReportUseCase {
	return CreateTroubleReportUseCase{service: service}
}

// Execute はリクエストをドメイン型へ変換し、報告書を生成する。
// トラブル種類・影響度が未知の場合や必須回答が不足している場合はエラーを返す。
func (u CreateTroubleReportUseCase) Execute(req CreateTroubleReportRequest) (*model.TroubleReport, error) {
	troubleType, err := casetype.ParseTroubleType(req.TroubleType)
	if err != nil {
		return nil, err
	}
	impactLevel, err := casetype.ParseImpactLevel(req.ImpactLevel)
	if err != nil {
		return nil, err
	}
	answers, err := model.NewAnswersFromStrings(req.Answers)
	if err != nil {
		return nil, err
	}
	return u.service.NewTroubleReport(troubleType, impactLevel, answers)
}
