package main

import (
	"fmt"

	"github.com/teukumulya-ichsan/go-loan/config"
	"github.com/teukumulya-ichsan/go-loan/src/loan/models"
	"github.com/teukumulya-ichsan/go-loan/src/loan/repository"
)

func main() {

	db, err := config.ConnectDB()

	if err != nil {
		fmt.Println(err)
	}

	loanRepo := repository.NewLoanRepositoryPg(db)
	loans, err := getAll(loanRepo)

	if err != nil {
		fmt.Print(err)
	}

	for _, v := range loans {
		fmt.Println(v)
	}

}

func getAll(repo repository.LoanRepository) (models.Loans, error) {
	loans, err := repo.FindAll()

	if err != nil {
		return nil, err
	}

	return loans, nil
}
