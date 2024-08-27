package routers

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/kidistbezabih/loan-tracker-api/Delivery/Controllers"
	infrastructure "github.com/kidistbezabih/loan-tracker-api/Infrastructure"
)

func SetUpRouter(r *gin.Engine, userController *controllers.UserController) {
	v1 := r.Group("/v1")
	authV1 := v1.Group("/users")
	SetUpAuthRouter(authV1, userController)
}
func SetUpAuthRouter(r *gin.RouterGroup, userController *controllers.UserController) {
	r.POST("/register", userController.RegisterUser)
	r.POST("/login", userController.Login)
	r.PUT("/profile", infrastructure.AuthMiddleware(), userController.GeteProfile)
	r.POST("/verify/:userID/:token", userController.ActivateUser)
	r.POST("/forget-password", userController.ForgetPassword)
	r.POST("/reset", infrastructure.AdminMidleware(), userController.ResetPassword)
}
