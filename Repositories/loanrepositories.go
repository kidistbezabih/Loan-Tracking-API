package repositories

import (
	"context"

	domain "github.com/kidistbezabih/loan-tracker-api/Domain"
	"github.com/kidistbezabih/loan-tracker-api/errs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LoanReposImple struct {
	loanCollection *mongo.Collection
}

func NewLoanRepoImple(loanCollection *mongo.Collection) domain.LoanRepository {
	return &LoanReposImple{
		loanCollection: loanCollection,
	}
}

func (lr *LoanReposImple) CreateLoan(ctx context.Context, loan domain.Loan) error {
	_, err := lr.loanCollection.InsertOne(ctx, loan)
	if err != nil {
		return err
	}
	return nil
}

func (lr *LoanReposImple) FindLoanById(ctx context.Context, id string) (domain.Loan, error) {
	var loan domain.Loan

	filter := bson.D{bson.E{Key: "_id", Value: id}}
	err := lr.loanCollection.FindOne(ctx, filter).Decode(&loan)

	if err != nil {
		return domain.Loan{}, errs.ErrNoUserWithId
	}
	return loan, nil
}

func (lr *LoanReposImple) FindLoans(ctx context.Context, userid string) ([]domain.Loan, error) {
	var loans []domain.Loan

	filter := bson.D{bson.E{Key: "userid", Value: userid}}
	cursor, err := lr.loanCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx) // Ensure the cursor is closed after the operation

	if err := cursor.All(ctx, &loans); err != nil {
		return nil, err
	}
	return loans, nil
}

func (lr *LoanReposImple) UpdateLoanStatus(ctx context.Context, id string, status string) error {
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "status", Value: status},
	}}}

	result, err := lr.loanCollection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	if result.MatchedCount != 1 {
		return errs.ErrNoUserWithId
	}
	return nil
}

func (lr *LoanReposImple) DeleteLoan(ctx context.Context, id string) error {
	filter := bson.D{bson.E{Key: "id", Value: id}}
	result, err := lr.loanCollection.DeleteOne(ctx, filter)

	if err != nil {
		return errs.ErrFailToDelete
	}
	if result.DeletedCount == 0 {
		return errs.ErrNoUserWithId
	}
	return nil
}
