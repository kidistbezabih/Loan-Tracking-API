package domain

import (
	"context"
	"time"
)

type Loan struct {
	ID        string    `json:"loanid,omitempty" bson:"_id"`
	UserId    string    `json:"userid"`
	Amount    int64     `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}

type LoanApplication struct {
	Amount int64 `json:"amount"`
}

type LoanRepository interface {
	CreateLoan(ctx context.Context, loan Loan) error
	FindLoanById(ctx context.Context, id string) (Loan, error)
	FindLoans(ctx context.Context, userid string) ([]Loan, error)
	UpdateLoanStatus(ctx context.Context, id string, status string) error
	DeleteLoan(ctx context.Context, id string) error
}
type LoanServices interface {
	ApplyForLoan(ctx context.Context, loanform LoanApplication, userid string) error
	ViewLoanStatus(ctx context.Context, id string) (string, error)
	ApproveLoanStatus(ctx context.Context, id string) error
	RejectLoanStatus(ctx context.Context, id string) error
	DeleteLoan(ctx context.Context, id string) error
	ViewLoans(ctx context.Context, userid string) ([]Loan, error)
}
