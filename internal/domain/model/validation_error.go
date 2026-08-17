package model

import (
	"fmt"
	"strings"
)

// ValidationIssue は1件の検証エラーを表す。
type ValidationIssue struct {
	questionID QuestionID
	reason     string
}

// NewValidationIssue は ValidationIssue を生成する。
func NewValidationIssue(questionID QuestionID, reason string) ValidationIssue {
	return ValidationIssue{
		questionID: questionID,
		reason:     reason,
	}
}

// QuestionID は問題のあった質問 ID を返す。
func (i ValidationIssue) QuestionID() QuestionID { return i.questionID }

// Reason はエラー理由を返す。
func (i ValidationIssue) Reason() string { return i.reason }

// ValidationError は質問ごとの問題を保持する。
// 呼び出し側は errors.As で判定し、不足している QuestionID を取り出せる。
type ValidationError struct {
	issues []ValidationIssue
}

// NewValidationError は ValidationError を生成する。
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

// Issues は検証エラーの一覧のコピーを返す。
func (e *ValidationError) Issues() []ValidationIssue {
	if e == nil {
		return nil
	}
	copied := make([]ValidationIssue, len(e.issues))
	copy(copied, e.issues)
	return copied
}

// QuestionIDs は問題のあった質問 ID の一覧を返す。
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
