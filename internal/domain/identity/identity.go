package identity

import "strings"

// Identity はドメインモデルの Brand 型。
// 型パラメータ Brand によって、同じ文字列でも別の身元として区別する。
// たとえば Identity[QuestionBrand] を別 Brand の Identity へ代入することはできない。
type Identity[Brand any] struct {
	value string
}

// NewIdentity は空でない文字列から Identity を生成する。
// 空白のみの値は EmptyIdentityError になる。
func NewIdentity[Brand any](value string) (Identity[Brand], error) {
	if strings.TrimSpace(value) == "" {
		return Identity[Brand]{}, &EmptyIdentityError{}
	}
	return Identity[Brand]{value: value}, nil
}

// MustIdentity は静的なカタログ定義向け。空文字はプログラミングエラーとして扱う。
func MustIdentity[Brand any](value string) Identity[Brand] {
	id, err := NewIdentity[Brand](value)
	if err != nil {
		panic(err)
	}
	return id
}

// String は身元の文字列表現を返す。
func (id Identity[Brand]) String() string { return id.value }

// IsZero は未設定の Identity かどうかを返す。
func (id Identity[Brand]) IsZero() bool { return id.value == "" }

// EmptyIdentityError は空の Identity を作ろうとしたときに返す。
type EmptyIdentityError struct{}

func (e *EmptyIdentityError) Error() string {
	return "identity must not be empty"
}
