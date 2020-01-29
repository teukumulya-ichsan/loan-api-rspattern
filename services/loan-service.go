package services

import (
	"github.com/teukumulya-ichsan/go-loan/models"
)

// LoanService inteface ...
type LoanService interface {
	Validate(loan *models.Loan) error
}

type service struct {
	models.Loan
}

// NewLoanServices Constructor ...
func NewLoanServices() LoanService {
	return &service{}
}

func (m *service) Validate(loan *models.Loan) error {
	err := loan.IsValidLengthKtp()

	err2 := loan.IsNotEmpty()

	return err
}
