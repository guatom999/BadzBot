package server

import "github.com/bwmarrin/discordgo"

type IModule interface {
	GetCommandHandler() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
	BotinfoModule() IBotinfoModule
}

type module struct {
	*discordServer
	commandHandler map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

func ModuleInit(s *discordServer) IModule {
	return &module{
		discordServer:  s,
		commandHandler: make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)),
	}
}

func (m *module) GetCommandHandler() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return m.commandHandler
}
