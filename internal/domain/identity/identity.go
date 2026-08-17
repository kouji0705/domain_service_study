package identity

import "strings"

// Identity はドメインモデルの Brand 型。
// 型パラメータ Brand によって、同じ文字列でも別の身元として区別する。
// たとえば Identity[QuestionBrand] を別 Brand の Identity へ代入することはできない。
type Identity[Brand any] struct {
	value string
}

func NewIdentity[Brand any](value string) (Identity[Brand], error) {
	if strings.TrimSpace(value) == "" {
		return Identity[Brand]{}, &EmptyIdentityError{}
	}
	return Identity[Brand]{value: value}, nil
}

func MustIdentity[Brand any](value string) Identity[Brand] {
	id, err := NewIdentity[Brand](value)
	if err != nil {
		panic(err)
	}
	return id
}

func (id Identity[Brand]) String() string { return id.value }
func (id Identity[Brand]) IsZero() bool   { return id.value == "" }

type EmptyIdentityError struct{}

func (e *EmptyIdentityError) Error() string {
	return "identity must not be empty"
}
