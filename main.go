package main

import (
	"database/sql"
	"log"

	"github.com/achal1304/simple_bank/api"
	db "github.com/achal1304/simple_bank/db/sqlc"
	"github.com/achal1304/simple_bank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load configurations")
	}
	conn, err := sql.Open(config.DBdriver, config.DBsoruce)
	if err != nil {
		log.Fatal("Cannot connect to db", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
}
