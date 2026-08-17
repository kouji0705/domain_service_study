package casetype

import (
	"domain_service_study/internal/domain/model"
	"domain_service_study/internal/domain/valueobject"
)

// PCModule はPCトラブル固有の質問と、回答に依存する分岐質問を返す。
type PCModule struct{}

func (PCModule) BaseDefinition() model.ReportDefinition {
	return mustDefinition(
		model.NewQuestionDefinition(
			QuestionPCPowerOn,
			valueobject.SectionSituation,
			valueobject.MustPrompt("PCの電源は入りますか？"),
			true,
		),
	)
}

func (PCModule) AnswerDependentDefinition(answers model.Answers) model.ReportDefinition {
	value, ok := answers.Get(QuestionPCPowerOn)
	if !ok || value.IsBlank() {
		return model.ReportDefinition{}
	}

	switch {
	case value.IsNo():
		return mustDefinition(
			model.NewQuestionDefinition(
				QuestionPCPowerLight,
				valueobject.SectionAssessment,
				valueobject.MustPrompt("電源ランプは点灯していますか？"),
				true,
			),
			model.NewQuestionDefinition(
				QuestionPCACAdapterConnected,
				valueobject.SectionBackground,
				valueobject.MustPrompt("ACアダプターは接続されていますか？"),
				true,
			),
		)
	case value.IsYes():
		return mustDefinition(
			model.NewQuestionDefinition(
				QuestionPCScreenVisible,
				valueobject.SectionSituation,
				valueobject.MustPrompt("画面は表示されていますか？"),
				true,
			),
		)
	default:
		return model.ReportDefinition{}
	}
}
