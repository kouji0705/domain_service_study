package casetype

import "domain_service_study/internal/domain/model"

func mustDefinition(questions ...model.QuestionDefinition) model.ReportDefinition {
	def, err := model.NewReportDefinition(questions...)
	if err != nil {
		// 事例カタログ内の静的な質問定義が不正なときのプログラミングエラー。
		panic(err)
	}
	return def
}
