package configs

import (
	"log"

	"github.com/joho/godotenv"
)

type IConfig interface {
	App() IAppConfig
	Ig() IIgConfig
}

type config struct {
	app *app
	ig  *ig
}

func NewConfig(path string) IConfig {
	envMap, err := godotenv.Read(path)
	if err != nil {
		log.Fatal("read .env file failed")
	}
	return &config{
		app: &app{
			token: envMap["APP_TOKEN"],
		},
		ig: &ig{
			username: envMap["IG_USERNAME"],
			password: envMap["IG_PASSWORD"],
		},
	}
}

type IAppConfig interface {
	GetToken() string
}

type app struct {
	token string
}

func (c config) App() IAppConfig {
	return c.app
}

func (a app) GetToken() string {
	return a.token
}

type IIgConfig interface {
	GetIgUserName() string
	GetIgPassword() string
}
type ig struct {
	username string
	password string
}

func (c *config) Ig() IIgConfig {
	return c.ig
}

func (i *ig) GetIgUserName() string {
	return i.username
}

func (i *ig) GetIgPassword() string {
	return i.password
}
