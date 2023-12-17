package grpcconn

import (
	"errors"
	"log"

	sharePricePb "github.com/guatom999/BadzBot/modules/sharePricePb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	GrcpClientFactoryHandler interface {
		SharePrice() sharePricePb.SharePriceGrpcServiceClient
	}

	grpcClientFactory struct {
		client *grpc.ClientConn
	}
)

func (g *grpcClientFactory) SharePrice() sharePricePb.SharePriceGrpcServiceClient {
	return sharePricePb.NewSharePriceGrpcServiceClient(g.client)
}

func NewGrpcClient(host string) (GrcpClientFactoryHandler, error) {

	options := make([]grpc.DialOption, 0)

	options = append(options, grpc.WithTransportCredentials(insecure.NewCredentials()))

	clientConn, err := grpc.Dial(host, options...)

	if err != nil {
		log.Printf("Error: Grpc client connection failed:%v", err)
		return nil, errors.New("error: grpc client connection failed:%")
	}

	// defer clientConn.Close()

	return &grpcClientFactory{client: clientConn}, nil
}
