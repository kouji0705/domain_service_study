package valueobject

import "strings"

// Answer は1つの質問に対する回答内容を表す値オブジェクト。
// 空や任意の文言も持てる。yes / no かどうかは IsYes / IsNo で判定する。
type Answer struct {
	value string
}

var (
	// AnswerYes は肯定の回答。
	AnswerYes = Answer{value: "yes"}
	// AnswerNo は否定の回答。
	AnswerNo = Answer{value: "no"}
)

// NewAnswer は回答内容から Answer を生成する。空文字も許可する。
func NewAnswer(value string) Answer {
	return Answer{value: value}
}

// String は回答の文字列表現を返す。
func (a Answer) String() string { return a.value }

// IsBlank は未入力（空白のみを含む）かどうかを返す。
func (a Answer) IsBlank() bool { return strings.TrimSpace(a.value) == "" }

// IsYes は回答が yes かどうかを返す。
func (a Answer) IsYes() bool { return a == AnswerYes }

// IsNo は回答が no かどうかを返す。
func (a Answer) IsNo() bool { return a == AnswerNo }
