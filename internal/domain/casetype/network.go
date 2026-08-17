package casetype

import (
	"domain_service_study/internal/domain/model"
	"domain_service_study/internal/domain/valueobject"
)

// NetworkModule はネットワークトラブル固有の質問と、回答に依存する分岐質問を返す。
type NetworkModule struct{}

func (NetworkModule) BaseDefinition() model.ReportDefinition {
	return mustDefinition(
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

func (NetworkModule) AnswerDependentDefinition(answers model.Answers) model.ReportDefinition {
	value, ok := answers.Get(QuestionNetworkOtherUsersAffected)
	if !ok || value.IsBlank() {
		return model.ReportDefinition{}
	}
	if !value.IsYes() {
		return model.ReportDefinition{}
	}

	return mustDefinition(
		model.NewQuestionDefinition(
			QuestionNetworkAffectedDeviceCount,
			valueobject.SectionOverview,
			valueobject.MustPrompt("影響を受けている端末は何台ですか？"),
			true,
		),
	)
}
