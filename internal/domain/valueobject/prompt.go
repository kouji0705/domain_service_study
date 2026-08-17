package valueobject

import "strings"

// Prompt は質問文を表す値オブジェクト。
type Prompt struct {
	value string
}

// NewPrompt は空でない質問文から Prompt を生成する。
// 空白のみの値は EmptyPromptError になる。
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

// String は質問文を返す。
func (p Prompt) String() string { return p.value }

// IsZero は未設定の Prompt かどうかを返す。
func (p Prompt) IsZero() bool { return p.value == "" }

// EmptyPromptError は空の質問文を作ろうとしたときに返す。
type EmptyPromptError struct{}

func (e *EmptyPromptError) Error() string {
	return "prompt must not be empty"
}
