package model

import (
	"testing"

	"domain_service_study/internal/domain/valueobject"
)

func TestNewQuestionID(t *testing.T) {
	id, err := NewQuestionID("overview.summary")
	if err != nil {
		t.Fatalf("NewQuestionID() error = %v", err)
	}
	if id.String() != "overview.summary" {
		t.Fatalf("NewQuestionID() = %q, want %q", id.String(), "overview.summary")
	}
}

func TestAnswersCopyInputMap(t *testing.T) {
	id := MustQuestionID("overview.summary")
	raw := map[QuestionID]valueobject.Answer{
		id: valueobject.NewAnswer("元の値"),
	}
	answers := NewAnswers(raw)
	raw[id] = valueobject.NewAnswer("書き換え")

	got, ok := answers.Get(id)
	if !ok {
		t.Fatal("answer missing")
	}
	if got.String() != "元の値" {
		t.Fatalf("Get() = %q, want %q", got.String(), "元の値")
	}
}
