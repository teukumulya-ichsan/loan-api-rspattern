package models

import "time"

type Loan struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	DateLoan  string    `json:"date_loan"`
	Gender    string    `json:"gender"`
	Ktp       string    `json:"nik"`
	BirthDate time.Time `json:"birth_date"`
	Amount    float64   `json:"amount"`
	Period    int       `json:"period"`
}

type Loans []Loan
