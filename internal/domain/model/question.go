package model

import (
	"domain_service_study/internal/domain/identity"
	"domain_service_study/internal/domain/valueobject"
)

// QuestionBrand は質問 ID の Brand。具体的な質問の一覧は casetype が持つ。
type QuestionBrand struct{}

// QuestionID は質問の身元。質問文が変わっても回答との対応が壊れないようにする。
type QuestionID = identity.Identity[QuestionBrand]

// NewQuestionID は空でない文字列から QuestionID を生成する。
func NewQuestionID(value string) (QuestionID, error) {
	return identity.NewIdentity[QuestionBrand](value)
}

// MustQuestionID は静的なカタログ定義向け。空文字はプログラミングエラーとして扱う。
func MustQuestionID(value string) QuestionID {
	return identity.MustIdentity[QuestionBrand](value)
}

// QuestionDefinition は報告書の1つの質問を表すドメインモデル。
// 質問文が変わっても回答との対応が壊れないよう、QuestionID で識別する。
type QuestionDefinition struct {
	id       QuestionID
	section  valueobject.Section
	prompt   valueobject.Prompt
	required bool
}

// NewQuestionDefinition は質問定義を生成する。
func NewQuestionDefinition(id QuestionID, section valueobject.Section, prompt valueobject.Prompt, required bool) QuestionDefinition {
	return QuestionDefinition{
		id:       id,
		section:  section,
		prompt:   prompt,
		required: required,
	}
}

// ID は質問の身元を返す。
func (q QuestionDefinition) ID() QuestionID { return q.id }

// Section は報告書のどの章に属すかを返す。
func (q QuestionDefinition) Section() valueobject.Section { return q.section }

// Prompt は質問文を返す。
func (q QuestionDefinition) Prompt() valueobject.Prompt { return q.prompt }

// Required は必須質問かどうかを返す。
func (q QuestionDefinition) Required() bool { return q.required }
