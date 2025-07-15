package main

import (
	"database/sql"
	"log"

	"github.com/KothariMansi/hospitalOPD/api"
	db "github.com/KothariMansi/hospitalOPD/db/sqlc"
	"github.com/KothariMansi/hospitalOPD/db/util"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configuration:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(*store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
