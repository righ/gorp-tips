package repositories_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"

	"gorp-tips/db"
	"gorp-tips/factories"
	"gorp-tips/models"
	"gorp-tips/repositories"
)

func TestRepository(t *testing.T) {
	var deps []db.Dependency

	ken, deps := factories.MakePilot(factories.Fields{"Name": "Ken"}, deps)
	kyle, deps := factories.MakePilot(factories.Fields{"Name": "Kyle"}, deps)
	kim, deps := factories.MakePilot(factories.Fields{"Name": "Kim"}, deps)

	jp, deps := factories.MakeLanguage(factories.Fields{"Language": "Japanese"}, deps)
	en, deps := factories.MakeLanguage(factories.Fields{"Language": "English"}, deps)
	kr, deps := factories.MakeLanguage(factories.Fields{"Language": "Korean"}, deps)

	_, deps = factories.MakePilotLanguage(factories.Fields{"PilotID": ken.ID, "LanguageID": jp.ID}, deps)
	_, deps = factories.MakePilotLanguage(factories.Fields{"PilotID": kyle.ID, "LanguageID": jp.ID}, deps)
	_, deps = factories.MakePilotLanguage(factories.Fields{"PilotID": kyle.ID, "LanguageID": en.ID}, deps)
	_, deps = factories.MakePilotLanguage(factories.Fields{"PilotID": kim.ID, "LanguageID": jp.ID}, deps)
	_, deps = factories.MakePilotLanguage(factories.Fields{"PilotID": kim.ID, "LanguageID": kr.ID}, deps)

	falcon, deps := factories.MakeJet(factories.Fields{"Age": uint8(40), "Name": "Falcon", "PilotID": ken.ID}, deps)
	hawk, deps := factories.MakeJet(factories.Fields{"Age": uint8(30), "Name": "Hawk", "PilotID": kyle.ID}, deps)
	swallow, deps := factories.MakeJet(factories.Fields{"Age": uint8(20), "Name": "Swallow", "PilotID": kyle.ID}, deps)
	dove, deps := factories.MakeJet(factories.Fields{"Age": uint8(10), "Name": "Dove", "Color": "gray"}, deps)
	eagle, deps := factories.MakeJet(factories.Fields{"Age": uint8(10), "Name": "Eagle", "PilotID": kim.ID}, deps)

	cases := []struct {
		name     string
		req      models.Request
		expected []models.Result
	}{
		{
			name: "age filter",
			req:  models.Request{Age: 10},
			expected: []models.Result{
				{JetName: dove.Name, JetAge: 10, JetColor: dove.Color, PilotName: "Tester", Language: nil},
				{JetName: eagle.Name, JetAge: 10, JetColor: eagle.Color, PilotName: kim.Name, Language: &jp.Language},
				{JetName: eagle.Name, JetAge: 10, JetColor: eagle.Color, PilotName: kim.Name, Language: &kr.Language},
			},
		},
		{
			name: "pilot name filter",
			req:  models.Request{PilotName: "en"},
			expected: []models.Result{
				{JetName: falcon.Name, JetAge: falcon.Age, JetColor: falcon.Color, PilotName: ken.Name, Language: &jp.Language},
			},
		},
		{
			name: "jet name filter",
			req:  models.Request{JetName: "awk"},
			expected: []models.Result{
				{JetName: hawk.Name, JetAge: hawk.Age, JetColor: hawk.Color, PilotName: kyle.Name, Language: &jp.Language},
				{JetName: hawk.Name, JetAge: hawk.Age, JetColor: hawk.Color, PilotName: kyle.Name, Language: &en.Language},
			},
		},
		{
			name: "language filter",
			req:  models.Request{Language: "English"},
			expected: []models.Result{
				{JetName: swallow.Name, JetAge: swallow.Age, JetColor: swallow.Color, PilotName: kyle.Name, Language: &en.Language},
				{JetName: hawk.Name, JetAge: hawk.Age, JetColor: hawk.Color, PilotName: kyle.Name, Language: &en.Language},
			},
		},
	}

	db.RunTest(context.Background(), t, func(ctx context.Context, ntx *db.NestableTx) {
		repo := repositories.NewJetRepository(ntx)
		for _, c := range cases {
			t.Run("GetJets "+c.name, func(t *testing.T) {
				got, err := repo.GetJets(ctx, c.req)
				if err != nil {
					t.Error(err)
					return
				}
				if r := cmp.Diff(got, c.expected); r != "" {
					t.Errorf("failed. expected: %v, got: %v", c.expected, got)
				}
			})
		}
	}, deps...)
}
