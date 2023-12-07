package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type IConfig interface {
}

type Config struct {
	App App
	Ig  Ig
	Db  Db
}

type App struct {
	Token string
}

type Ig struct {
	Username string
	Password string
}

type Db struct {
	Url string
}

func NewConfig(path string) Config {
	if err := godotenv.Load(path); err != nil {
		log.Fatal("Error loading env :")
		panic(err)
	}

	return Config{
		App: App{
			Token: os.Getenv("APP_TOKEN"),
		},
		Ig: Ig{
			Username: os.Getenv("IG_USERNAME"),
			Password: os.Getenv("IG_PASSWORD"),
		},
		Db: Db{
			Url: os.Getenv("DB_URL"),
		},
	}
}
