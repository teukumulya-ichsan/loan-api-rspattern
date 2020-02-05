package models

import (
	"errors"
	"reflect"
	"strconv"
)

// Loan type
type Loan struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`

	DateLoan  string `json:"date_loan"`
	Gender    string `json:"gender"`
	Ktp       string `json:"ktp"`
	BirthDate string `json:"birth_date"`
	Amount    int64  `json:"amount"`
	Period    int64  `json:"period"`
}

// IsValidLengthKtp method to check is KTP fixed on 16 character
func (l *Loan) IsValidLengthKtp() error {
	if len(l.Ktp) > 16 || len(l.Ktp) < 16 {
		return errors.New("the Length KTP it must 16 character")
	}
	return nil
}

// IsNotEmpty method to check if the input is not null
func (l *Loan) IsNotEmpty() error {
	var reflectValue = reflect.ValueOf(l)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}
	var reflectType = reflectValue.Type()
	for i := 0; i < reflectValue.NumField(); i++ {
		if reflectValue.Field(i).Interface() == "" {
			msg := "The " + reflectType.Field(i).Name + " is cannot be null."

			return errors.New(msg)
		}
	}

	return nil
}

// KtpIsNumber method to check if the ktp number must be number
func (l *Loan) KtpIsNumber() error {
	if _, err := strconv.Atoi(l.Ktp); err != nil {
		return errors.New("KTP must the number")
	}
	return nil
}
