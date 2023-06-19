package server

import (
	"github.com/bwmarrin/discordgo"
	"github.com/guatom999/BadzBot/modules/botinfo/botinfoUsecases"
	"github.com/guatom999/BadzBot/modules/botinfo/botinfohandlers"
)

type IBotinfoModule interface {
	Init()
	Handler() botinfohandlers.IBotinfoHandler
	Usecase() botinfoUsecases.IBotinfoUsecase
}

type botInfoModule struct {
	*module
	hanlder botinfohandlers.IBotinfoHandler
	usecase botinfoUsecases.IBotinfoUsecase
}

func (m *module) BotinfoModule() IBotinfoModule {

	botinfoUsecases := botinfoUsecases.NewBotinfoUsecase()
	botinfoHanlder := botinfohandlers.NewBotinfoHandler(botinfoUsecases)

	return &botInfoModule{
		module:  m,
		hanlder: botinfoHanlder,
		usecase: botinfoUsecases,
	}
}

func (b *botInfoModule) Init() {

	b.module.commands = append(b.module.commands, &discordgo.ApplicationCommand{
		Name:        "help",
		Description: "Help Bot Menu",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "feature",
				Description: "Test Bot",
				Type:        discordgo.ApplicationCommandOptionString,
			},
		},
	})

	b.commandHandler["help"] = b.hanlder.Help
}

func (b *botInfoModule) Handler() botinfohandlers.IBotinfoHandler {
	return b.hanlder
}
func (b *botInfoModule) Usecase() botinfoUsecases.IBotinfoUsecase {
	return b.usecase
}
