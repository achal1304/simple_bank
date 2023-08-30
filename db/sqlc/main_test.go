package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/achal1304/simple_bank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	var err error
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load configurations")
	}

	testDb, err = sql.Open(config.DBdriver, config.DBsoruce)
	if err != nil {
		log.Fatal("Cannot connect to db", err)
	}
	testQueries = New(testDb)

	os.Exit(m.Run())
}
