package model

import "fmt"

// DuplicateQuestionIDError は質問 ID が重複した FormSchema を作ろうとしたときに返す。
type DuplicateQuestionIDError struct {
	ID QuestionID
}

func (e *DuplicateQuestionIDError) Error() string {
	return fmt.Sprintf("duplicate question id: %s", e.ID)
}

// QuestionNotFoundError は存在しない質問 ID を参照したときに返す。
type QuestionNotFoundError struct {
	ID QuestionID
}

func (e *QuestionNotFoundError) Error() string {
	return fmt.Sprintf("question not found: %s", e.ID)
}

// FormSchema は今回表示する質問項目と Validation ルールの集合。
type FormSchema struct {
	questions []QuestionDefinition
}

// NewFormSchema は質問項目の定義から FormSchema を生成する。
// 質問 ID が重複している場合は DuplicateQuestionIDError を返す。
func NewFormSchema(questions ...QuestionDefinition) (FormSchema, error) {
	seen := make(map[QuestionID]struct{}, len(questions))
	copied := make([]QuestionDefinition, 0, len(questions))
	for _, question := range questions {
		if _, exists := seen[question.id]; exists {
			return FormSchema{}, &DuplicateQuestionIDError{ID: question.id}
		}
		seen[question.id] = struct{}{}
		copied = append(copied, question)
	}
	return FormSchema{questions: copied}, nil
}

// Questions は含まれる質問定義のコピーを返す。
func (d FormSchema) Questions() []QuestionDefinition {
	copied := make([]QuestionDefinition, len(d.questions))
	copy(copied, d.questions)
	return copied
}

// Contains は指定した質問 ID が定義に含まれるかどうかを返す。
func (d FormSchema) Contains(id QuestionID) bool {
	_, ok := d.Question(id)
	return ok
}

// Question は指定した質問 ID の定義を返す。存在しない場合は false を返す。
func (d FormSchema) Question(id QuestionID) (QuestionDefinition, bool) {
	for _, question := range d.questions {
		if question.id == id {
			return question, true
		}
	}
	return QuestionDefinition{}, false
}

// Combine は元の定義を変更せず、新しい FormSchema を返す。
// 質問IDが重複した場合は黙って上書きせず、エラーにする。
func (d FormSchema) Combine(other FormSchema) (FormSchema, error) {
	questions := make([]QuestionDefinition, 0, len(d.questions)+len(other.questions))
	questions = append(questions, d.questions...)
	questions = append(questions, other.questions...)
	return NewFormSchema(questions...)
}

// OverrideRequired は既存質問の必須条件を上書きする。
// 同じ質問IDを重複追加する代わりに、合成後の必須条件を変えられるようにする。
func (d FormSchema) OverrideRequired(id QuestionID, required bool) (FormSchema, error) {
	questions := make([]QuestionDefinition, len(d.questions))
	copy(questions, d.questions)

	found := false
	for i, question := range questions {
		if question.id == id {
			questions[i].required = required
			found = true
			break
		}
	}
	if !found {
		return FormSchema{}, &QuestionNotFoundError{ID: id}
	}
	return FormSchema{questions: questions}, nil
}

// Validate は必須質問に対する回答が揃っているか検証する。
// 不足がある場合は ValidationError を返す。
func (d FormSchema) Validate(answers Answers) error {
	issues := make([]ValidationIssue, 0)
	for _, question := range d.questions {
		if !question.required {
			continue
		}
		value, ok := answers.Get(question.id)
		if !ok || value.IsBlank() {
			issues = append(issues, NewValidationIssue(question.id, "required answer is missing"))
		}
	}
	if len(issues) > 0 {
		return NewValidationError(issues)
	}
	return nil
}
