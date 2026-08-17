package valueobject

import "strings"

// Answer は1つの質問に対する回答内容を表す値オブジェクト。
type Answer struct {
	value string
}

var (
	AnswerYes = Answer{value: "yes"}
	AnswerNo  = Answer{value: "no"}
)

func NewAnswer(value string) Answer {
	return Answer{value: value}
}

func (a Answer) String() string { return a.value }
func (a Answer) IsBlank() bool  { return strings.TrimSpace(a.value) == "" }
func (a Answer) IsYes() bool    { return a == AnswerYes }
func (a Answer) IsNo() bool     { return a == AnswerNo }
