package repository

import (
	"context"
	"fmt"

	"github.com/adiletelf/authentication-go/internal/model"
	"github.com/adiletelf/authentication-go/internal/token"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
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
	ur.BeforeSave(u)
	result, err := ur.collection.InsertOne(ur.ctx, u)
	if err != nil {
		return "", err
	}

	uuid, ok := result.InsertedID.(primitive.Binary)
	if !ok {
		return "", fmt.Errorf("error while inserting %v", u)
	}
	return fmt.Sprintf("%x", uuid.Data), nil
}

func (ur *UserRepoImpl) BeforeSave(u *model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (ur *UserRepoImpl) LoginCheck(uuid uuid.UUID, password string) (model.TokenDetails, error) {
	var err error
	var user model.User

	filter := bson.D{{Key: "_id", Value: uuid}}
	err = ur.collection.FindOne(ur.ctx, filter).Decode(&user)
	if err != nil {
		return model.TokenDetails{}, err
	}

	err = VerifyPassword(password, user.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return model.TokenDetails{}, err
	}

	tokenDetails, err := token.Generate(user.UUID)
	return tokenDetails, err
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
