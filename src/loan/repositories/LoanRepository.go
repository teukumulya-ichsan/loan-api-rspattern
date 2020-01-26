package repositories

import (
	"context"
	"database/sql"

	"github.com/teukumulya-ichsan/go-loan/config"
	"github.com/teukumulya-ichsan/go-loan/src/loan/models"
)

// loanRepositoryPg struct
type loanRepositoryPg struct {
	db *sql.DB
}

// NewLoanRepositoryPg constructor
func NewLoanRepositoryPg() *loanRepositoryPg {

	// ConnectDB PostgresSql
	dbCon, _ := config.ConnectDB()

	return &loanRepositoryPg{
		dbCon.SQL,
	}
}

func (r *loanRepositoryPg) FindAll() (models.Loans, error) {
	query := `SELECT * FROM "pinjaman"`

	var loans models.Loans

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var loan models.Loan

		err = rows.Scan(&loan.ID, &loan.Name, &loan.DateLoan, &loan.Gender, &loan.Ktp, &loan.BirthDate, &loan.Amount, &loan.Period)

		if err != nil {
			return nil, err
		}

		loans = append(loans, loan)

	}
	return loans, nil
}

func (r *loanRepositoryPg) FindByID(id string) (*models.Loan, error) {
	query := `SELECT * FROM "pinjaman" WHERE "id" = $1`

	// variable utk menampung hasil return
	var loan models.Loan

	statement, err := r.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	err = statement.QueryRow(id).Scan(
		&loan.ID,
		&loan.Name,
		&loan.DateLoan,
		&loan.Gender,
		&loan.Ktp,
		&loan.BirthDate,
		&loan.Amount,
		&loan.Period,
	)

	defer statement.Close()

	if err != nil {
		return nil, err
	}

	return &loan, nil
}

func (r *loanRepositoryPg) CreateLoan(ctx context.Context, l *models.Loan) (int64, error) {

	query := `INSERT INTO "pinjaman"("name","date_loan", "gender", "ktp", "birthdate", "amount", "period")
	VALUES($1, $2,$3,$4,$5,$6,$7)`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.ExecContext(ctx,
		l.Name,
		l.DateLoan,
		l.Gender,
		l.Ktp,
		l.BirthDate,
		l.Amount,
		l.Period,
	)
	defer stmt.Close()

	if err != nil {
		return -1, err
	}

	return res.LastInsertId()

}
