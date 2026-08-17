package casetype

import (
	"fmt"

	"domain_service_study/internal/domain/model"
)

var (
	TroubleTypePC      = model.MustTroubleType("pc")
	TroubleTypeNetwork = model.MustTroubleType("network")
)

func ParseTroubleType(value string) (model.TroubleType, error) {
	for _, candidate := range []model.TroubleType{TroubleTypePC, TroubleTypeNetwork} {
		if value == candidate.String() {
			return candidate, nil
		}
	}
	return model.TroubleType{}, &UnknownTroubleTypeError{Value: value}
}

func IsKnownTroubleType(t model.TroubleType) bool {
	return t == TroubleTypePC || t == TroubleTypeNetwork
}

func IsPC(t model.TroubleType) bool      { return t == TroubleTypePC }
func IsNetwork(t model.TroubleType) bool { return t == TroubleTypeNetwork }

type UnknownTroubleTypeError struct {
	Value string
}

func (e *UnknownTroubleTypeError) Error() string {
	return fmt.Sprintf("unknown trouble type: %q", e.Value)
}
