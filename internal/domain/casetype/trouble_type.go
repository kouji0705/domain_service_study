package casetype

import (
	"fmt"

	"domain_service_study/internal/domain/model"
)

var (
	// TroubleTypePC は PC トラブル。
	TroubleTypePC = model.MustTroubleType("pc")
	// TroubleTypeNetwork はネットワークトラブル。
	TroubleTypeNetwork = model.MustTroubleType("network")
)

// ParseTroubleType は文字列を既知の TroubleType に変換する。
func ParseTroubleType(value string) (model.TroubleType, error) {
	for _, candidate := range []model.TroubleType{TroubleTypePC, TroubleTypeNetwork} {
		if value == candidate.String() {
			return candidate, nil
		}
	}
	return model.TroubleType{}, &UnknownTroubleTypeError{Value: value}
}

// IsKnownTroubleType はトラブル種類がこのサンプルで定義済みかどうかを返す。
func IsKnownTroubleType(t model.TroubleType) bool {
	return t == TroubleTypePC || t == TroubleTypeNetwork
}

// IsPC は PC トラブルかどうかを返す。
func IsPC(t model.TroubleType) bool { return t == TroubleTypePC }

// IsNetwork はネットワークトラブルかどうかを返す。
func IsNetwork(t model.TroubleType) bool { return t == TroubleTypeNetwork }

// UnknownTroubleTypeError は未知のトラブル種類文字列を渡されたときに返す。
type UnknownTroubleTypeError struct {
	Value string
}

func (e *UnknownTroubleTypeError) Error() string {
	return fmt.Sprintf("unknown trouble type: %q", e.Value)
}
