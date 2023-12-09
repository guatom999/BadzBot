package main

import (
	"context"
	"log"
	"os"

	"github.com/guatom999/BadzBot/config"
	"github.com/guatom999/BadzBot/modules/server"
	"github.com/guatom999/BadzBot/pkg/database"
)

func main() {

	ctx := context.Background()

	cfg := config.NewConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error: .env path is required")
		}
		log.Printf("choosen env is %s", os.Args[1])

		return os.Args[1]
	}())

	db := database.DbConnect(ctx, &cfg)

	defer db.Disconnect(ctx)

	switch cfg.App.Name {
	case "discord":
		server.NewDiscordServer(&cfg).Start()
	case "normal":
		server.NewHttpServer(db, &cfg).Start(ctx)
	}

	log.Println(db)

}
