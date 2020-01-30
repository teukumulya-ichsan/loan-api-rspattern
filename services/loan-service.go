package services

import (
	"errors"
	"strconv"

	"github.com/teukumulya-ichsan/loan-api-rspattern/models"
)

// LoanService inteface ...
type LoanService interface {
	Validate(loan *models.Loan) error
}

type service struct{}

// NewLoanServices Constructor ...
func NewLoanServices() LoanService {
	return &service{}
}

func (*service) Validate(loan *models.Loan) error {
	isNotNull := loan.IsNotEmpty()
	if isNotNull != nil {
		return isNotNull
	}

	ktpLengthValid := loan.IsValidLengthKtp()
	if ktpLengthValid != nil {
		return ktpLengthValid
	}

	ktpIsNumber := loan.KtpIsNumber()
	if ktpIsNumber != nil {
		return ktpIsNumber
	}

	validProvince := verifyProvince(loan)
	if validProvince == false {
		return errors.New("Provinsi dari KTP tidak di izinkan")
	}

	return nil
}

func verifyProvince(loan *models.Loan) bool {
	sliceProv, _ := strconv.Atoi(loan.Ktp[0:2])

	var provPermitted = map[int]string{
		12: "Sumatera Utara",
		31: "DKI Jakarta",
		32: "Jawa Barat",
		35: "Jawa Timur",
	}

	_, isExist := provPermitted[sliceProv]

	if isExist {
		return true
	}
	return false
}
