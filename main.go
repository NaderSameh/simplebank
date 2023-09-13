package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/nadersameh_/simplebank/api"
	db "github.com/nadersameh_/simplebank/db/sqlc"
	"github.com/nadersameh_/simplebank/util"
)

func main() {
	config, err := util.GetConfig(".")
	if err != nil {
		log.Fatal("can't load configurations")
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.Serveraddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
