package casetype

import (
	"domain_service_study/internal/domain/model"
	"domain_service_study/internal/domain/valueobject"
)

// CommonModule はすべての報告書に共通する質問とルールを返す。
type CommonModule struct{}

// NewFormSchema はすべての報告書に共通する質問項目定義を生成する。
func (CommonModule) NewFormSchema() model.FormSchema {
	return mustFormSchema(
		model.NewQuestionDefinition(
			QuestionOverviewSummary,
			valueobject.SectionOverview,
			valueobject.MustPrompt("どのような問題が発生していますか？"),
			true,
		),
		model.NewQuestionDefinition(
			QuestionSituationOccurredAt,
			valueobject.SectionSituation,
			valueobject.MustPrompt("いつ問題が発生しましたか？"),
			true,
		),
		model.NewQuestionDefinition(
			QuestionBackgroundBeforeOccurrence,
			valueobject.SectionBackground,
			valueobject.MustPrompt("問題が発生する前に何をしていましたか？"),
			true,
		),
		model.NewQuestionDefinition(
			QuestionAssessmentPossibleCause,
			valueobject.SectionAssessment,
			valueobject.MustPrompt("原因として考えられることはありますか？"),
			false,
		),
		model.NewQuestionDefinition(
			QuestionRecommendationRequestedAction,
			valueobject.SectionRecommendation,
			valueobject.MustPrompt("どのような対応を希望しますか？"),
			false,
		),
		model.NewQuestionDefinition(
			QuestionOtherNotes,
			valueobject.SectionOther,
			valueobject.MustPrompt("そのほかに伝えたいことはありますか？"),
			false,
		),
	)
}
