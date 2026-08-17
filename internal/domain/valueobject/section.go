package valueobject

import "fmt"

// Section は報告書の章を表す値オブジェクト。
type Section struct {
	value string
}

var (
	SectionOverview       = Section{value: "overview"}
	SectionSituation      = Section{value: "situation"}
	SectionBackground     = Section{value: "background"}
	SectionAssessment     = Section{value: "assessment"}
	SectionRecommendation = Section{value: "recommendation"}
	SectionOther          = Section{value: "other"}
)

func ParseSection(value string) (Section, error) {
	for _, candidate := range []Section{
		SectionOverview,
		SectionSituation,
		SectionBackground,
		SectionAssessment,
		SectionRecommendation,
		SectionOther,
	} {
		if value == candidate.value {
			return candidate, nil
		}
	}
	return Section{}, &UnknownSectionError{Value: value}
}

func (s Section) String() string { return s.value }
func (s Section) IsZero() bool   { return s.value == "" }
func (s Section) IsValid() bool {
	switch s {
	case SectionOverview, SectionSituation, SectionBackground, SectionAssessment, SectionRecommendation, SectionOther:
		return true
	default:
		return false
	}
}

type UnknownSectionError struct {
	Value string
}

func (e *UnknownSectionError) Error() string {
	return fmt.Sprintf("unknown section: %q", e.Value)
}
