package domain

import (
	"context"
	"time"
)

type Loan struct {
	ID        string    `json:"id,omitempty" bson:"_id"`
	UserId    string    `json:"user"`
	Amount    int64     `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdat"`
	UpdatedAt time.Time `json:"updatedat"`
}

type LoanApplication struct {
	UserId string `json:"user"`
	Amount int64  `json:"amount"`
}

type LoanRepository interface {
	CreateLoan(ctx context.Context, loan Loan) error
	FindLoanById(ctx context.Context, id string) (Loan, error)
	FindLoans(ctx context.Context) ([]Loan, error)
	UpdateLoanStatus(ctx context.Context, id string, status string) error
	DeleteLoan(ctx context.Context, id string) error
}
type LoanServices interface {
	ApplyForLoan(ctx context.Context, loanform LoanApplication) error
	ViewLoanStatus(ctx context.Context, id string) (string, error)
	ApproveLoanStatus(ctx context.Context, id string) error
	RejectLoanStatus(ctx context.Context, id string) error
	DeleteLoan(ctx context.Context, id string) error
}
