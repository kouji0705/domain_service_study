package model

import (
	"domain_service_study/internal/domain/identity"
	"domain_service_study/internal/domain/valueobject"
)

// QuestionBrand は質問 ID の Brand。具体的な質問の一覧は casetype が持つ。
type QuestionBrand struct{}

type QuestionID = identity.Identity[QuestionBrand]

func NewQuestionID(value string) (QuestionID, error) {
	return identity.NewIdentity[QuestionBrand](value)
}

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

func NewQuestionDefinition(id QuestionID, section valueobject.Section, prompt valueobject.Prompt, required bool) QuestionDefinition {
	return QuestionDefinition{
		id:       id,
		section:  section,
		prompt:   prompt,
		required: required,
	}
}

func (q QuestionDefinition) ID() QuestionID               { return q.id }
func (q QuestionDefinition) Section() valueobject.Section { return q.section }
func (q QuestionDefinition) Prompt() valueobject.Prompt   { return q.prompt }
func (q QuestionDefinition) Required() bool               { return q.required }
