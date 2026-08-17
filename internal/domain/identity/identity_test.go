package identity

import (
	"errors"
	"testing"
)

func TestIdentityRejectsEmpty(t *testing.T) {
	type testBrand struct{}
	_, err := NewIdentity[testBrand](" ")
	var empty *EmptyIdentityError
	if !errors.As(err, &empty) {
		t.Fatalf("NewIdentity() error = %v, want EmptyIdentityError", err)
	}
}

func TestIdentityBrandsAreDistinct(t *testing.T) {
	type brandA struct{}
	type brandB struct{}
	a := MustIdentity[brandA]("same-raw-value")
	b := MustIdentity[brandB]("same-raw-value")

	if a.String() != b.String() {
		t.Fatal("raw values should be equal")
	}

	// Brand が違うため、次はコンパイルできない。
	// var _ Identity[brandA] = b
	_ = b
}
