package model

import "domain_service_study/internal/domain/valueobject"

// AnswerItem は報告書に確定した1件の回答。
type AnswerItem struct {
	questionID QuestionID
	prompt     valueobject.Prompt
	value      valueobject.Answer
}

// NewAnswerItem は報告書に確定した1件の回答を生成する。
func NewAnswerItem(questionID QuestionID, prompt valueobject.Prompt, value valueobject.Answer) AnswerItem {
	return AnswerItem{
		questionID: questionID,
		prompt:     prompt,
		value:      value,
	}
}

// QuestionID は回答対象の質問 ID を返す。
func (i AnswerItem) QuestionID() QuestionID { return i.questionID }

// Prompt は質問文を返す。
func (i AnswerItem) Prompt() valueobject.Prompt { return i.prompt }

// Value は回答内容を返す。
func (i AnswerItem) Value() valueobject.Answer { return i.value }

// SectionAnswers は1つのセクションに属する回答の集まり。
type SectionAnswers struct {
	items []AnswerItem
}

// NewSectionAnswers は1つのセクションに属する回答の集まりを生成する。
func NewSectionAnswers(items []AnswerItem) SectionAnswers {
	copied := make([]AnswerItem, len(items))
	copy(copied, items)
	return SectionAnswers{items: copied}
}

// Items は含まれる回答のコピーを返す。
func (s SectionAnswers) Items() []AnswerItem {
	copied := make([]AnswerItem, len(s.items))
	copy(copied, s.items)
	return copied
}

// Value は指定した質問 ID の回答を返す。
func (s SectionAnswers) Value(id QuestionID) (valueobject.Answer, bool) {
	for _, item := range s.items {
		if item.questionID == id {
			return item.value, true
		}
	}
	return valueobject.Answer{}, false
}

// Clone は SectionAnswers のコピーを返す。
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

// NewSBAR は SBAR 構造を生成する。
func NewSBAR(situation, background, assessment, recommendation SectionAnswers) SBAR {
	return SBAR{
		situation:      situation.Clone(),
		background:     background.Clone(),
		assessment:     assessment.Clone(),
		recommendation: recommendation.Clone(),
	}
}

// Situation は Situation セクションの回答を返す。
func (s SBAR) Situation() SectionAnswers { return s.situation.Clone() }

// Background は Background セクションの回答を返す。
func (s SBAR) Background() SectionAnswers { return s.background.Clone() }

// Assessment は Assessment セクションの回答を返す。
func (s SBAR) Assessment() SectionAnswers { return s.assessment.Clone() }

// Recommendation は Recommendation セクションの回答を返す。
func (s SBAR) Recommendation() SectionAnswers { return s.recommendation.Clone() }

// Clone は SBAR のコピーを返す。
func (s SBAR) Clone() SBAR {
	return NewSBAR(s.situation, s.background, s.assessment, s.recommendation)
}
