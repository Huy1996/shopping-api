package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"shopping-cart/src/api"
	db "shopping-cart/src/db/sqlc"
	"shopping-cart/src/util"
)

func main() {
	config, err := util.LoadConfig("../../../")
	if err != nil {
		log.Fatal("Unable to load configuration file:", err)
	}

	conn, err := sql.Open(config.DBEngine, config.DBSource)
	if err != nil {
		log.Fatal("Unable to create connection with db:", err)
	}

	store := db.NewStore(conn)

	server := api.NewServer(config, store)
	err = server.Start(config.HTTPAddress)
	if err != nil {
		log.Fatal("Unable to start server:", err)
	}
}
