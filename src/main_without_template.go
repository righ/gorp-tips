package main

import (
	"database/sql"
	"flag"
	"fmt"
	"strings"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
)

func initDb() *gorp.DbMap {
	db, _ := sql.Open("mysql", "usr:pw@tcp(mysql:3306)/db")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	return dbmap
}

type Request struct {
	Age       int
	PilotName string
	JetName   string
	Language  string
}

type Result struct {
	JetName   string
	JetAge    int
	JetColor  string
	PilotName string
	Language  string
}

func main() {
	dbmap := initDb()
	defer dbmap.Db.Close()

	var (
		age       = flag.Int("age", 0, "jet age")
		pilotName = flag.String("pilot_name", "", "pilot name")
		jetName   = flag.String("jet name", "", "jet name")
		language  = flag.String("language", "", "language")
	)
	flag.Parse()
	req := Request{
		Age:       *age,
		PilotName: *pilotName,
		JetName:   *jetName,
		Language:  *language,
	}

	query := "SELECT jets.name AS jetName, jets.age AS jetAge, jets.color AS jetColor, pilots.name AS pilotName, languages.language "
	query += "FROM jets "
	query += "JOIN pilots ON pilots.id = jets.pilot_id "
	query += "JOIN pilot_languages ON pilot_languages.pilot_id = jets.pilot_id "
	query += "JOIN languages ON languages.id = pilot_languages.language_id "

	conds, context := makeCondition(req)
	if conds != "" {
		query += "WHERE " + conds
	}

	query += " ORDER BY jets.age"

	log.Debug(query)

	var results []Result
	if _, err := dbmap.Select(&results, query, context); err != nil {
		log.Error(err)
		return
	}
	for _, record := range results {
		fmt.Printf("%+v\n", record)
	}

}

func makeCondition(req Request) (string, map[string]interface{}) {
	conds := []string{}
	context := map[string]interface{}{}

	if req.Age > 0 {
		conds = append(conds, "jets.age = :age")
		context["age"] = req.Age
	}
	if req.PilotName != "" {
		conds = append(conds, "pilots.name LIKE :pilot_name")
		context["pilot_name"] = "%" + req.PilotName + "%"
	}
	if req.JetName != "" {
		conds = append(conds, "jets.name LIKE :jet_name")
		context["jet_name"] = "%" + req.JetName + "%"
	}
	if req.Language != "" {
		conds = append(conds, "languages.language = :language")
		context["language"] = req.Language
	}

	return strings.Join(conds, " AND "), context
}
