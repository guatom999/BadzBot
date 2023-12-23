package database

import (
	"context"
	"log"
	"time"

	"github.com/guatom999/BadzBot/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func DbConnect(pctx context.Context, cfg *config.Config) *mongo.Client {
	ctx, _ := context.WithTimeout(pctx, time.Second*15)
	// defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Db.Url))
	if err != nil {
		log.Fatalf("Error: Connect to database error : %v", err)
		panic(err)
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("Error: Connect to database error : %v", err)
		panic(err)
	}

	return client

}
