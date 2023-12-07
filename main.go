package main

import (
	"github.com/guatom999/BadzBot/config"
	"github.com/guatom999/BadzBot/modules/server"
)

func main() {
	cfg := config.NewConfig("./.env")

	server.NewDiscordServer(&cfg).Start()

}
