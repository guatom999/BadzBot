package sharePriceHandlers

import (
	"context"
	"log"

	"github.com/guatom999/BadzBot/modules/shaprice/sharePriceUseCases"
	"github.com/labstack/echo/v4"
)

type (
	ISharePriceHandler interface {
		Test(e echo.Context) error
		SharePrice(e echo.Context) error
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

func (h *sharePriceHandler) SharePrice(e echo.Context) error {

	ctx := context.Background()

	req := e.Param("sharesymbol")

	log.Printf("sharesymbol req is =============>%v", req)

	result, err := h.sharePriceUseCase.SharePrice(ctx, req)
	if err != nil {
		log.Println("error: error is :", err)
		return e.JSON(400, "error: failed to get sharePrice grpc")
	}

	return e.JSON(200, result)
}
