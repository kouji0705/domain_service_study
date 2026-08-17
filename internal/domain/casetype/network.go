package casetype

import (
	"domain_service_study/internal/domain/model"
	"domain_service_study/internal/domain/valueobject"
)

// NetworkModule はネットワークトラブル固有の質問と、回答に依存する分岐質問を返す。
type NetworkModule struct{}

// NewBaseFormSchema はネットワークトラブル固有の基本質問項目定義を生成する。
func (NetworkModule) NewBaseFormSchema() model.FormSchema {
	return mustFormSchema(
		model.NewQuestionDefinition(
			QuestionNetworkConnectionType,
			valueobject.SectionBackground,
			valueobject.MustPrompt("Wi-Fiと有線のどちらを利用していますか？"),
			true,
		),
		model.NewQuestionDefinition(
			QuestionNetworkOtherUsersAffected,
			valueobject.SectionSituation,
			valueobject.MustPrompt("他の利用者にも影響がありますか？"),
			true,
		),
	)
}

// NewAnswerDependentFormSchema は回答内容に応じた分岐質問項目定義を生成する。
// 前提となる回答が未入力または分岐条件を満たさない場合は空の定義を返す。
func (NetworkModule) NewAnswerDependentFormSchema(answers model.Answers) model.FormSchema {
	value, ok := answers.Get(QuestionNetworkOtherUsersAffected)
	if !ok || value.IsBlank() {
		return model.FormSchema{}
	}
	if !value.IsYes() {
		return model.FormSchema{}
	}

	return mustFormSchema(
		model.NewQuestionDefinition(
			QuestionNetworkAffectedDeviceCount,
			valueobject.SectionOverview,
			valueobject.MustPrompt("影響を受けている端末は何台ですか？"),
			true,
		),
	)
}
