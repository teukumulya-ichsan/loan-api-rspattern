package repository

import (
	"database/sql"
	"github.com/teukumulya-ichsan/go-loan/src/loan/models"
)

type loanRepositoryPg struct {
	db *sql.DB
}

func NewLoanRepositoryPg(db *sql.DB) *loanRepositoryPg {
	return &loanRepositoryPg{db}
}

func (r *loanRepositoryPg) FindAll() (models.Loans, error) {
	query := `SELECT * FROM "loans"`

	var loans models.Loans

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var loan models.Loan

		err = rows.Scan(&loan.ID, &loan.Name, &loan.DateLoan, &loan.Gender, &loan.Ktp, &loan.Birthdate, &loan.Amount, &loan.Period)

		if err != nil {
			return nil, err
		}

		loans = append(loans, loan)

	}
	return loans, nil
}
