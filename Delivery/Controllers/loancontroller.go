package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	domain "github.com/kidistbezabih/loan-tracker-api/Domain"
)

type LoanController struct {
	loanservices domain.LoanServices
}

func NewLoanController(loanservices domain.LoanServices) *LoanController {
	return &LoanController{
		loanservices: loanservices,
	}

}

func (lc *LoanController) ApplyForLoan(ctx *gin.Context) {
	var loanform domain.LoanApplication
	userid, ok := ctx.Value("user_id").(string)

	if !ok || userid == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user ID not found or invalid"})
		return
	}

	if err := ctx.ShouldBindJSON(&loanform); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := lc.loanservices.ApplyForLoan(ctx, loanform, userid)
	if err != nil {
		ctx.JSON(http.StatusFound, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "application successfully submited"})
}

func (lc *LoanController) ViewLoanStatus(ctx *gin.Context) {
	loanid := ctx.Param("loanid")

	userprofile, err := lc.loanservices.ViewLoanStatus(ctx, loanid)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"the user status is": userprofile})
}

func (lc *LoanController) ApproveLoanStatus(ctx *gin.Context) {
	loanid := ctx.Param("loanid")

	err := lc.loanservices.ApproveLoanStatus(ctx, loanid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "approved"})
}

func (lc *LoanController) RejectLoanStatus(ctx *gin.Context) {
	loanid := ctx.Param("loanid")

	err := lc.loanservices.RejectLoanStatus(ctx, loanid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "rejected"})
}

func (lc *LoanController) DeleteLoan(ctx *gin.Context) {
	loanid := ctx.Param("loanid")

	err := lc.loanservices.DeleteLoan(ctx, loanid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "successfully deleted"})
}

func (lc *LoanController) ViewLoans(ctx *gin.Context) {
	userid, ok := ctx.Value("user_id").(string)
	if !ok || userid == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user ID not found or invalid"})
		return
	}
	loans, err := lc.loanservices.ViewLoans(ctx, userid)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}
	ctx.JSON(http.StatusOK, gin.H{"loan list": loans})
}
