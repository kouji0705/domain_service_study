package casetype

import (
	"fmt"

	"domain_service_study/internal/domain/model"
)

var (
	// ImpactLevelIndividual は個人のみに影響する段階。
	ImpactLevelIndividual = model.MustImpactLevel("individual")
	// ImpactLevelTeam はチームに影響する段階。
	ImpactLevelTeam = model.MustImpactLevel("team")
	// ImpactLevelCompanyWide は全社に影響する段階。
	ImpactLevelCompanyWide = model.MustImpactLevel("company_wide")
)

// ParseImpactLevel は文字列を既知の ImpactLevel に変換する。
func ParseImpactLevel(value string) (model.ImpactLevel, error) {
	for _, candidate := range []model.ImpactLevel{ImpactLevelIndividual, ImpactLevelTeam, ImpactLevelCompanyWide} {
		if value == candidate.String() {
			return candidate, nil
		}
	}
	return model.ImpactLevel{}, &UnknownImpactLevelError{Value: value}
}

// IsKnownImpactLevel は影響度がこのサンプルで定義済みかどうかを返す。
func IsKnownImpactLevel(l model.ImpactLevel) bool {
	return l == ImpactLevelIndividual || l == ImpactLevelTeam || l == ImpactLevelCompanyWide
}

// IsIndividual は個人影響かどうかを返す。
func IsIndividual(l model.ImpactLevel) bool { return l == ImpactLevelIndividual }

// IsTeam はチーム影響かどうかを返す。
func IsTeam(l model.ImpactLevel) bool { return l == ImpactLevelTeam }

// IsCompanyWide は全社影響かどうかを返す。
func IsCompanyWide(l model.ImpactLevel) bool { return l == ImpactLevelCompanyWide }

// UnknownImpactLevelError は未知の影響度文字列を渡されたときに返す。
type UnknownImpactLevelError struct {
	Value string
}

func (e *UnknownImpactLevelError) Error() string {
	return fmt.Sprintf("unknown impact level: %q", e.Value)
}
