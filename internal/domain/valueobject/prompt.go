package valueobject

import "strings"

// Prompt は質問文を表す値オブジェクト。
type Prompt struct {
	value string
}

func NewPrompt(value string) (Prompt, error) {
	if strings.TrimSpace(value) == "" {
		return Prompt{}, &EmptyPromptError{}
	}
	return Prompt{value: value}, nil
}

// MustPrompt は静的な質問定義向け。空文字はプログラミングエラーとして扱う。
func MustPrompt(value string) Prompt {
	prompt, err := NewPrompt(value)
	if err != nil {
		panic(err)
	}
	return prompt
}

func (p Prompt) String() string { return p.value }
func (p Prompt) IsZero() bool   { return p.value == "" }

type EmptyPromptError struct{}

func (e *EmptyPromptError) Error() string {
	return "prompt must not be empty"
}
