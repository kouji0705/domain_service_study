package model

import "domain_service_study/internal/domain/valueobject"

// AnswerItem は報告書に確定した1件の回答。
type AnswerItem struct {
	questionID QuestionID
	prompt     valueobject.Prompt
	value      valueobject.Answer
}

func NewAnswerItem(questionID QuestionID, prompt valueobject.Prompt, value valueobject.Answer) AnswerItem {
	return AnswerItem{
		questionID: questionID,
		prompt:     prompt,
		value:      value,
	}
}

func (i AnswerItem) QuestionID() QuestionID     { return i.questionID }
func (i AnswerItem) Prompt() valueobject.Prompt { return i.prompt }
func (i AnswerItem) Value() valueobject.Answer  { return i.value }

// SectionAnswers は1つのセクションに属する回答の集まり。
type SectionAnswers struct {
	items []AnswerItem
}

func NewSectionAnswers(items []AnswerItem) SectionAnswers {
	copied := make([]AnswerItem, len(items))
	copy(copied, items)
	return SectionAnswers{items: copied}
}

func (s SectionAnswers) Items() []AnswerItem {
	copied := make([]AnswerItem, len(s.items))
	copy(copied, s.items)
	return copied
}

func (s SectionAnswers) Value(id QuestionID) (valueobject.Answer, bool) {
	for _, item := range s.items {
		if item.questionID == id {
			return item.value, true
		}
	}
	return valueobject.Answer{}, false
}

func (s SectionAnswers) Clone() SectionAnswers {
	return NewSectionAnswers(s.items)
}

// SBAR は Situation / Background / Assessment / Recommendation をまとめた構造。
type SBAR struct {
	situation      SectionAnswers
	background     SectionAnswers
	assessment     SectionAnswers
	recommendation SectionAnswers
}

func NewSBAR(situation, background, assessment, recommendation SectionAnswers) SBAR {
	return SBAR{
		situation:      situation.Clone(),
		background:     background.Clone(),
		assessment:     assessment.Clone(),
		recommendation: recommendation.Clone(),
	}
}

func (s SBAR) Situation() SectionAnswers      { return s.situation.Clone() }
func (s SBAR) Background() SectionAnswers     { return s.background.Clone() }
func (s SBAR) Assessment() SectionAnswers     { return s.assessment.Clone() }
func (s SBAR) Recommendation() SectionAnswers { return s.recommendation.Clone() }

func (s SBAR) Clone() SBAR {
	return NewSBAR(s.situation, s.background, s.assessment, s.recommendation)
}
