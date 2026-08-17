package casetype

import (
	"domain_service_study/internal/domain/model"
	"domain_service_study/internal/domain/valueobject"
)

// ImpactModule は影響度固有の質問と、既存質問の必須条件の上書きを返す。
type ImpactModule struct{}

func (ImpactModule) Definition(level model.ImpactLevel) model.ReportDefinition {
	questions := make([]model.QuestionDefinition, 0, 2)

	if IsTeam(level) || IsCompanyWide(level) {
		questions = append(questions, model.NewQuestionDefinition(
			QuestionImpactAffectedPeople,
			valueobject.SectionOverview,
			valueobject.MustPrompt("影響を受けている人数を入力してください"),
			true,
		))
	}
	if IsCompanyWide(level) {
		questions = append(questions, model.NewQuestionDefinition(
			QuestionImpactWorkaround,
			valueobject.SectionRecommendation,
			valueobject.MustPrompt("現在利用できる代替手段を入力してください"),
			true,
		))
	}

	return mustDefinition(questions...)
}

func (ImpactModule) ApplyRequiredOverrides(def model.ReportDefinition, level model.ImpactLevel) (model.ReportDefinition, error) {
	if !IsCompanyWide(level) {
		return def, nil
	}
	// company_wide では共通質問の希望対応も必須にする。同じIDを重複追加せず、必須条件だけ上書きする。
	return def.OverrideRequired(QuestionRecommendationRequestedAction, true)
}
