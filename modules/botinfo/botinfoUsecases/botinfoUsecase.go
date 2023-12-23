package botinfoUsecases

import (
	"context"
	"fmt"

	"github.com/guatom999/BadzBot/modules/botinfo/botinfoRepositories"
	sharePb "github.com/guatom999/BadzBot/modules/sharePricePb"
)

type IBotinfoUsecase interface {
	Feature(message string) string
	JetTest(message string) string
	GetSharePrice(pctx context.Context, target string) (*sharePb.SharePriceRes, error)
}

type botinfoUsecase struct {
	botinfoRepo botinfoRepositories.IBotRepositoryService
}

func NewBotinfoUsecase(botinfoRepo botinfoRepositories.IBotRepositoryService) IBotinfoUsecase {
	return &botinfoUsecase{botinfoRepo: botinfoRepo}
}

func (u *botinfoUsecase) Feature(message string) string {
	return fmt.Sprintf("`Test Said: %v`", message)
}

func (u *botinfoUsecase) Forecast() string {
	// return fmt.Sprintf("`not avaliable now i sus: %v`", message)
	return fmt.Sprintln("`not avaliable right now`")
}

func (u *botinfoUsecase) JetTest(message string) string {
	return fmt.Sprintf("`Pen Kuay Rai: %v`", message)
}

func (u *botinfoUsecase) GetSharePrice(pctx context.Context, target string) (*sharePb.SharePriceRes, error) {

	result, err := u.botinfoRepo.GetSharePrice(pctx, target)
	if err != nil {
		return nil, err
	}

	return result, nil
}
