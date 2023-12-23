package botinfoRepositories

import (
	"context"
	"errors"
	"log"
	"time"

	sharePb "github.com/guatom999/BadzBot/modules/sharePricePb"
	"github.com/guatom999/BadzBot/pkg/grpcconn"
)

type (
	IBotRepositoryService interface {
		GetSharePrice(pctx context.Context, shareSymbol string) (*sharePb.SharePriceRes, error)
	}

	botrepository struct {
		// db *mongo.Client
	}
)

func NewBotRepository() IBotRepositoryService {
	return &botrepository{
		// db: db,
	}
}

func (r *botrepository) GetSharePrice(pctx context.Context, shareSymbol string) (*sharePb.SharePriceRes, error) {
	log.Println("target is ========================:", shareSymbol)
	ctx, cancel := context.WithTimeout(pctx, time.Second*10)
	defer cancel()

	conn, err := grpcconn.NewGrpcClient("0.0.0.0:1425")
	if err != nil {
		log.Printf("Error: failed to connect Grpc : %v", err)
		return nil, errors.New("error: cannot connect grpc")
	}

	result, err := conn.SharePrice().SharePriceSearch(ctx, &sharePb.SharePriceReq{
		ShareSymbol: shareSymbol,
	})

	log.Println("result is :", result)

	if err != nil {
		log.Printf("Error: Grpc JustTest Failed: %v", err)
		return nil, errors.New("error: cannot connect grpc")
	}

	log.Printf("result is ====> %s", result)
	return result, nil
}
