package model

// TroubleReportDraft は入力途中の報告書。
// 未完成な回答を許容する。提出可能な状態かどうかは Domain Service が判定する。
type TroubleReportDraft struct {
	troubleType TroubleType
	impactLevel ImpactLevel
	answers     Answers
}

func NewTroubleReportDraft(
	troubleType TroubleType,
	impactLevel ImpactLevel,
	answers Answers,
) TroubleReportDraft {
	return TroubleReportDraft{
		troubleType: troubleType,
		impactLevel: impactLevel,
		answers:     answers.Clone(),
	}
}

func (d TroubleReportDraft) TroubleType() TroubleType { return d.troubleType }
func (d TroubleReportDraft) ImpactLevel() ImpactLevel { return d.impactLevel }
func (d TroubleReportDraft) Answers() Answers {
	return d.answers.Clone()
}
