package main

import (
	"GoBank/api"
	db "GoBank/db/sqlc"
	"GoBank/util"
	"context"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	// 读取配置文件中的内容
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, config.DbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
