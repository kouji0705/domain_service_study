package domainservice

import (
	"errors"
	"slices"
	"testing"

	"domain_service_study/internal/domain/casetype"
	"domain_service_study/internal/domain/model"
	"domain_service_study/internal/domain/valueobject"
)

func TestDefinition(t *testing.T) {
	service := NewTroubleReportService(NewDefinitionComposer())

	tests := []struct {
		name         string
		troubleType  model.TroubleType
		impactLevel  model.ImpactLevel
		answers      map[model.QuestionID]string
		wantIDs      []model.QuestionID
		wantAbsent   []model.QuestionID
		wantRequired []model.QuestionID
	}{
		{
			name:        "共通質問が常に含まれる",
			troubleType: casetype.TroubleTypePC,
			impactLevel: casetype.ImpactLevelIndividual,
			wantIDs: []model.QuestionID{
				casetype.QuestionOverviewSummary,
				casetype.QuestionSituationOccurredAt,
				casetype.QuestionBackgroundBeforeOccurrence,
				casetype.QuestionAssessmentPossibleCause,
				casetype.QuestionRecommendationRequestedAction,
				casetype.QuestionOtherNotes,
			},
		},
		{
			name:        "PCの場合はpc.power_onが含まれる",
			troubleType: casetype.TroubleTypePC,
			impactLevel: casetype.ImpactLevelIndividual,
			wantIDs:     []model.QuestionID{casetype.QuestionPCPowerOn},
		},
		{
			name:        "PCでpc.power_on=noの場合は電源ランプとACアダプターの質問が追加される",
			troubleType: casetype.TroubleTypePC,
			impactLevel: casetype.ImpactLevelIndividual,
			answers: map[model.QuestionID]string{
				casetype.QuestionPCPowerOn: valueobject.AnswerNo.String(),
			},
			wantIDs: []model.QuestionID{
				casetype.QuestionPCPowerLight,
				casetype.QuestionPCACAdapterConnected,
			},
			wantAbsent: []model.QuestionID{casetype.QuestionPCScreenVisible},
		},
		{
			name:        "PCでpc.power_on=yesの場合は画面表示の質問が追加される",
			troubleType: casetype.TroubleTypePC,
			impactLevel: casetype.ImpactLevelIndividual,
			answers: map[model.QuestionID]string{
				casetype.QuestionPCPowerOn: valueobject.AnswerYes.String(),
			},
			wantIDs: []model.QuestionID{casetype.QuestionPCScreenVisible},
			wantAbsent: []model.QuestionID{
				casetype.QuestionPCPowerLight,
				casetype.QuestionPCACAdapterConnected,
			},
		},
		{
			name:        "PCで電源の回答がない場合は分岐後の質問がまだ追加されない",
			troubleType: casetype.TroubleTypePC,
			impactLevel: casetype.ImpactLevelIndividual,
			wantAbsent: []model.QuestionID{
				casetype.QuestionPCPowerLight,
				casetype.QuestionPCACAdapterConnected,
				casetype.QuestionPCScreenVisible,
			},
		},
		{
			name:        "ネットワークの場合はネットワーク固有の質問が含まれる",
			troubleType: casetype.TroubleTypeNetwork,
			impactLevel: casetype.ImpactLevelIndividual,
			wantIDs: []model.QuestionID{
				casetype.QuestionNetworkConnectionType,
				casetype.QuestionNetworkOtherUsersAffected,
			},
			wantAbsent: []model.QuestionID{casetype.QuestionPCPowerOn},
		},
		{
			name:        "他の利用者にも影響がある場合は影響端末数の質問が追加される",
			troubleType: casetype.TroubleTypeNetwork,
			impactLevel: casetype.ImpactLevelIndividual,
			answers: map[model.QuestionID]string{
				casetype.QuestionNetworkOtherUsersAffected: valueobject.AnswerYes.String(),
			},
			wantIDs: []model.QuestionID{casetype.QuestionNetworkAffectedDeviceCount},
		},
		{
			name:        "影響度が個人の場合は影響人数の質問が追加されない",
			troubleType: casetype.TroubleTypePC,
			impactLevel: casetype.ImpactLevelIndividual,
			wantAbsent:  []model.QuestionID{casetype.QuestionImpactAffectedPeople, casetype.QuestionImpactWorkaround},
		},
		{
			name:        "影響度がチームの場合は影響人数の質問が追加される",
			troubleType: casetype.TroubleTypePC,
			impactLevel: casetype.ImpactLevelTeam,
			wantIDs:     []model.QuestionID{casetype.QuestionImpactAffectedPeople},
			wantAbsent:  []model.QuestionID{casetype.QuestionImpactWorkaround},
		},
		{
			name:        "影響度が全社の場合は代替手段と希望対応が必須になる",
			troubleType: casetype.TroubleTypePC,
			impactLevel: casetype.ImpactLevelCompanyWide,
			wantIDs: []model.QuestionID{
				casetype.QuestionImpactAffectedPeople,
				casetype.QuestionImpactWorkaround,
				casetype.QuestionRecommendationRequestedAction,
			},
			wantRequired: []model.QuestionID{
				casetype.QuestionImpactWorkaround,
				casetype.QuestionRecommendationRequestedAction,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			def, err := service.Definition(tt.troubleType, tt.impactLevel, answersOf(tt.answers))
			if err != nil {
				t.Fatalf("Definition() error = %v", err)
			}
			assertContains(t, def, tt.wantIDs...)
			assertAbsent(t, def, tt.wantAbsent...)
			assertRequired(t, def, tt.wantRequired...)
		})
	}
}

