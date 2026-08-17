package valueobject

import "fmt"

// Section は報告書の章を表す値オブジェクト。
// 既知の章以外では生成できない。
type Section struct {
	value string
}

var (
	// SectionOverview は概要。
	SectionOverview = Section{value: "overview"}
	// SectionSituation は SBAR の Situation。
	SectionSituation = Section{value: "situation"}
	// SectionBackground は SBAR の Background。
	SectionBackground = Section{value: "background"}
	// SectionAssessment は SBAR の Assessment。
	SectionAssessment = Section{value: "assessment"}
	// SectionRecommendation は SBAR の Recommendation。
	SectionRecommendation = Section{value: "recommendation"}
	// SectionOther はその他。
	SectionOther = Section{value: "other"}
)

// ParseSection は既知の章名から Section を生成する。
// 未知の値は UnknownSectionError になる。
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

// String は章名を返す。
func (s Section) String() string { return s.value }

// IsZero は未設定の Section かどうかを返す。
func (s Section) IsZero() bool { return s.value == "" }

// IsValid は既知の章かどうかを返す。
func (s Section) IsValid() bool {
	switch s {
	case SectionOverview, SectionSituation, SectionBackground, SectionAssessment, SectionRecommendation, SectionOther:
		return true
	default:
		return false
	}
}

// UnknownSectionError は未知の章名を渡されたときに返す。
type UnknownSectionError struct {
	Value string
}

func (e *UnknownSectionError) Error() string {
	return fmt.Sprintf("unknown section: %q", e.Value)
}
