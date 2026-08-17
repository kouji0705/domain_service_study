package casetype

import (
	"domain_service_study/internal/domain/model"
	"domain_service_study/internal/domain/valueobject"
)

// PCModule はPCトラブル固有の質問と、回答に依存する分岐質問を返す。
type PCModule struct{}

// NewBaseFormSchema は PC トラブル固有の基本質問項目定義を生成する。
func (PCModule) NewBaseFormSchema() model.FormSchema {
	return mustFormSchema(
		model.NewQuestionDefinition(
			QuestionPCPowerOn,
			valueobject.SectionSituation,
			valueobject.MustPrompt("PCの電源は入りますか？"),
			true,
		),
	)
}

// NewAnswerDependentFormSchema は回答内容に応じた分岐質問項目定義を生成する。
// 前提となる回答が未入力または分岐条件を満たさない場合は空の定義を返す。
func (PCModule) NewAnswerDependentFormSchema(answers model.Answers) model.FormSchema {
	value, ok := answers.Get(QuestionPCPowerOn)
	if !ok || value.IsBlank() {
		return model.FormSchema{}
	}

	switch {
	case value.IsNo():
		return mustFormSchema(
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
		return mustFormSchema(
			model.NewQuestionDefinition(
				QuestionPCScreenVisible,
				valueobject.SectionSituation,
				valueobject.MustPrompt("画面は表示されていますか？"),
				true,
			),
		)
	default:
		return model.FormSchema{}
	}
}
