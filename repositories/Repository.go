package repositories

import (
	"github.com/teukumulya-ichsan/go-loan/models"
)

// LoanRepository Interface
type LoanRepository interface {
	Save(loan *models.Loan) (*models.Loan, error)
	FindAll() ([]models.Loan, error)
}