func TestDefinitionUnknownValues(t *testing.T) {
	service := NewTroubleReportService(NewDefinitionComposer())

	t.Run("未知のトラブル種類はエラーになる", func(t *testing.T) {
		_, err := service.Definition(model.TroubleType{}, casetype.ImpactLevelIndividual, answersOf(nil))
		var unknown *casetype.UnknownTroubleTypeError
		if !errors.As(err, &unknown) {
			t.Fatalf("Definition() error = %v, want UnknownTroubleTypeError", err)
		}
	})

	t.Run("未知の影響度はエラーになる", func(t *testing.T) {
		_, err := service.Definition(casetype.TroubleTypePC, model.ImpactLevel{}, answersOf(nil))
		var unknown *casetype.UnknownImpactLevelError
		if !errors.As(err, &unknown) {
			t.Fatalf("Definition() error = %v, want UnknownImpactLevelError", err)
		}
	})
}

func TestDefinitionDuplicateQuestionID(t *testing.T) {
	common := casetype.CommonModule{}.Definition()
	duplicated, err := model.NewReportDefinition(
		model.NewQuestionDefinition(casetype.QuestionOverviewSummary, valueobject.SectionOverview, valueobject.MustPrompt("重複"), true),
	)
	if err != nil {
		t.Fatalf("NewReportDefinition() error = %v", err)
	}

	_, err = common.Combine(duplicated)
	var dup *model.DuplicateQuestionIDError
	if !errors.As(err, &dup) {
		t.Fatalf("Combine() error = %v, want DuplicateQuestionIDError", err)
	}
	if dup.ID != casetype.QuestionOverviewSummary {
		t.Fatalf("DuplicateQuestionIDError.ID = %s, want %s", dup.ID, casetype.QuestionOverviewSummary)
	}
}

