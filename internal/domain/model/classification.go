package model

import "domain_service_study/internal/domain/identity"

// TroubleBrand はトラブル種類の Brand。具体的な種類の値は casetype が持つ。
type TroubleBrand struct{}

// TroubleType はトラブル種類の身元。pc や network といった既知の値は casetype が定義する。
type TroubleType = identity.Identity[TroubleBrand]

// NewTroubleType は空でない文字列から TroubleType を生成する。
// 既知かどうかは判定しない。既知の値への変換は casetype.ParseTroubleType を使う。
func NewTroubleType(value string) (TroubleType, error) {
	return identity.NewIdentity[TroubleBrand](value)
}

// MustTroubleType は静的なカタログ定義向け。空文字はプログラミングエラーとして扱う。
func MustTroubleType(value string) TroubleType {
	return identity.MustIdentity[TroubleBrand](value)
}

// ImpactBrand は影響度の Brand。具体的な段階の値は casetype が持つ。
type ImpactBrand struct{}

// ImpactLevel は影響度の身元。individual などの既知の値は casetype が定義する。
type ImpactLevel = identity.Identity[ImpactBrand]

// NewImpactLevel は空でない文字列から ImpactLevel を生成する。
// 既知かどうかは判定しない。既知の値への変換は casetype.ParseImpactLevel を使う。
func NewImpactLevel(value string) (ImpactLevel, error) {
	return identity.NewIdentity[ImpactBrand](value)
}

// MustImpactLevel は静的なカタログ定義向け。空文字はプログラミングエラーとして扱う。
func MustImpactLevel(value string) ImpactLevel {
	return identity.MustIdentity[ImpactBrand](value)
}
