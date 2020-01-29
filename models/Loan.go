package models

import "errors"

// Loan type
type Loan struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	DateLoan  string `json:"date_loan"`
	Gender    string `json:"gender"`
	Ktp       string `json:"nik"`
	BirthDate string `json:"birth_date"`
	Amount    int64  `json:"amount"`
	Period    int64  `json:"period"`
}

// IsValidLengthKtp method to check is KTP fixed on 16 character
func (l *Loan) IsValidLengthKtp() error {
	if len(l.Ktp) > 16 {
		return errors.New("The Length KTP is 16")
	}
	return nil
}
