package botEvents

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func CreateAmazingEvent(s *discordgo.Session) *discordgo.GuildScheduledEvent {
	VoiceChannelID := "c1186018141927653407"
	GuildID := "335799834965573642"
	// Define the starting time (must be in future)
	startingTime := time.Now().Add(10 * time.Second)
	// Define the ending time (must be after starting time)
	endingTime := startingTime.Add(10 * time.Second)
	// Create the event
	scheduledEvent, err := s.GuildScheduledEventCreate(GuildID, &discordgo.GuildScheduledEventParams{
		Name:               "Amazing Event",
		Description:        "This event will start in 1 hour and last 30 minutes",
		ScheduledStartTime: &startingTime,
		ScheduledEndTime:   &endingTime,
		EntityType:         discordgo.GuildScheduledEventEntityTypeVoice,
		ChannelID:          VoiceChannelID,
		PrivacyLevel:       discordgo.GuildScheduledEventPrivacyLevelGuildOnly,
	})
	if err != nil {
		log.Printf("Error creating scheduled event: %v", err)
		return nil
	}

	fmt.Println("Created scheduled event:", scheduledEvent.Name)
	return scheduledEvent
}

func TransformEventToExternalEvent(s *discordgo.Session, event *discordgo.GuildScheduledEvent) {
	GuildID := "335799834965573642"
	scheduledEvent, err := s.GuildScheduledEventEdit(GuildID, event.ID, &discordgo.GuildScheduledEventParams{
		Name:       "Amazing Event @ Discord Website",
		EntityType: discordgo.GuildScheduledEventEntityTypeExternal,
		EntityMetadata: &discordgo.GuildScheduledEventEntityMetadata{
			Location: "https://discord.com",
		},
	})
	if err != nil {
		log.Printf("Error during transformation of scheduled voice event into external event: %v", err)
		return
	}

	fmt.Println("Created scheduled event:", scheduledEvent.Name)
}
