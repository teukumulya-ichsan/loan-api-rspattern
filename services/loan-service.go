package services

import (
	"errors"
	"math/rand"

	"github.com/teukumulya-ichsan/go-loan/models"
	"github.com/teukumulya-ichsan/go-loan/repositories"
)

// LoanService inteface ...
type LoanService interface {
	Validate(loan *models.Loan) error
	Create(loan *models.Loan) (*models.Loan, error)
	FindAll() ([]models.Loan, error)
}

type service struct{}

var (
	repo repositories.LoanRepository
)

// NewLoanServices Constructor ...
func NewLoanServices(repository repositories.LoanRepository) LoanService {
	repo = repository
	return &service{}
}

func (*service) Validate(loan *models.Loan) error {
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

func (*service) Create(loan *models.Loan) (*models.Loan, error) {
	loan.ID = rand.Int63()
	return repo.Save(loan)
}

func (*service) FindAll() ([]models.Loan, error) {
	return repo.FindAll()
}
