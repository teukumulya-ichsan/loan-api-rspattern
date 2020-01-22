package controller

import (
	"encoding/json"
	"net/http"

	"github.com/teukumulya-ichsan/go-loan/config"
	"github.com/teukumulya-ichsan/go-loan/src/loan/repository"
)

// NewPostHandler ...
func NewLoanHandler(db *config.DB) *Loan {
	return &Loan{
		repo: repository.NewLoanRepositoryPg(db.SQL),
	}
}

type Loan struct {
	repo repository.LoanRepository
}

func (p *Loan) Fetch(w http.ResponseWriter, r *http.Request) {
	payload, _ := p.repo.FindAll()

	respondwithJSON(w, http.StatusOK, payload)
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
