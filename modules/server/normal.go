package server

import (
	"github.com/guatom999/BadzBot/modules/shaprice/sharePriceHandlers"
	"github.com/guatom999/BadzBot/modules/shaprice/sharePriceRepositories"
	"github.com/guatom999/BadzBot/modules/shaprice/sharePriceUseCases"
)

func (s *HttpServer) sharePriceServer() {
	sharePriceRepository := sharePriceRepositories.NewSharePriceRepository(s.db)
	sharePriceUseCase := sharePriceUseCases.NewSharePriceUseCase(sharePriceRepository)
	sharePriceHandler := sharePriceHandlers.NewSharePriceHandler(sharePriceUseCase)

	router := s.app.Group("/share_price_test")

	router.GET("/test", sharePriceHandler.Test)
	router.GET("/shareprice/:sharesymbol", sharePriceHandler.SharePrice)

}
