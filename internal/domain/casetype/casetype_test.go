package casetype

import (
	"errors"
	"testing"

	"domain_service_study/internal/domain/model"
)

func TestParseTroubleType(t *testing.T) {
	t.Run("既知の値から生成できる", func(t *testing.T) {
		got, err := ParseTroubleType("pc")
		if err != nil {
			t.Fatalf("ParseTroubleType() error = %v", err)
		}
		if got != TroubleTypePC {
			t.Fatalf("ParseTroubleType() = %v, want %v", got, TroubleTypePC)
		}
	})

	t.Run("未知の値は生成できない", func(t *testing.T) {
		_, err := ParseTroubleType("unknown")
		var unknown *UnknownTroubleTypeError
		if !errors.As(err, &unknown) {
			t.Fatalf("ParseTroubleType() error = %v, want UnknownTroubleTypeError", err)
		}
		if unknown.Value != "unknown" {
			t.Fatalf("UnknownTroubleTypeError.Value = %q, want %q", unknown.Value, "unknown")
		}
	})
}

func TestParseImpactLevel(t *testing.T) {
	_, err := ParseImpactLevel("unknown")
	var unknown *UnknownImpactLevelError
	if !errors.As(err, &unknown) {
		t.Fatalf("ParseImpactLevel() error = %v, want UnknownImpactLevelError", err)
	}
}

func TestKnownClassifications(t *testing.T) {
	if !IsPC(TroubleTypePC) || IsNetwork(TroubleTypePC) {
		t.Fatal("pc classification is wrong")
	}
	if !IsTeam(ImpactLevelTeam) || IsCompanyWide(ImpactLevelTeam) {
		t.Fatal("team classification is wrong")
	}
	if !IsKnownTroubleType(TroubleTypeNetwork) || IsKnownTroubleType(model.MustTroubleType("other")) {
		t.Fatal("known trouble type check is wrong")
	}
}
