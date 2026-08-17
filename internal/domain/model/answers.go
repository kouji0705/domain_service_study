package model

import "domain_service_study/internal/domain/valueobject"

// Answers は質問IDと回答の対応。
// 生成後に外部のmapを書き換えても影響しないよう、受け取り時にコピーする。
type Answers struct {
	values map[QuestionID]valueobject.Answer
}

func NewAnswers(raw map[QuestionID]valueobject.Answer) Answers {
	copied := make(map[QuestionID]valueobject.Answer, len(raw))
	for id, value := range raw {
		copied[id] = value
	}
	return Answers{values: copied}
}

func NewAnswersFromText(raw map[QuestionID]string) Answers {
	copied := make(map[QuestionID]valueobject.Answer, len(raw))
	for id, value := range raw {
		copied[id] = valueobject.NewAnswer(value)
	}
	return NewAnswers(copied)
}

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

func (a Answers) Get(id QuestionID) (valueobject.Answer, bool) {
	if a.values == nil {
		return valueobject.Answer{}, false
	}
	value, ok := a.values[id]
	return value, ok
}

func (a Answers) Answer(id QuestionID) valueobject.Answer {
	value, _ := a.Get(id)
	return value
}

func (a Answers) Clone() Answers {
	return NewAnswers(a.values)
}
