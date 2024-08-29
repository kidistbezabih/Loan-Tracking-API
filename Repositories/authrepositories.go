package repositories

import (
	"context"

	domain "github.com/kidistbezabih/loan-tracker-api/Domain"
	"github.com/kidistbezabih/loan-tracker-api/errs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthStorage struct {
	AuthTokenImple
	AuthUserImple
}
type AuthUserImple struct {
	usercollection *mongo.Collection
}

type AuthTokenImple struct {
	tokencollection *mongo.Collection
}

func NewAuthStorage(usercollection *mongo.Collection, tokencollection *mongo.Collection) domain.AuthRepository {
	return &AuthStorage{
		AuthTokenImple: AuthTokenImple{
			tokencollection: tokencollection,
		},
		AuthUserImple: AuthUserImple{
			usercollection: usercollection,
		},
	}
}

func (au *AuthUserImple) CreateUser(ctx context.Context, user domain.User) (string, error) {
	user.ID = primitive.NewObjectID().Hex()

	result, err := au.usercollection.InsertOne(ctx, user)
	if err != nil {
		return "", errs.ErrFailToCreateUser
	}
	return result.InsertedID.(string), nil
}

func (au *AuthUserImple) UpdateUser(ctx context.Context, user domain.User) (domain.User, error) {
	_, err := primitive.ObjectIDFromHex(user.ID)

	if err != nil {
		return domain.User{}, errs.ErrIsnvalidID
	}
	filter := bson.D{bson.E{Key: "_id", Value: user.ID}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "name", Value: user.Name},
		{Key: "username", Value: user.Username},
		{Key: "email", Value: user.Email},
		{Key: "isactive", Value: user.IsActive},
		{Key: "isadmin", Value: user.IsAdmin},
		{Key: "createdat", Value: user.CreatedAt},
		{Key: "updatedat", Value: user.UpdatedAt},
	}}}

	result, err := au.usercollection.UpdateOne(ctx, filter, update)

	if err != nil {
		return domain.User{}, err
	}

	if result.MatchedCount != 1 {
		return domain.User{}, errs.ErrNoUserWithId
	}
	return user, nil
}

func (au *AuthUserImple) GetUserByUsername(ctx context.Context, username string) (domain.User, error) {
	var user domain.User

	filter := bson.D{bson.E{Key: "username", Value: username}}
	err := au.usercollection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		return domain.User{}, errs.ErrNoUserWithUsername
	}
	return user, nil
}

func (au *AuthUserImple) GetUserByID(ctx context.Context, id string) (domain.User, error) {
	_, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return domain.User{}, errs.ErrIsnvalidID
	}

	var user domain.User

	filter := bson.D{bson.E{Key: "_id", Value: id}}
	err = au.usercollection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		return domain.User{}, errs.ErrNoUserWithId
	}
	return user, nil
}

func (au *AuthUserImple) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User

	filter := bson.D{bson.E{Key: "email", Value: email}}
	err := au.usercollection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		return domain.User{}, errs.ErrNoUserWithEmail
	}
	return user, nil
}

func (au *AuthUserImple) GetUsers(ctx context.Context) ([]domain.User, error) {
	var users []domain.User

	filter := bson.D{}
	cursor, err := au.usercollection.Find(ctx, filter)

	if err != nil {
		return []domain.User{}, err
	}

	if err := cursor.All(ctx, &users); err != nil {
		return []domain.User{}, err
	}
	return users, nil
}

func (au *AuthUserImple) DeleteUser(ctx context.Context, id string) error {
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	result, err := au.usercollection.DeleteOne(ctx, filter)

	if err != nil {
		return errs.ErrFailToDelete
	}
	if result.DeletedCount == 0 {
		return errs.ErrNoUserWithId
	}
	return nil
}

func (at *AuthTokenImple) RegisterRefreshToken(ctx context.Context, userId string, token string) error {
	savedToken := domain.Token{
		UserId:       userId,
		RefreshToken: token,
	}
	savedToken.ID = primitive.NewObjectID().Hex()

	_, err := at.tokencollection.InsertOne(ctx, savedToken)
	return err
}

func (at *AuthTokenImple) GetRefreshToken(ctx context.Context, userId string) (string, error) {
	var token domain.Token

	filter := bson.D{bson.E{Key: "userid", Value: userId}}
	err := at.tokencollection.FindOne(ctx, filter).Decode(&token)

	if err != nil {
		return "", err
	}
	return token.RefreshToken, nil
}

func (at *AuthUserImple) GetCollectionCount(ctx context.Context) (int64, error) {
	count, err := at.usercollection.CountDocuments(ctx, bson.D{})
	if err != nil {
		return 0, err
	}

	return count, nil
}
