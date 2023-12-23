package server

import (
	"github.com/bwmarrin/discordgo"
	"github.com/guatom999/BadzBot/modules/botinfo/botinfoRepositories"
	"github.com/guatom999/BadzBot/modules/botinfo/botinfoUsecases"
	"github.com/guatom999/BadzBot/modules/botinfo/botinfohandlers"
)

var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "help",
			Description: "Help Bot Menu",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "feature",
					Description: "Test Bot",
					Type:        discordgo.ApplicationCommandOptionString,
				},
			},
		},
		{
			Name:        "response",
			Description: "Test Response",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "text",
					Description: "Test Feature 2",
					Type:        discordgo.ApplicationCommandOptionString,
				},
			},
		},
		{
			Name:        "getshareprice",
			Description: "Get Share Price",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "query",
					Description: "get price of share",
					Type:        discordgo.ApplicationCommandOptionString,
				},
			},
		},
	}

	// _ = commands
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

	botinfoRepo := botinfoRepositories.NewBotRepository()
	botinfoUsecases := botinfoUsecases.NewBotinfoUsecase(botinfoRepo)
	botinfoHanlder := botinfohandlers.NewBotinfoHandler(botinfoUsecases)

	return &botInfoModule{
		module:  m,
		hanlder: botinfoHanlder,
		usecase: botinfoUsecases,
	}
}

func (b *botInfoModule) Init() {

	// registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))

	// b.module.commands = append(b.module.commands, &discordgo.ApplicationCommand{
	// 	Name:        "help",
	// 	Description: "Help Bot Menu",
	// 	Options: []*discordgo.ApplicationCommandOption{
	// 		{
	// 			Name:        "feature",
	// 			Description: "Test Bot",
	// 			Type:        discordgo.ApplicationCommandOptionString,
	// 		},
	// 	},
	// }, &discordgo.ApplicationCommand{
	// 	Name:        "response",
	// 	Description: "Response Bot Test",
	// 	Options: []*discordgo.ApplicationCommandOption{
	// 		{
	// 			Name:        "feature",
	// 			Description: "Test Response",
	// 			Type:        discordgo.ApplicationCommandOptionString,
	// 		},
	// 	},
	// })

	b.module.commands = commands

	b.commandHandler["help"] = b.hanlder.Help
	b.commandHandler["response"] = b.hanlder.Test
	b.commandHandler["getshareprice"] = b.hanlder.GetSharePrice
	// b.commandHandler["getfollower"] = b.hanlder.

}

func (b *botInfoModule) Handler() botinfohandlers.IBotinfoHandler {
	return b.hanlder
}
func (b *botInfoModule) Usecase() botinfoUsecases.IBotinfoUsecase {
	return b.usecase
}
