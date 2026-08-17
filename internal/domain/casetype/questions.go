package casetype

import "domain_service_study/internal/domain/model"

var (
	QuestionOverviewSummary               = model.MustQuestionID("overview.summary")
	QuestionSituationOccurredAt           = model.MustQuestionID("situation.occurred_at")
	QuestionBackgroundBeforeOccurrence    = model.MustQuestionID("background.before_occurrence")
	QuestionAssessmentPossibleCause       = model.MustQuestionID("assessment.possible_cause")
	QuestionRecommendationRequestedAction = model.MustQuestionID("recommendation.requested_action")
	QuestionOtherNotes                    = model.MustQuestionID("other.notes")

	QuestionPCPowerOn            = model.MustQuestionID("pc.power_on")
	QuestionPCPowerLight         = model.MustQuestionID("pc.power_light")
	QuestionPCACAdapterConnected = model.MustQuestionID("pc.ac_adapter_connected")
	QuestionPCScreenVisible      = model.MustQuestionID("pc.screen_visible")

	QuestionNetworkConnectionType      = model.MustQuestionID("network.connection_type")
	QuestionNetworkOtherUsersAffected  = model.MustQuestionID("network.other_users_affected")
	QuestionNetworkAffectedDeviceCount = model.MustQuestionID("network.affected_device_count")

	QuestionImpactAffectedPeople = model.MustQuestionID("impact.affected_people")
	QuestionImpactWorkaround     = model.MustQuestionID("impact.workaround")
)
