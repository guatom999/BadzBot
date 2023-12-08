package server

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/guatom999/BadzBot/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	GuildID = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	// BotToken       = flag.String("token", "", "Bot access token") || doesn't use this
	RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")
)

type IDiscordServer interface {
	Start()
}

type discordServer struct {
	cfg      *config.Config
	dg       *discordgo.Session
	commands []*discordgo.ApplicationCommand
}

func NewDiscordServer(cfg *config.Config) IDiscordServer {
	dg, err := discordgo.New("Bot " + cfg.App.Token)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}

	return &discordServer{
		cfg:      cfg,
		dg:       dg,
		commands: make([]*discordgo.ApplicationCommand, 0),
	}
}

func (s *discordServer) Start() {
	s.dg.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})

	err := s.dg.Open()

	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)

	}

	module := ModuleInit(s)
	module.BotinfoModule().Init()

	log.Println("command avaliable is ==========> :", len(s.commands))

	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(s.commands))
	for i, v := range s.commands {
		cmd, err := s.dg.ApplicationCommandCreate(s.dg.State.User.ID, *GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	defer s.dg.Close()

	s.dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		h, ok := module.GetCommandHandler()[i.ApplicationCommandData().Name]
		if ok {
			h(s, i)
		}
	})

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	if *RemoveCommands {
		log.Println("Removing commands...")

		for _, v := range registeredCommands {
			err := s.dg.ApplicationCommandDelete(s.dg.State.User.ID, *GuildID, v.ID)
			if err != nil {
				log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
			}
		}
	}

	log.Println("Gracefully shutting down.")

}

type IHttpServer interface {
	Start()
}
type HttpServer struct {
	app *echo.Echo
	db  *mongo.Client
	cfg *config.Config
}

func NewHttpServer(db *mongo.Client, cfg *config.Config) IHttpServer {
	return &HttpServer{
		app: echo.New(),
		db:  db,
		cfg: cfg,
	}
}

// func (s *HttpServer) gracefulShutdown(pctx context.Context, close <-chan os.Signal) {

// 	resClose := <-close

// 	if resClose != nil {
// 		log.Println("Shutting down server")

// 		ctx, cancel := context.WithTimeout(pctx, time.Second*10)
// 		defer cancel()

// 		if err := s.app.Shutdown(ctx); err != nil {
// 			log.Fatalf("Failed to shutdown:%v", err)
// 		}
// 	}

// }

func (s *HttpServer) Start(pctx context.Context) {

	s.app.Use(middleware.Logger())

	if err := s.app.Start(s.cfg.App.AppUrl); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to shutdown:%v", err)
	}

	close := make(chan os.Signal, 1)
	signal.Notify(close, syscall.SIGINT, syscall.SIGTERM)

	go s.gracefulShutdown(pctx, close)

}
