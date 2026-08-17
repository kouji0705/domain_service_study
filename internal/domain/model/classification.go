package model

import "domain_service_study/internal/domain/identity"

// TroubleBrand はトラブル種類の Brand。具体的な種類の値は casetype が持つ。
type TroubleBrand struct{}

type TroubleType = identity.Identity[TroubleBrand]

func NewTroubleType(value string) (TroubleType, error) {
	return identity.NewIdentity[TroubleBrand](value)
}

func MustTroubleType(value string) TroubleType {
	return identity.MustIdentity[TroubleBrand](value)
}

// ImpactBrand は影響度の Brand。具体的な段階の値は casetype が持つ。
type ImpactBrand struct{}

type ImpactLevel = identity.Identity[ImpactBrand]

func NewImpactLevel(value string) (ImpactLevel, error) {
	return identity.NewIdentity[ImpactBrand](value)
}

func MustImpactLevel(value string) ImpactLevel {
	return identity.MustIdentity[ImpactBrand](value)
}
