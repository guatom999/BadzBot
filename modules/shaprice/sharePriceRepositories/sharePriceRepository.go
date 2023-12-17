package sharePriceRepositories

import (
	"context"
	"errors"
	"log"
	"time"

	sharePb "github.com/guatom999/BadzBot/modules/sharePricePb"
	"github.com/guatom999/BadzBot/pkg/grpcconn"
	"go.mongodb.org/mongo-driver/mongo"
)

type ISharePriceRepository interface {
	Test() (string, error)
}

type sharePriceRepository struct {
	db *mongo.Client
}

func NewSharePriceRepository(db *mongo.Client) ISharePriceRepository {
	return &sharePriceRepository{db: db}
}

func (r *sharePriceRepository) Test() (string, error) {

	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	// defer cancel()

	conn, err := grpcconn.NewGrpcClient("0.0.0.0:1425")
	if err != nil {
		log.Printf("Error: failed to connect Grpc : %v", err)
		return "", errors.New("error: cannot connect grpc")
	}

	result, err := conn.SharePrice().JustTest(ctx, &sharePb.Test{})
	if err != nil {
		log.Printf("Error: Grpc JustTest Failed: %v", err)
		return "", errors.New("error: cannot connect grpc")
	}

	log.Printf("result is ====> %s", result.Message)
	return result.Message, nil
}
