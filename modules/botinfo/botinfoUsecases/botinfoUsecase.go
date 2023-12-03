package botinfoUsecases

import (
	"fmt"
	import "github.com/gocolly/colly"

)

type IBotinfoUsecase interface {
	Feature(message string) string
	JetTest(message string) string
	GetFollower(target string) string
}

type botintoUsecase struct {
}

func NewBotinfoUsecase() IBotinfoUsecase {
	return &botintoUsecase{}
}

func (u *botintoUsecase) Feature(message string) string {
	return fmt.Sprintf("`Test Said: %v`", message)
}

func (u *botintoUsecase) Forecast() string {
	// return fmt.Sprintf("`not avaliable now i sus: %v`", message)
	return fmt.Sprintln("`not avaliable right now`")
}

func (u *botintoUsecase) JetTest(message string) string {
	return fmt.Sprintf("`Pen Kuay Rai: %v`", message)
}

func (u *botintoUsecase) GetFollower(target string) string {

	if err := scrapper.GetFollower(target); err != nil {
		return ""
	}
	return "1"
}
