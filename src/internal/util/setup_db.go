package util

import (
	"context"

	"github.com/adiletelf/authentication-go/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupDB(ctx context.Context, config *config.Config) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(config.DB.ConnectionString)
	client, err := mongo.Connect(ctx, clientOptions)
	return client, err
}
