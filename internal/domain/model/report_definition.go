package model

import "fmt"

type DuplicateQuestionIDError struct {
	ID QuestionID
}

func (e *DuplicateQuestionIDError) Error() string {
	return fmt.Sprintf("duplicate question id: %s", e.ID)
}

type QuestionNotFoundError struct {
	ID QuestionID
}

func (e *QuestionNotFoundError) Error() string {
	return fmt.Sprintf("question not found: %s", e.ID)
}

// ReportDefinition は今回表示する質問と Validation ルールの集合。
type ReportDefinition struct {
	questions []QuestionDefinition
}

func NewReportDefinition(questions ...QuestionDefinition) (ReportDefinition, error) {
	seen := make(map[QuestionID]struct{}, len(questions))
	copied := make([]QuestionDefinition, 0, len(questions))
	for _, question := range questions {
		if _, exists := seen[question.id]; exists {
			return ReportDefinition{}, &DuplicateQuestionIDError{ID: question.id}
		}
		seen[question.id] = struct{}{}
		copied = append(copied, question)
	}
	return ReportDefinition{questions: copied}, nil
}

func (d ReportDefinition) Questions() []QuestionDefinition {
	copied := make([]QuestionDefinition, len(d.questions))
	copy(copied, d.questions)
	return copied
}

func (d ReportDefinition) Contains(id QuestionID) bool {
	_, ok := d.Question(id)
	return ok
}

func (d ReportDefinition) Question(id QuestionID) (QuestionDefinition, bool) {
	for _, question := range d.questions {
		if question.id == id {
			return question, true
		}
	}
	return QuestionDefinition{}, false
}

// Combine は元の定義を変更せず、新しい ReportDefinition を返す。
// 質問IDが重複した場合は黙って上書きせず、エラーにする。
func (d ReportDefinition) Combine(other ReportDefinition) (ReportDefinition, error) {
	questions := make([]QuestionDefinition, 0, len(d.questions)+len(other.questions))
	questions = append(questions, d.questions...)
	questions = append(questions, other.questions...)
	return NewReportDefinition(questions...)
}

// OverrideRequired は既存質問の必須条件を上書きする。
// 同じ質問IDを重複追加する代わりに、合成後の必須条件を変えられるようにする。
func (d ReportDefinition) OverrideRequired(id QuestionID, required bool) (ReportDefinition, error) {
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
		return ReportDefinition{}, &QuestionNotFoundError{ID: id}
	}
	return ReportDefinition{questions: questions}, nil
}

func (d ReportDefinition) Validate(answers Answers) error {
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
