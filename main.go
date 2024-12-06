package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	db "github.com/rashid642/banking/Database/sqlc"
	"github.com/rashid642/banking/api"
	"github.com/rashid642/banking/utils"
)

func main() {
	config, err := utils.LoadConfig(".") 
	if err != nil {
		log.Fatal("Can not load config files, err:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database, error :", err)
	}

	store := db.NewStore(conn) 
	server := api.NewServer(store) 

	err = server.Start(config.ServerAddress) 
	if err != nil {
		log.Fatal("Can not start the server, err :", err)
	}
}