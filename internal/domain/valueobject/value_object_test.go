package valueobject

import (
	"errors"
	"testing"
)

func TestParseSection(t *testing.T) {
	got, err := ParseSection("assessment")
	if err != nil {
		t.Fatalf("ParseSection() error = %v", err)
	}
	if got != SectionAssessment {
		t.Fatalf("ParseSection() = %v, want %v", got, SectionAssessment)
	}

	_, err = ParseSection("unknown")
	var unknown *UnknownSectionError
	if !errors.As(err, &unknown) {
		t.Fatalf("ParseSection() error = %v, want UnknownSectionError", err)
	}
}

func TestNewPrompt(t *testing.T) {
	prompt, err := NewPrompt("どのような問題が発生していますか？")
	if err != nil {
		t.Fatalf("NewPrompt() error = %v", err)
	}
	if prompt.String() != "どのような問題が発生していますか？" {
		t.Fatalf("Prompt.String() = %q", prompt.String())
	}

	_, err = NewPrompt(" ")
	var empty *EmptyPromptError
	if !errors.As(err, &empty) {
		t.Fatalf("NewPrompt() error = %v, want EmptyPromptError", err)
	}
}

func TestAnswer(t *testing.T) {
	if !NewAnswer("yes").IsYes() {
		t.Fatal("NewAnswer(\"yes\") should be yes")
	}
	if !AnswerNo.IsNo() {
		t.Fatal("AnswerNo should be no")
	}
	if !NewAnswer(" ").IsBlank() {
		t.Fatal("whitespace answer should be blank")
	}
}
