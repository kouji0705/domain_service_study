package application

import (
	"errors"
	"testing"

	"domain_service_study/internal/domain/casetype"
	"domain_service_study/internal/domain/domainservice"
)

func TestExecuteUnknownValues(t *testing.T) {
	usecase := NewCreateTroubleReportUseCase(
		domainservice.NewTroubleReportService(domainservice.NewDefinitionComposer()),
	)

	t.Run("未知のトラブル種類はエラーになる", func(t *testing.T) {
		_, err := usecase.Execute(CreateTroubleReportRequest{
			TroubleType: "unknown",
			ImpactLevel: casetype.ImpactLevelIndividual.String(),
		})
		var unknown *casetype.UnknownTroubleTypeError
		if !errors.As(err, &unknown) {
			t.Fatalf("Execute() error = %v, want UnknownTroubleTypeError", err)
		}
	})

	t.Run("未知の影響度はエラーになる", func(t *testing.T) {
		_, err := usecase.Execute(CreateTroubleReportRequest{
			TroubleType: casetype.TroubleTypePC.String(),
			ImpactLevel: "unknown",
		})
		var unknown *casetype.UnknownImpactLevelError
		if !errors.As(err, &unknown) {
			t.Fatalf("Execute() error = %v, want UnknownImpactLevelError", err)
		}
	})
}
