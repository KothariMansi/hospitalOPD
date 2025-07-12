package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/KothariMansi/hospitalOPD/db/util"
	_ "github.com/go-sql-driver/mysql"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load configurations:", err)
	}
	testDb, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(testDb)
	os.Exit(m.Run())
}
