package application

import (
	"domain_service_study/internal/domain/casetype"
	"domain_service_study/internal/domain/domainservice"
	"domain_service_study/internal/domain/model"
)

// CreateTroubleReportRequest はアプリケーション層の入力DTO。
// Domain Service にはこの型を渡さず、model.TroubleReportDraft へ変換してから渡す。
type CreateTroubleReportRequest struct {
	TroubleType string
	ImpactLevel string
	Answers     map[string]string
}

type CreateTroubleReportUseCase struct {
	service domainservice.TroubleReportService
}

func NewCreateTroubleReportUseCase(service domainservice.TroubleReportService) CreateTroubleReportUseCase {
	return CreateTroubleReportUseCase{service: service}
}

func (u CreateTroubleReportUseCase) Execute(req CreateTroubleReportRequest) (*model.TroubleReport, error) {
	draft, err := ToDraft(req)
	if err != nil {
		return nil, err
	}
	return u.service.Create(draft)
}

func ToDraft(req CreateTroubleReportRequest) (model.TroubleReportDraft, error) {
	troubleType, err := casetype.ParseTroubleType(req.TroubleType)
	if err != nil {
		return model.TroubleReportDraft{}, err
	}
	impactLevel, err := casetype.ParseImpactLevel(req.ImpactLevel)
	if err != nil {
		return model.TroubleReportDraft{}, err
	}
	answers, err := model.NewAnswersFromStrings(req.Answers)
	if err != nil {
		return model.TroubleReportDraft{}, err
	}
	return model.NewTroubleReportDraft(troubleType, impactLevel, answers), nil
}
