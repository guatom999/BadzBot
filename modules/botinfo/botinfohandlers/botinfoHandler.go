package botinfohandlers

import (
	"context"

	"github.com/bwmarrin/discordgo"
	"github.com/guatom999/BadzBot/modules/botinfo/botinfoUsecases"
)

type IBotinfoHandler interface {
	Help(s *discordgo.Session, i *discordgo.InteractionCreate)
	Test(s *discordgo.Session, i *discordgo.InteractionCreate)
	// GetSharePrice(s *discordgo.Session, i *discordgo.InteractionCreate)
	// GetFollower(s *discordgo.Session, i *discordgo.InteractionCreate)
}

type botinfohandler struct {
	botinfoUsecase botinfoUsecases.IBotinfoUsecase
}

// func NewBotinfoHandler() IBotinfoHandler {
// 	return &botinfohandler{}
// }

func NewBotinfoHandler(botinfoUsecase botinfoUsecases.IBotinfoUsecase) IBotinfoHandler {
	return &botinfohandler{
		botinfoUsecase: botinfoUsecase,
	}
}

func (h *botinfohandler) Help(s *discordgo.Session, i *discordgo.InteractionCreate) {
	command := i.ApplicationCommandData()
	messageContent := command.Options[0].StringValue()

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: h.botinfoUsecase.Feature(messageContent),
		},
	})
}

func (h *botinfohandler) Test(s *discordgo.Session, i *discordgo.InteractionCreate) {
	command := i.ApplicationCommandData()
	messageContent := command.Options[0].StringValue()

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: h.botinfoUsecase.JetTest(messageContent),
		},
	})
}

func (h *botinfohandler) GetSharePrice(s *discordgo.Session, i *discordgo.InteractionCreate) {

	ctx := context.Background()

	command := i.ApplicationCommandData()
	messageContent := command.Options[0].StringValue()

	result, err := h.botinfoUsecase.GetSharePrice(ctx, messageContent)
	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: err.Error(),
			},
		})
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: result.Name,
		},
	})
}

// func (h *botinfohandler) GetFollower(s *discordgo.Session, i *discordgo.InteractionCreate) {
// 	command := i.ApplicationCommandData()
// 	messageContent := command.Options[0].StringValue()

// 	if messageContent == "" {
// 		s.ChannelMessageSend(m.ChannelID, "Pong!")
// 	}

// 	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
// 		Type: discordgo.InteractionResponseChannelMessageWithSource,
// 		Data: &discordgo.InteractionResponseData{
// 			Content: h.botinfoUsecase.GetSharePrice(messageContent),
// 		},
// 	})
// }
