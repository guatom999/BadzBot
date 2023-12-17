package botinfoUsecases

import (
	"context"
	"fmt"

	"github.com/guatom999/BadzBot/modules/botinfo"
	"github.com/guatom999/BadzBot/modules/botinfo/botinfoRepositories"
	sharePb "github.com/guatom999/BadzBot/modules/sharePricePb"
)

type IBotinfoUsecase interface {
	Feature(message string) string
	JetTest(message string) string
	GetSharePrice(pctx context.Context, target string) (*botinfo.BotInfoSharePriceRes, error)
}

type botinfoUsecase struct {
	botinfoRepo botinfoRepositories.IBotRepositoryService
}

func NewBotinfoUsecase() IBotinfoUsecase {
	return &botinfoUsecase{}
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

func (u *botinfoUsecase) GetSharePrice(pctx context.Context, target string) (*botinfo.BotInfoSharePriceRes, error) {

	shareResult, err := u.botinfoRepo.GetSharePrice(pctx, &sharePb.SharePriceReq{
		ShareSymbol: target,
	})
	if err != nil {
		return nil, err
	}

	share := new(botinfo.BotInfoSharePriceRes)
	share = &botinfo.BotInfoSharePriceRes{
		Name:  shareResult.Name,
		Price: share.Price,
	}

	return share, nil
}
