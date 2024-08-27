package auth

import (
	"context"
	"time"

	domain "github.com/kidistbezabih/loan-tracker-api/Domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoanUsecases struct {
	loanrepositories domain.LoanRepository
}

func (lu *LoanUsecases) ApplyForLoan(ctx context.Context, loanform domain.LoanApplication) error {
	var loan domain.Loan
	loan.ID = primitive.NewObjectID().Hex()
	loan.UserId = loanform.UserId
	loan.Status = "pending"
	loan.CreatedAt = time.Now()
	loan.UpdatedAt = time.Now()

	err := lu.loanrepositories.CreateLoan(ctx, loan)
	if err != nil {
		return err
	}
	return nil
}

func (lu *LoanUsecases) ViewLoanStatus(ctx context.Context, id string) (string, error) {
	loan, err := lu.loanrepositories.FindLoanById(ctx, id)
	if err != nil {
		return "", err
	}
	return loan.Status, nil
}

func (lu *LoanUsecases) ViewLoans(ctx context.Context) ([]domain.Loan, error) {
	loans, err := lu.loanrepositories.FindLoans(ctx)
	if err != nil {
		return []domain.Loan{}, err
	}
	return loans, nil
}

func (lu *LoanUsecases) ApproveLoanStatus(ctx context.Context, id string) error {
	_, err := lu.loanrepositories.FindLoanById(ctx, id)
	if err != nil {
		return err
	}
	status := "approved"
	err = lu.loanrepositories.UpdateLoanStatus(ctx, id, status)
	if err != nil {
		return err
	}
	return nil
}

func (lu *LoanUsecases) RejectLoanStatus(ctx context.Context, id string) error {
	_, err := lu.loanrepositories.FindLoanById(ctx, id)
	if err != nil {
		return err
	}
	status := "rejected"
	err = lu.loanrepositories.UpdateLoanStatus(ctx, id, status)
	if err != nil {
		return err
	}
	return nil
}

func (lu *LoanUsecases) DeleteLoan(ctx context.Context, id string) error {
	err := lu.loanrepositories.DeleteLoan(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
