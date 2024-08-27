package usecases

import (
	domain "github.com/kidistbezabih/loan-tracker-api/Domain"
	infrastructure "github.com/kidistbezabih/loan-tracker-api/Infrastructure"
)

type UserUsecases struct {
	authusecases  domain.AuthUsecases
	emailservices infrastructure.EmailService
}
