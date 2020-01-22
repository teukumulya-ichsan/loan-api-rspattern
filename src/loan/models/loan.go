package models

import "time"

type Loan struct {
	ID        string
	Name      string
	DateLoan  string
	Gender    string
	Ktp       string
	Birthdate time.Time
	Amount    float64
	Period    int
}

type Loans []Loan
