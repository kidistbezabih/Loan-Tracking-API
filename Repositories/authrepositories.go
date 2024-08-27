package repositories

import (
	"context"

	domain "github.com/kidistbezabih/loan-tracker-api/Domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepositories struct {
	userCollection *mongo.Collection
	tokenColletion *mongo.Collection
}

func NewAuthService(userCollection *mongo.Collection) domain.AuthServices {
	return &AuthRepositories{
		userCollection: userCollection,
	}
}

func (ar *AuthRepositories) CreateUser(ctx context.Context, user domain.User) (string, error) {
	user.ID = primitive.NewObjectID().Hex()

	result, err := ar.userCollection.InsertOne(ctx, user)
	if err != nil {
		return "", ErrFailToCreateUser
	}
	return result.InsertedID.(string), nil
}

func (ar *AuthRepositories) GetUserWithID(ctx context.Context, id string) (domain.User, error) {
	var user domain.User
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	err := ar.userCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (ar *AuthRepositories) GetUserWithEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User
	filter := bson.D{bson.E{Key: "email", Value: email}}
	err := ar.userCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return domain.User{}, ErrNoUesrWitThisEmail
	}
	return user, nil
}

func (ar *AuthRepositories) GetUserWithUsername(ctx context.Context, username string) (domain.User, error) {
	var user domain.User
	filter := bson.D{bson.E{Key: "user_name", Value: username}}
	err := ar.userCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return domain.User{}, ErrNoUesrWitThisUsername
	}
	return user, nil
}

func (ar *AuthRepositories) RegisterRefreshToken(ctx context.Context, userid, tokenstring string) error {
	var token domain.Token

	token.ID = primitive.NewObjectID().Hex()
	token.UserID = userid
	token.TokenString = tokenstring

	_, err := ar.tokenColletion.InsertOne(ctx, token)
	if err != nil {
		return err
	}
	return err
}
