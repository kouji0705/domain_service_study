package model

import (
	"errors"
	"testing"

	"domain_service_study/internal/domain/valueobject"
)

func TestFormSchemaCombineRejectsDuplicateQuestionID(t *testing.T) {
	summary := MustQuestionID("overview.summary")
	base, err := NewFormSchema(
		NewQuestionDefinition(summary, valueobject.SectionOverview, valueobject.MustPrompt("概要"), true),
	)
	if err != nil {
		t.Fatalf("NewFormSchema() error = %v", err)
	}
	other, err := NewFormSchema(
		NewQuestionDefinition(summary, valueobject.SectionOverview, valueobject.MustPrompt("別の概要"), false),
	)
	if err != nil {
		t.Fatalf("NewFormSchema() error = %v", err)
	}

	_, err = base.Combine(other)
	if err == nil {
		t.Fatal("Combine() error = nil, want duplicate question id error")
	}

	var dup *DuplicateQuestionIDError
	if !errors.As(err, &dup) {
		t.Fatalf("Combine() error type = %T, want *DuplicateQuestionIDError", err)
	}
	if dup.ID != summary {
		t.Fatalf("DuplicateQuestionIDError.ID = %s, want %s", dup.ID, summary)
	}
}

func TestFormSchemaCombineDoesNotMutateOriginal(t *testing.T) {
	summary := MustQuestionID("overview.summary")
	powerOn := MustQuestionID("pc.power_on")
	base, err := NewFormSchema(
		NewQuestionDefinition(summary, valueobject.SectionOverview, valueobject.MustPrompt("概要"), true),
	)
	if err != nil {
		t.Fatalf("NewFormSchema() error = %v", err)
	}
	other, err := NewFormSchema(
		NewQuestionDefinition(powerOn, valueobject.SectionSituation, valueobject.MustPrompt("電源"), true),
	)
	if err != nil {
		t.Fatalf("NewFormSchema() error = %v", err)
	}

	combined, err := base.Combine(other)
	if err != nil {
		t.Fatalf("Combine() error = %v", err)
	}
	if !combined.Contains(powerOn) {
		t.Fatal("combined definition should contain pc.power_on")
	}
	if base.Contains(powerOn) {
		t.Fatal("original definition was mutated")
	}
}

func TestFormSchemaOverrideRequired(t *testing.T) {
	requested := MustQuestionID("recommendation.requested_action")
	def, err := NewFormSchema(
		NewQuestionDefinition(requested, valueobject.SectionRecommendation, valueobject.MustPrompt("対応"), false),
	)
	if err != nil {
		t.Fatalf("NewFormSchema() error = %v", err)
	}

	overridden, err := def.OverrideRequired(requested, true)
	if err != nil {
		t.Fatalf("OverrideRequired() error = %v", err)
	}

	original, ok := def.Question(requested)
	if !ok || original.Required() {
		t.Fatal("original required flag should stay false")
	}
	updated, ok := overridden.Question(requested)
	if !ok || !updated.Required() {
		t.Fatal("overridden required flag should be true")
	}
}
