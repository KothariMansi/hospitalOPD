package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/KothariMansi/hospitalOPD/db/util"
)

var testqueries *Queries
var testdb *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load configuration:", err)
	}
	testdb, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testqueries = New(testdb)
	os.Exit(m.Run())
}
