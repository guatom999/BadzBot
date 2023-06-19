package main

import (
	"github.com/guatom999/BadzBot/configs"
	"github.com/guatom999/BadzBot/modules/server"
)

func main() {
	cfg := configs.NewConfig("./.env")

	server.NewDiscordServer(cfg).Start()
}
