package botinfoRepositories

import "go.mongodb.org/mongo-driver/mongo"

type (
	IBotRepositoryService interface {
	}

	botrepository struct {
		db *mongo.Client
	}
)

func MewBotRepository(db *mongo.Client) IBotRepositoryService {
	return &botrepository{
		db: db,
	}
}

func GetData() (string, error) {
	return "", nil
}
