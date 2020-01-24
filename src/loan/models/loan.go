package models

// Loan Struct ...
type Loan struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	DateLoan  string  `json:"date_loan"`
	Gender    string  `json:"gender"`
	Ktp       string  `json:"nik"`
	BirthDate string  `json:"birth_date"`
	Amount    float64 `json:"amount"`
	Period    int     `json:"period"`
}

// Loans Collection ...
type Loans []Loan
