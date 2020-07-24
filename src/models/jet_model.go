package models

import (
	"github.com/go-gorp/gorp"
)

type Jet struct {
	ID      int    `db:"id"`
	PilotID int    `db:"pilot_id"`
	Age     uint8  `db:"age"`
	Name    string `db:"name"`
	Color   string `db:"color"`
}

type Pilot struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type Language struct {
	ID       int    `db:"id"`
	Language string `db:"language"`
}

type PilotLanguage struct {
	PilotID    int `db:"pilot_id"`
	LanguageID int `db:"language_id"`
}

// MapStructsToTables 構造体と物理テーブルの紐付け
func MapStructsToTables(dbmap *gorp.DbMap) {
	dbmap.AddTableWithName(Pilot{}, "pilots")
	dbmap.AddTableWithName(Jet{}, "jets")
	dbmap.AddTableWithName(Language{}, "languages")
	dbmap.AddTableWithName(PilotLanguage{}, "pilot_languages")
}
