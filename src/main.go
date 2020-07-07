package main

import (
	"database/sql"
	"flag"
	"fmt"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"

	"gorp-with-template/lib"
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
	query := lib.GetSQL("query.sql", req)
	log.Debug(query)

	var results []Result
	if _, err := dbmap.Select(&results, query, map[string]interface{}{
		"age":        req.Age,
		"pilot_name": "%" + req.PilotName + "%",
		"jet_name":   "%" + req.JetName + "%",
		"language":   req.Language,
	}); err != nil {
		log.Error(err)
		return
	}
	for _, record := range results {
		fmt.Printf("%+v\n", record)
	}

}
