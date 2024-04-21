package mongodb

import (
	"context"
	"os"

	"github.com/mfcbentes/meu-primeiro-crud-go/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGODB_URL     = "MONGODB_URL"
	MONGODB_USER_DB = "MONGODB_USER_DB"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	logger.Info("Init NewMongoDBConnection")
	mongodb_uri := os.Getenv(MONGODB_URL)
	mongodb_database := os.Getenv(MONGODB_USER_DB)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodb_uri))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database(mongodb_database), nil
}
