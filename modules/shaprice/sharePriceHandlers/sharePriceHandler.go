package sharePriceHandlers

import (
	"github.com/guatom999/BadzBot/modules/shaprice/sharePriceUseCases"
	"github.com/labstack/echo/v4"
)

type (
	ISharePriceHandler interface {
		Test(e echo.Context) error
	}

	sharePriceHandler struct {
		sharePriceUseCase sharePriceUseCases.ISharePriceUseCase
	}
)

func NewSharePriceHandler(sharePriceUseCase sharePriceUseCases.ISharePriceUseCase) ISharePriceHandler {
	return &sharePriceHandler{sharePriceUseCase: sharePriceUseCase}
}

func (h *sharePriceHandler) Test(e echo.Context) error {

	result, err := h.sharePriceUseCase.Test()
	if err != nil {
		return e.JSON(400, "error: failed to test grpc")
	}

	return e.JSON(200, result)
}
