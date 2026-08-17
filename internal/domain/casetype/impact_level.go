package casetype

import (
	"fmt"

	"domain_service_study/internal/domain/model"
)

var (
	ImpactLevelIndividual  = model.MustImpactLevel("individual")
	ImpactLevelTeam        = model.MustImpactLevel("team")
	ImpactLevelCompanyWide = model.MustImpactLevel("company_wide")
)

func ParseImpactLevel(value string) (model.ImpactLevel, error) {
	for _, candidate := range []model.ImpactLevel{ImpactLevelIndividual, ImpactLevelTeam, ImpactLevelCompanyWide} {
		if value == candidate.String() {
			return candidate, nil
		}
	}
	return model.ImpactLevel{}, &UnknownImpactLevelError{Value: value}
}

func IsKnownImpactLevel(l model.ImpactLevel) bool {
	return l == ImpactLevelIndividual || l == ImpactLevelTeam || l == ImpactLevelCompanyWide
}

func IsIndividual(l model.ImpactLevel) bool  { return l == ImpactLevelIndividual }
func IsTeam(l model.ImpactLevel) bool        { return l == ImpactLevelTeam }
func IsCompanyWide(l model.ImpactLevel) bool { return l == ImpactLevelCompanyWide }

type UnknownImpactLevelError struct {
	Value string
}

func (e *UnknownImpactLevelError) Error() string {
	return fmt.Sprintf("unknown impact level: %q", e.Value)
}
