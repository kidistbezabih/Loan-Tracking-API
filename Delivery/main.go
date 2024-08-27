package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	controllers "github.com/kidistbezabih/loan-tracker-api/Delivery/Controllers"
	db "github.com/kidistbezabih/loan-tracker-api/Delivery/Internal"
	routers "github.com/kidistbezabih/loan-tracker-api/Delivery/Routers"
	infrastructure "github.com/kidistbezabih/loan-tracker-api/Infrastructure"
	repositories "github.com/kidistbezabih/loan-tracker-api/Repositories"
	auth "github.com/kidistbezabih/loan-tracker-api/Usecases"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getHost() string {
	hostUrl := os.Getenv("HOST_URL")
	if hostUrl != "" {
		return hostUrl
	}

	return "localhost:8000"
}

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	mongoClient := db.NewMongoClient()
	mongoDB := mongoClient.Database(os.Getenv("MONGO_DB"))

	authRepository := repositories.NewAuthStorage(mongoDB.Collection("users"), mongoDB.Collection("tokens"))
	loanrepo := repositories.NewLoanRepoImple(mongoDB.Collection("loans"))

	authUsecase := auth.NewAuthUserUsecase(authRepository, infrastructure.NewEmail(
		os.Getenv("EMAIL_USERNAME"),
		os.Getenv("EMAIL_PASSWORD"),
		os.Getenv("SMTP_HOST"),
		os.Getenv("SMTP_PORT"),
	))
	loanUsecase := auth.NewLoanUsecases(loanrepo)

	authController := controllers.NewUserController(authUsecase)
	loanController := controllers.NewLoanController(loanUsecase)

	routers.SetUpRouter(r, authController, loanController)

	r.Run(getHost())
}