func TestCreate(t *testing.T) {
	service := NewTroubleReportService(NewDefinitionComposer())

	t.Run("必須回答がすべて揃っていればTroubleReportを生成できる", func(t *testing.T) {
		report, err := service.Create(casetype.TroubleTypePC, casetype.ImpactLevelTeam, answersOf(validPCTeamPowerOffAnswers()))
		if err != nil {
			t.Fatalf("Create() error = %v", err)
		}
		if report == nil {
			t.Fatal("Create() report = nil")
		}
	})

	t.Run("必須回答が不足していれば生成できない", func(t *testing.T) {
		answers := validPCTeamPowerOffAnswers()
		delete(answers, casetype.QuestionOverviewSummary)

		_, err := service.Create(casetype.TroubleTypePC, casetype.ImpactLevelTeam, answersOf(answers))
		if err == nil {
			t.Fatal("Create() error = nil, want validation error")
		}
	})

	t.Run("エラーから不足しているQuestionIDを取得できる", func(t *testing.T) {
		answers := validPCTeamPowerOffAnswers()
		delete(answers, casetype.QuestionSituationOccurredAt)
		delete(answers, casetype.QuestionImpactAffectedPeople)

		_, err := service.Create(casetype.TroubleTypePC, casetype.ImpactLevelTeam, answersOf(answers))
		var verr *model.ValidationError
		if !errors.As(err, &verr) {
			t.Fatalf("Create() error = %v, want ValidationError", err)
		}
		got := verr.QuestionIDs()
		if !containsID(got, casetype.QuestionSituationOccurredAt) {
			t.Fatalf("missing QuestionIDs = %v, want %s", got, casetype.QuestionSituationOccurredAt)
		}
		if !containsID(got, casetype.QuestionImpactAffectedPeople) {
			t.Fatalf("missing QuestionIDs = %v, want %s", got, casetype.QuestionImpactAffectedPeople)
		}
	})

	t.Run("PC固有のValidationが実行される", func(t *testing.T) {
		answers := validPCTeamPowerOffAnswers()
		delete(answers, casetype.QuestionPCPowerLight)

		_, err := service.Create(casetype.TroubleTypePC, casetype.ImpactLevelTeam, answersOf(answers))
		assertMissing(t, err, casetype.QuestionPCPowerLight)
	})

	t.Run("ネットワーク固有のValidationが実行される", func(t *testing.T) {
		answers := validNetworkIndividualAnswers()
		answers[casetype.QuestionNetworkOtherUsersAffected] = valueobject.AnswerYes.String()

		_, err := service.Create(casetype.TroubleTypeNetwork, casetype.ImpactLevelIndividual, answersOf(answers))
		assertMissing(t, err, casetype.QuestionNetworkAffectedDeviceCount)
	})

	t.Run("影響度固有のValidationが実行される", func(t *testing.T) {
		answers := validPCTeamPowerOffAnswers()
		delete(answers, casetype.QuestionImpactAffectedPeople)

		_, err := service.Create(casetype.TroubleTypePC, casetype.ImpactLevelTeam, answersOf(answers))
		assertMissing(t, err, casetype.QuestionImpactAffectedPeople)

		answers = validPCCompanyWidePowerOffAnswers()
		delete(answers, casetype.QuestionRecommendationRequestedAction)
		delete(answers, casetype.QuestionImpactWorkaround)

		_, err = service.Create(casetype.TroubleTypePC, casetype.ImpactLevelCompanyWide, answersOf(answers))
		assertMissing(t, err, casetype.QuestionRecommendationRequestedAction)
		assertMissing(t, err, casetype.QuestionImpactWorkaround)
	})

	t.Run("生成された報告書がOverview、SBAR、Otherに回答を正しく保持する", func(t *testing.T) {
		answers := validPCTeamPowerOffAnswers()
		report, err := service.Create(casetype.TroubleTypePC, casetype.ImpactLevelTeam, answersOf(answers))
		if err != nil {
			t.Fatalf("Create() error = %v", err)
		}

		assertSectionValue(t, report.Overview(), casetype.QuestionOverviewSummary, answers[casetype.QuestionOverviewSummary])
		assertSectionValue(t, report.Overview(), casetype.QuestionImpactAffectedPeople, answers[casetype.QuestionImpactAffectedPeople])
		assertSectionValue(t, report.SBAR().Situation(), casetype.QuestionSituationOccurredAt, answers[casetype.QuestionSituationOccurredAt])
		assertSectionValue(t, report.SBAR().Situation(), casetype.QuestionPCPowerOn, answers[casetype.QuestionPCPowerOn])
		assertSectionValue(t, report.SBAR().Background(), casetype.QuestionBackgroundBeforeOccurrence, answers[casetype.QuestionBackgroundBeforeOccurrence])
		assertSectionValue(t, report.SBAR().Background(), casetype.QuestionPCACAdapterConnected, answers[casetype.QuestionPCACAdapterConnected])
		assertSectionValue(t, report.SBAR().Assessment(), casetype.QuestionAssessmentPossibleCause, answers[casetype.QuestionAssessmentPossibleCause])
		assertSectionValue(t, report.SBAR().Assessment(), casetype.QuestionPCPowerLight, answers[casetype.QuestionPCPowerLight])
		assertSectionValue(t, report.SBAR().Recommendation(), casetype.QuestionRecommendationRequestedAction, answers[casetype.QuestionRecommendationRequestedAction])
		assertSectionValue(t, report.Other(), casetype.QuestionOtherNotes, answers[casetype.QuestionOtherNotes])
	})

	t.Run("入力に使用したmapを生成後に変更しても生成済み報告書が変化しない", func(t *testing.T) {
		raw := map[string]string{
			casetype.QuestionOverviewSummary.String():               "元の概要",
			casetype.QuestionSituationOccurredAt.String():           "2026-08-17 09:00",
			casetype.QuestionBackgroundBeforeOccurrence.String():    "作業中",
			casetype.QuestionAssessmentPossibleCause.String():       "不明",
			casetype.QuestionRecommendationRequestedAction.String(): "点検希望",
			casetype.QuestionOtherNotes.String():                    "特になし",
			casetype.QuestionPCPowerOn.String():                     valueobject.AnswerNo.String(),
			casetype.QuestionPCPowerLight.String():                  valueobject.AnswerNo.String(),
			casetype.QuestionPCACAdapterConnected.String():          valueobject.AnswerYes.String(),
			casetype.QuestionImpactAffectedPeople.String():          "3",
		}
		answers, err := model.NewAnswersFromStrings(raw)
		if err != nil {
			t.Fatalf("NewAnswersFromStrings() error = %v", err)
		}

		report, err := service.Create(casetype.TroubleTypePC, casetype.ImpactLevelTeam, answers)
		if err != nil {
			t.Fatalf("Create() error = %v", err)
		}

		raw[casetype.QuestionOverviewSummary.String()] = "書き換え後の概要"
		got, ok := report.Overview().Value(casetype.QuestionOverviewSummary)
		if !ok {
			t.Fatal("overview.summary is missing")
		}
		if got.String() != "元の概要" {
			t.Fatalf("Overview summary = %q, want %q", got, "元の概要")
		}

		items := report.Overview().Items()
		items[0] = model.AnswerItem{}
		got, _ = report.Overview().Value(casetype.QuestionOverviewSummary)
		if got.String() != "元の概要" {
			t.Fatalf("Overview summary after slice mutation = %q, want %q", got, "元の概要")
		}
	})
}

