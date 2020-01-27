package models

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
