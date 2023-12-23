package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/guatom999/BadzBot/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

type IHttpServer interface {
	Start(pctx context.Context)
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

func (s *HttpServer) gracefulShutdown(pctx context.Context, close <-chan os.Signal) {

	resClose := <-close

	if resClose != nil {
		log.Println("Shutting down server")

		ctx, _ := context.WithTimeout(pctx, time.Second*10)
		// defer cancel()

		if err := s.app.Shutdown(ctx); err != nil {
			log.Fatalf("Failed to shutdown:%v", err)
		}
	}

	log.Println("Staring HttpServer")

}

func (s *HttpServer) AppListening() {

	if err := s.app.Start(s.cfg.App.AppUrl); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to shutdown:%v", err)
	}

}

func (s *HttpServer) Start(pctx context.Context) {

	s.app.Use(middleware.Logger())

	close := make(chan os.Signal, 1)
	signal.Notify(close, syscall.SIGINT, syscall.SIGTERM)

	go s.gracefulShutdown(pctx, close)

	s.sharePriceServer()

	s.AppListening()

}
