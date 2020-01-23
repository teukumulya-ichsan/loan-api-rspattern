package repository

import (
	"github.com/teukumulya-ichsan/go-loan/src/loan/models"
)

type LoanRepository interface {
	FindAll() (models.Loans, error)
	FindById(string) (*models.Loan, error)
}
