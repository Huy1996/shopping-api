package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"shopping-cart/src/util"
	"testing"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../../")
	if err != nil {
		log.Fatal("Unable to load configuration file:", err)
	}

	testDB, err = sql.Open(config.DBEngine, config.DBSource)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
