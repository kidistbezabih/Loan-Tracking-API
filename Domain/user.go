package domain

import (
	"context"
	"time"
)

type User struct {
	ID        string
	Username  string
	Name      string
	Email     string
	Passoword string
	IsActive  bool
	IsAdmin   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Token struct {
	ID          string
	UserID      string
	TokenString string
}

type LoginForm struct {
	Email    string
	Password string
}

type AuthServices interface {
	CreateUser(ctx context.Context, user User) (string, error)
	GetUserWithID(ctx context.Context, id string) (User, error)
	GetUserWithEmail(ctx context.Context, email string) (User, error)
	GetUserWithUsername(ctx context.Context, usename string) (User, error)
	RegisterRefreshToken(ctx context.Context, userid, token string) error
}

type AuthUsecases interface {
	UserLogin(ctx context.Context, loginform LoginForm) (string, string, error)
	RegisterUser(ctx context.Context, user User) error
	EmailVerification(ctx context.Context, id, token string) error
	GenerateActivationToken(hashedpassword string, updatedat time.Time) string
	GenerateToken(user User, tokentype string) (string, error)
	GetUserProfile()
}
