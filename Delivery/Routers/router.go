package routers

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/kidistbezabih/loan-tracker-api/Delivery/Controllers"
	infrastructure "github.com/kidistbezabih/loan-tracker-api/Infrastructure"
)

func SetUpRouter(r *gin.Engine, userController *controllers.UserController, loanController *controllers.LoanController) {
	v1 := r.Group("/v1")
	authV1 := v1.Group("/users")
	loanV1 := v1.Group("/loans")
	SetUpAuthRouter(authV1, userController)
	SetUpLoanRouter(loanV1, loanController)

}
func SetUpAuthRouter(r *gin.RouterGroup, userController *controllers.UserController) {
	r.POST("/register", userController.RegisterUser)
	r.POST("/login", userController.Login)
	r.GET("/profile", infrastructure.AuthMiddleware(), userController.GeteProfile)
	r.GET("/activate/:userID/:token", userController.ActivateUser)
	r.GET("/forget-password", userController.ForgetPassword)
	r.PUT("/reset/:userid/:tokentime/:token", userController.ResetPassword)
	r.GET("/all-users", infrastructure.AuthMiddleware(), infrastructure.AdminMidleware(), userController.GetUsers)
	r.GET("/delete/:id", infrastructure.AuthMiddleware(), infrastructure.AdminMidleware(), userController.DeleteUser)
}

func SetUpLoanRouter(r *gin.RouterGroup, loanController *controllers.LoanController) {
	r.POST("/", infrastructure.AuthMiddleware(), loanController.ApplyForLoan)
	r.GET("/loan-status/:loanid", infrastructure.AuthMiddleware(), loanController.ViewLoanStatus)
	r.GET("/all-loans", infrastructure.AuthMiddleware(), infrastructure.AdminMidleware(), loanController.ViewLoans)
	r.PUT("/approve-status/:loanid", infrastructure.AuthMiddleware(), loanController.ApproveLoanStatus)
	r.PUT("/reject-status/:loanid", infrastructure.AuthMiddleware(), loanController.RejectLoanStatus)
	r.DELETE("/delete:id", infrastructure.AdminMidleware(), loanController.DeleteLoan)
}
