package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	domain "github.com/kidistbezabih/loan-tracker-api/Domain"
)

type UserController struct {
	authuserusecase domain.AuthServices
}

func NewUserController(authServices domain.AuthServices) *UserController {
	return &UserController{
		authuserusecase: authServices,
	}

}

func (uc *UserController) Login(ctx *gin.Context) {
	var userInfo domain.LoginForm
	if err := ctx.ShouldBindJSON(&userInfo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	refToken, accessToken, err := uc.authuserusecase.Login(ctx, userInfo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"access_token": accessToken, "refresh_token": refToken})
}

func (uc *UserController) RegisterUser(ctx *gin.Context) {
	var user domain.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.authuserusecase.RegisterUser(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "successfully registered"})
}

func (uc *UserController) GeteProfile(ctx *gin.Context) {
	userid := ctx.Value("userID").(string)

	userprofile, err := uc.authuserusecase.GetProfile(ctx, userid)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": userprofile})
}

func (uc *UserController) ActivateUser(ctx *gin.Context) {
	token := ctx.Param("token")
	userID := ctx.Param("userID")

	err := uc.authuserusecase.Activate(ctx, userID, token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "activated"})
}

func (uc *UserController) ForgetPassword(ctx *gin.Context) {
	var email domain.Email
	if err := ctx.ShouldBindJSON(&email); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err := uc.authuserusecase.ForgetPassword(ctx, email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "reset email is sent"})
}

func (uc *UserController) ResetPassword(ctx *gin.Context) {
	userid := ctx.Param("userid")
	token := ctx.Param("token")
	var resetForm domain.ResetForm
	if err := ctx.ShouldBindJSON(&resetForm); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := uc.authuserusecase.ResetPassword(ctx, userid, token, resetForm.Passowrd, resetForm.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "password reseted"})
}

func (uc *UserController) GetUsers(ctx *gin.Context) {
	users, err := uc.authuserusecase.GetUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	err := uc.authuserusecase.DeleteUser(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "successfully deleted"})
}
