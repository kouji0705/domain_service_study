package model

import (
	"fmt"
	"strings"
)

type ValidationIssue struct {
	questionID QuestionID
	reason     string
}

func NewValidationIssue(questionID QuestionID, reason string) ValidationIssue {
	return ValidationIssue{
		questionID: questionID,
		reason:     reason,
	}
}

func (i ValidationIssue) QuestionID() QuestionID { return i.questionID }
func (i ValidationIssue) Reason() string         { return i.reason }

// ValidationError は質問ごとの問題を保持する。
// 呼び出し側は errors.As で判定し、不足している QuestionID を取り出せる。
type ValidationError struct {
	issues []ValidationIssue
}

func NewValidationError(issues []ValidationIssue) *ValidationError {
	copied := make([]ValidationIssue, len(issues))
	copy(copied, issues)
	return &ValidationError{issues: copied}
}

func (e *ValidationError) Error() string {
	if e == nil || len(e.issues) == 0 {
		return "validation failed"
	}
	parts := make([]string, 0, len(e.issues))
	for _, issue := range e.issues {
		parts = append(parts, fmt.Sprintf("%s: %s", issue.questionID, issue.reason))
	}
	return "validation failed: " + strings.Join(parts, "; ")
}

func (e *ValidationError) Issues() []ValidationIssue {
	if e == nil {
		return nil
	}
	copied := make([]ValidationIssue, len(e.issues))
	copy(copied, e.issues)
	return copied
}

func (e *ValidationError) QuestionIDs() []QuestionID {
	if e == nil {
		return nil
	}
	ids := make([]QuestionID, 0, len(e.issues))
	for _, issue := range e.issues {
		ids = append(ids, issue.questionID)
	}
	return ids
}
