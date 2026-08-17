package main

import (
	"fmt"
	"log"
	"os"

	"domain_service_study/internal/application"
	"domain_service_study/internal/domain/casetype"
	"domain_service_study/internal/domain/domainservice"
	"domain_service_study/internal/domain/model"
	"domain_service_study/internal/domain/valueobject"
)

func main() {
	service := domainservice.NewTroubleReportService(domainservice.NewDefinitionComposer())
	usecase := application.NewCreateTroubleReportUseCase(service)

	report, err := usecase.Execute(application.CreateTroubleReportRequest{
		TroubleType: casetype.TroubleTypePC.String(),
		ImpactLevel: casetype.ImpactLevelTeam.String(),
		Answers: map[string]string{
			casetype.QuestionOverviewSummary.String():               "始業時にノートPCの電源が入らない",
			casetype.QuestionSituationOccurredAt.String():           "2026-08-17 09:00",
			casetype.QuestionBackgroundBeforeOccurrence.String():    "会議室へ移動して電源ボタンを押した",
			casetype.QuestionAssessmentPossibleCause.String():       "ACアダプター未接続のままバッテリーが切れた可能性がある",
			casetype.QuestionRecommendationRequestedAction.String(): "代替PCの貸出と点検を希望します",
			casetype.QuestionOtherNotes.String():                    "本日中に会議資料を編集する必要がある",
			casetype.QuestionPCPowerOn.String():                     valueobject.AnswerNo.String(),
			casetype.QuestionPCPowerLight.String():                  valueobject.AnswerNo.String(),
			casetype.QuestionPCACAdapterConnected.String():          valueobject.AnswerYes.String(),
			casetype.QuestionImpactAffectedPeople.String():          "4",
		},
	})
	if err != nil {
		log.Fatalf("failed to create trouble report: %v", err)
	}

	printReport(report)
}

func printReport(report *model.TroubleReport) {
	fmt.Fprintf(os.Stdout, "Trouble Type: %s\n", report.TroubleType())
	fmt.Fprintf(os.Stdout, "Impact Level: %s\n", report.ImpactLevel())
	printSection("Overview", report.Overview())
	sbar := report.SBAR()
	printSection("Situation", sbar.Situation())
	printSection("Background", sbar.Background())
	printSection("Assessment", sbar.Assessment())
	printSection("Recommendation", sbar.Recommendation())
	printSection("Other", report.Other())
}

func printSection(name string, answers model.SectionAnswers) {
	fmt.Fprintf(os.Stdout, "\n[%s]\n", name)
	for _, item := range answers.Items() {
		fmt.Fprintf(os.Stdout, "- %s (%s): %s\n", item.Prompt(), item.QuestionID(), item.Value())
	}
}
