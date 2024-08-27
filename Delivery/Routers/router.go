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
	r.PUT("/profile", infrastructure.AuthMiddleware(), userController.GeteProfile)
	r.POST("/verify/:userID/:token", userController.ActivateUser)
	r.POST("/forget-password", userController.ForgetPassword)
	r.POST("/reset", infrastructure.AdminMidleware(), userController.ResetPassword)
}
func SetUpLoanRouter(r *gin.RouterGroup, loanController *controllers.LoanController) {
	r.POST("/", loanController.ApplyForLoan)
	r.POST("/my-status", loanController.ViewLoanStatus)
	r.PUT("/all-loans", infrastructure.AuthMiddleware(), loanController.ViewLoans)
	r.POST("/approve-status", infrastructure.AuthMiddleware(), loanController.ApproveLoanStatus)
	r.POST("/reject-status", infrastructure.AuthMiddleware(), loanController.RejectLoanStatus)
	r.POST("/delete:id", infrastructure.AdminMidleware(), loanController.DeleteLoan)
}
