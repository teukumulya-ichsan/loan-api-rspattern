package repositories

import (
	"context"
	"github.com/teukumulya-ichsan/go-loan/src/loan/models"
)

// LoanRepository interface
type LoanRepository interface {
	FindAll() (models.Loans, error)
	FindByID(string) (*models.Loan, error)
	CreateLoan(ctc context.Context, l *models.Loan) (int64, error)
}
