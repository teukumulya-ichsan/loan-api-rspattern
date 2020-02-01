package repositories

import (
	"github.com/teukumulya-ichsan/loan-api-rspattern/models"
)

// LoanRepository Interface
type LoanRepository interface {
	Save(loan *models.Loan) (*models.Loan, error)
	FindAll() ([]models.Loan, error)
	GetLast7days(date string) ([]models.Loan, error)
}
