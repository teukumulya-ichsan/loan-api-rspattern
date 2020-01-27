package repositories

import (
	"context"
	"database/sql"
	"log"

	"github.com/teukumulya-ichsan/go-loan/config"
	"github.com/teukumulya-ichsan/go-loan/models"
)

type repoPQ struct {
	db *sql.DB
}

// NewPostgresRepository Constructor
func NewPostgresRepository() LoanRepository {

	// ConnectDB PostgresSql
	dbCon, _ := config.ConnectPG()

	return &repoPQ{
		dbCon.SQL,
	}
}

// const (
// 	tableName string = "loans"
// )

func (r *repoPQ) FindAll() ([]models.Loan, error) {
	ctx := context.Background()
	query := `SELECT * FROM "loans"`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		log.Fatalf("Failed to Connect to %v", err)
		return nil, err
	}
	defer rows.Close()

	var loans []models.Loan
	for rows.Next() {
		var loan models.Loan
		err = rows.Scan(
			&loan.ID,
			&loan.Name,
			&loan.DateLoan,
			&loan.Gender,
			&loan.Ktp,
			&loan.BirthDate,
			&loan.Amount,
			&loan.Period,
		)
		if err != nil {
			log.Fatalf("Failed to iterate the list of loan: %v", err)
			return nil, err
		}

		loans = append(loans, loan)
	}
	return loans, nil

}

func (r *repoPQ) Save(loan *models.Loan) (*models.Loan, error) {
	ctx := context.Background()

	query := `INSERT INTO "loans"("name","date_loan", "gender", "ktp", "birthdate", "amount", "period")
	VALUES($1, $2,$3,$4,$5,$6,$7)`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatalf("Failed to Connect a Pg: %v", err)

		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		loan.Name,
		loan.DateLoan,
		loan.Gender,
		loan.Ktp,
		loan.BirthDate,
		loan.Amount,
		loan.Period,
	)

	if err != nil {
		log.Fatalf("Failed to Add New Loan: %v", err)

		return nil, err
	}

	return loan, nil
}
