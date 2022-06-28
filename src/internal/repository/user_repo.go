package repository

import (
	"context"
	"fmt"

	"github.com/adiletelf/authentication-go/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepoImpl struct {
	ctx        context.Context
	collection *mongo.Collection
}

func NewUserRepo(ctx context.Context, collection *mongo.Collection, client *mongo.Client) *UserRepoImpl {
	return &UserRepoImpl{
		ctx:        ctx,
		collection: collection,
	}
}

func (ur *UserRepoImpl) Save(u *model.User) (string, error) {
	result, err := ur.collection.InsertOne(ur.ctx, u)
	if err != nil {
		return "", err
	}

	id, ok := result.InsertedID.(primitive.Binary)
	if !ok {
		return "", fmt.Errorf("unable to get insertedId")
	}
	return fmt.Sprintf("%x", id.Data), nil
}
