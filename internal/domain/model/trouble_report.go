package model

import "domain_service_study/internal/domain/valueobject"

// TroubleReport は提出可能なITトラブル報告書。
// フィールドは非公開にし、生成後に外部から不正な状態へ変更できないようにする。
type TroubleReport struct {
	troubleType TroubleType
	impactLevel ImpactLevel
	overview    SectionAnswers
	sbar        SBAR
	other       SectionAnswers
}

// TroubleType はトラブル種類を返す。
func (r *TroubleReport) TroubleType() TroubleType { return r.troubleType }

// ImpactLevel は影響度を返す。
func (r *TroubleReport) ImpactLevel() ImpactLevel { return r.impactLevel }

// Overview は概要セクションの回答を返す。
func (r *TroubleReport) Overview() SectionAnswers { return r.overview.Clone() }

// SBAR は Situation / Background / Assessment / Recommendation を返す。
func (r *TroubleReport) SBAR() SBAR { return r.sbar.Clone() }

// Other はその他セクションの回答を返す。
func (r *TroubleReport) Other() SectionAnswers { return r.other.Clone() }

// NewTroubleReport は定義と回答から報告書を組み立てる。
//
// Go には兄弟パッケージだけに公開する仕組みがないため、この関数は model パッケージに置く必要がある。
// プロジェクト内ではアプリケーション層から直接呼ばず、必ず domainservice.TroubleReportService.Create を経由する。
func NewTroubleReport(
	troubleType TroubleType,
	impactLevel ImpactLevel,
	def FormSchema,
	answers Answers,
) *TroubleReport {
	overview := make([]AnswerItem, 0)
	situation := make([]AnswerItem, 0)
	background := make([]AnswerItem, 0)
	assessment := make([]AnswerItem, 0)
	recommendation := make([]AnswerItem, 0)
	other := make([]AnswerItem, 0)

	for _, question := range def.Questions() {
		item := NewAnswerItem(question.ID(), question.Prompt(), answers.Answer(question.ID()))
		switch question.Section() {
		case valueobject.SectionOverview:
			overview = append(overview, item)
		case valueobject.SectionSituation:
			situation = append(situation, item)
		case valueobject.SectionBackground:
			background = append(background, item)
		case valueobject.SectionAssessment:
			assessment = append(assessment, item)
		case valueobject.SectionRecommendation:
			recommendation = append(recommendation, item)
		case valueobject.SectionOther:
			other = append(other, item)
		}
	}

	return &TroubleReport{
		troubleType: troubleType,
		impactLevel: impactLevel,
		overview:    NewSectionAnswers(overview),
		sbar: NewSBAR(
			NewSectionAnswers(situation),
			NewSectionAnswers(background),
			NewSectionAnswers(assessment),
			NewSectionAnswers(recommendation),
		),
		other: NewSectionAnswers(other),
	}
}