func answersOf(raw map[model.QuestionID]string) model.Answers {
	return model.NewAnswersFromText(raw)
}

func validPCTeamPowerOffAnswers() map[model.QuestionID]string {
	return map[model.QuestionID]string{
		casetype.QuestionOverviewSummary:               "PCの電源が入らない",
		casetype.QuestionSituationOccurredAt:           "2026-08-17 09:00",
		casetype.QuestionBackgroundBeforeOccurrence:    "会議室へ移動した",
		casetype.QuestionAssessmentPossibleCause:       "バッテリー切れ",
		casetype.QuestionRecommendationRequestedAction: "代替機の貸出",
		casetype.QuestionOtherNotes:                    "会議資料の編集が必要",
		casetype.QuestionPCPowerOn:                     valueobject.AnswerNo.String(),
		casetype.QuestionPCPowerLight:                  valueobject.AnswerNo.String(),
		casetype.QuestionPCACAdapterConnected:          valueobject.AnswerYes.String(),
		casetype.QuestionImpactAffectedPeople:          "4",
	}
}

func validPCCompanyWidePowerOffAnswers() map[model.QuestionID]string {
	answers := validPCTeamPowerOffAnswers()
	answers[casetype.QuestionImpactWorkaround] = "共有端末を利用する"
	return answers
}

func validNetworkIndividualAnswers() map[model.QuestionID]string {
	return map[model.QuestionID]string{
		casetype.QuestionOverviewSummary:            "社内ネットワークに接続できない",
		casetype.QuestionSituationOccurredAt:        "2026-08-17 10:00",
		casetype.QuestionBackgroundBeforeOccurrence: "出社してノートPCを起動した",
		casetype.QuestionNetworkConnectionType:      "wifi",
		casetype.QuestionNetworkOtherUsersAffected:  valueobject.AnswerNo.String(),
	}
}

func assertContains(t *testing.T, def model.ReportDefinition, ids ...model.QuestionID) {
	t.Helper()
	for _, id := range ids {
		if !def.Contains(id) {
			t.Fatalf("definition does not contain %s", id)
		}
	}
}

func assertAbsent(t *testing.T, def model.ReportDefinition, ids ...model.QuestionID) {
	t.Helper()
	for _, id := range ids {
		if def.Contains(id) {
			t.Fatalf("definition unexpectedly contains %s", id)
		}
	}
}

func assertRequired(t *testing.T, def model.ReportDefinition, ids ...model.QuestionID) {
	t.Helper()
	for _, id := range ids {
		question, ok := def.Question(id)
		if !ok {
			t.Fatalf("definition does not contain %s", id)
		}
		if !question.Required() {
			t.Fatalf("%s should be required", id)
		}
	}
}

func assertMissing(t *testing.T, err error, id model.QuestionID) {
	t.Helper()
	var verr *model.ValidationError
	if !errors.As(err, &verr) {
		t.Fatalf("error = %v, want ValidationError", err)
	}
	if !containsID(verr.QuestionIDs(), id) {
		t.Fatalf("missing QuestionIDs = %v, want %s", verr.QuestionIDs(), id)
	}
}

func assertSectionValue(t *testing.T, section model.SectionAnswers, id model.QuestionID, want string) {
	t.Helper()
	got, ok := section.Value(id)
	if !ok {
		t.Fatalf("section does not contain %s", id)
	}
	if got.String() != want {
		t.Fatalf("%s = %q, want %q", id, got, want)
	}
}

func containsID(ids []model.QuestionID, id model.QuestionID) bool {
	return slices.Contains(ids, id)
}
