package sharePriceUseCases

import "github.com/guatom999/BadzBot/modules/shaprice/sharePriceRepositories"

type (
	ISharePriceUseCase interface {
		Test() (string, error)
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
