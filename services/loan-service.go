package services

import (
	"errors"

	"github.com/teukumulya-ichsan/go-loan/models"
	"github.com/teukumulya-ichsan/go-loan/repositories"
)

// LoanService inteface ...
type LoanService interface {
	Validate(loan *models.Loan) error
	Create(loan *models.Loan) (*models.Loan, error)
	FindAll() ([]models.Loan, error)
}

type service struct {
	repositories.LoanRepository
}

// NewLoanServices Constructor ...
func NewLoanServices(repository repositories.LoanRepository) LoanService {
	return &service{
		repository,
	}
}

func (repo *service) Create(loan *models.Loan) (*models.Loan, error) {
	return repo.Save(loan)
}

func (repo *service) FindAll() ([]models.Loan, error) {
	return repo.FindAll()
}

func (repo *service) Validate(loan *models.Loan) error {
	// validate if loan is null
	if loan == nil {
		err := errors.New("The Loan is Empty")
		return err
	}
	if loan.Name == "" {
		err := errors.New("The Name is Empty")
		return err
	}
	return nil
}
