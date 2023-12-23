package sharePriceUseCases

import (
	"context"

	"github.com/guatom999/BadzBot/modules/shaprice/sharePriceRepositories"
	sharePb "github.com/guatom999/BadzBot/modules/sharePricePb"
)

type (
	ISharePriceUseCase interface {
		Test() (string, error)
		SharePrice(pctx context.Context, shareSymbol string) (*sharePb.SharePriceRes, error)
	}

	sharePriceUseCase struct {
		sharePriceRepo sharePriceRepositories.ISharePriceRepository
	}
)

func NewSharePriceUseCase(sharePriceRepo sharePriceRepositories.ISharePriceRepository) ISharePriceUseCase {
	return &sharePriceUseCase{sharePriceRepo: sharePriceRepo}
}

func (u *sharePriceUseCase) Test() (string, error) {

	result, err := u.sharePriceRepo.Test()
	if err != nil {
		return "", err
	}

	return result, nil
}

func (u *sharePriceUseCase) SharePrice(pctx context.Context, shareSymbol string) (*sharePb.SharePriceRes, error) {

	result, err := u.sharePriceRepo.SharePrice(pctx, shareSymbol)
	if err != nil {
		return nil, err
	}

	return result, nil
}
