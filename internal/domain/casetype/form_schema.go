package casetype

import "domain_service_study/internal/domain/model"

// mustFormSchema は静的な質問項目定義から FormSchema を生成する。
// カタログ内の定義が不正な場合は panic する。
func mustFormSchema(questions ...model.QuestionDefinition) model.FormSchema {
	def, err := model.NewFormSchema(questions...)
	if err != nil {
		// 事例カタログ内の静的な質問定義が不正なときのプログラミングエラー。
		panic(err)
	}
	return def
}
