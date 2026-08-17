package casetype

import "domain_service_study/internal/domain/model"

// 質問 ID のカタログ。QuestionID は model パッケージで定義し、
// 具体的な値はこのパッケージが事例知識として保持する。
var (
	// QuestionOverviewSummary は概要セクションの問題内容。
	QuestionOverviewSummary = model.MustQuestionID("overview.summary")
	// QuestionSituationOccurredAt は発生時刻。
	QuestionSituationOccurredAt = model.MustQuestionID("situation.occurred_at")
	// QuestionBackgroundBeforeOccurrence は発生前の操作内容。
	QuestionBackgroundBeforeOccurrence = model.MustQuestionID("background.before_occurrence")
	// QuestionAssessmentPossibleCause は想定される原因。
	QuestionAssessmentPossibleCause = model.MustQuestionID("assessment.possible_cause")
	// QuestionRecommendationRequestedAction は希望する対応。
	QuestionRecommendationRequestedAction = model.MustQuestionID("recommendation.requested_action")
	// QuestionOtherNotes はその他の連絡事項。
	QuestionOtherNotes = model.MustQuestionID("other.notes")

	// QuestionPCPowerOn は PC の電源状態。
	QuestionPCPowerOn = model.MustQuestionID("pc.power_on")
	// QuestionPCPowerLight は電源ランプの点灯状態。
	QuestionPCPowerLight = model.MustQuestionID("pc.power_light")
	// QuestionPCACAdapterConnected は AC アダプター接続状態。
	QuestionPCACAdapterConnected = model.MustQuestionID("pc.ac_adapter_connected")
	// QuestionPCScreenVisible は画面表示の有無。
	QuestionPCScreenVisible = model.MustQuestionID("pc.screen_visible")

	// QuestionNetworkConnectionType は接続方式（Wi-Fi / 有線）。
	QuestionNetworkConnectionType = model.MustQuestionID("network.connection_type")
	// QuestionNetworkOtherUsersAffected は他利用者への影響有無。
	QuestionNetworkOtherUsersAffected = model.MustQuestionID("network.other_users_affected")
	// QuestionNetworkAffectedDeviceCount は影響端末台数。
	QuestionNetworkAffectedDeviceCount = model.MustQuestionID("network.affected_device_count")

	// QuestionImpactAffectedPeople は影響を受けている人数。
	QuestionImpactAffectedPeople = model.MustQuestionID("impact.affected_people")
	// QuestionImpactWorkaround は利用可能な代替手段。
	QuestionImpactWorkaround = model.MustQuestionID("impact.workaround")
)
