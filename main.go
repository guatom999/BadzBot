package main

import (
	"context"

	"github.com/guatom999/BadzBot/config"
	"github.com/guatom999/BadzBot/modules/server"
	"github.com/guatom999/BadzBot/pkg/database"
)

func main() {

	ctx := context.Background()

	cfg := config.NewConfig("./env/.env")

	server.NewDiscordServer(&cfg).Start()

	db := database.DbConnect(ctx, &cfg)

	defer db.Disconnect(ctx)

	// server.NewHttpServer()

}
