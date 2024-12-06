package Database

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"github.com/rashid642/banking/utils"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("can not load config, error:", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database, error :", err)
	}
	testQueries = New(testDB)

	os.Exit(m.Run()) 
}