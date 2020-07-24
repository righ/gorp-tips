package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"gorp-tips/controllers"
	"gorp-tips/models"
	"gorp-tips/repositories"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/log"
)

func initDb() *gorp.DbMap {
	db, _ := sql.Open("mysql", "usr:pw@tcp(mysql:3306)/db")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "utf8mb4"}}
	models.MapStructsToTables(dbmap)
	return dbmap
}

func main() {
	dbmap := initDb()
	defer dbmap.Db.Close()

	var (
		age       = flag.Int("age", 0, "jet age")
		pilotName = flag.String("pilot_name", "", "pilot name")
		jetName   = flag.String("jet_name", "", "jet name")
		language  = flag.String("language", "", "language")
	)
	flag.Parse()
	req := models.Request{
		Age:       *age,
		PilotName: *pilotName,
		JetName:   *jetName,
		Language:  *language,
	}
	repo := repositories.NewJetRepository2(dbmap)
	results, err := controllers.GetJets(context.Background(), repo, req)
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, record := range results {
		fmt.Printf("%+v\n", record)
	}
}
