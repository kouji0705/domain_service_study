package model

import "domain_service_study/internal/domain/valueobject"

// Answers は質問IDと回答の対応。
// 生成後に外部のmapを書き換えても影響しないよう、受け取り時にコピーする。
type Answers struct {
	values map[QuestionID]valueobject.Answer
}

// NewAnswers は QuestionID と Answer の対応から Answers を生成する。
// 受け取った map はコピーされるため、呼び出し側の変更の影響を受けない。
func NewAnswers(raw map[QuestionID]valueobject.Answer) Answers {
	copied := make(map[QuestionID]valueobject.Answer, len(raw))
	for id, value := range raw {
		copied[id] = value
	}
	return Answers{values: copied}
}

// NewAnswersFromText は QuestionID と文字列回答の対応から Answers を生成する。
func NewAnswersFromText(raw map[QuestionID]string) Answers {
	copied := make(map[QuestionID]valueobject.Answer, len(raw))
	for id, value := range raw {
		copied[id] = valueobject.NewAnswer(value)
	}
	return NewAnswers(copied)
}

// NewAnswersFromStrings は文字列キーの map から Answers を生成する。
// キーが空の QuestionID になる場合は EmptyIdentityError を返す。
func NewAnswersFromStrings(raw map[string]string) (Answers, error) {
	copied := make(map[QuestionID]valueobject.Answer, len(raw))
	for id, value := range raw {
		questionID, err := NewQuestionID(id)
		if err != nil {
			return Answers{}, err
		}
		copied[questionID] = valueobject.NewAnswer(value)
	}
	return NewAnswers(copied), nil
}

// Get は指定した質問 ID の回答を返す。
func (a Answers) Get(id QuestionID) (valueobject.Answer, bool) {
	if a.values == nil {
		return valueobject.Answer{}, false
	}
	value, ok := a.values[id]
	return value, ok
}

// Answer は指定した質問 ID の回答を返す。未回答の場合は空の Answer を返す。
func (a Answers) Answer(id QuestionID) valueobject.Answer {
	value, _ := a.Get(id)
	return value
}

// Clone は Answers のコピーを返す。
func (a Answers) Clone() Answers {
	return NewAnswers(a.values)
}
