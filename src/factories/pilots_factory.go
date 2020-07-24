package factories

import (
	"github.com/bluele/factory-go/factory"

	"gorp-tips/db"
	"gorp-tips/models"
)

var PilotFactory = factory.NewFactory(
	&models.Pilot{},
).SeqInt("ID", func(n int) (interface{}, error) {
	return n, nil
}).Attr("Name", func(args factory.Args) (interface{}, error) {
	return "Tester", nil
})

// MakePilot Pilotのファクトリを作る
func MakePilot(fields Fields, deps []db.Dependency) (*models.Pilot, []db.Dependency) {
	m := PilotFactory.MustCreateWithOption(fields).(*models.Pilot)
	deps = append(deps, m)
	return m, deps
}

var LanguageFactory = factory.NewFactory(
	&models.Language{},
).SeqInt("ID", func(n int) (interface{}, error) {
	return n, nil
}).Attr("Language", func(args factory.Args) (interface{}, error) {
	return "English", nil
})

// MakeLanguage Languageのファクトリを作る
func MakeLanguage(fields Fields, deps []db.Dependency) (*models.Language, []db.Dependency) {
	m := LanguageFactory.MustCreateWithOption(fields).(*models.Language)
	deps = append(deps, m)
	return m, deps
}

var PilotLanguageFactory = factory.NewFactory(
	&models.PilotLanguage{},
)

// MakePilotLanguage PilotLanguageのファクトリを作る
func MakePilotLanguage(fields Fields, deps []db.Dependency) (*models.PilotLanguage, []db.Dependency) {
	m := PilotLanguageFactory.MustCreateWithOption(fields).(*models.PilotLanguage)
	if m.PilotID == 0 {
		pilot, _deps := MakePilot(nil, nil)
		m.PilotID = pilot.ID
		deps = append(deps, _deps...)
	}
	if m.LanguageID == 0 {
		lang, _deps := MakeLanguage(nil, nil)
		m.LanguageID = lang.ID
		deps = append(deps, _deps...)
	}
	deps = append(deps, m)
	return m, deps
}
