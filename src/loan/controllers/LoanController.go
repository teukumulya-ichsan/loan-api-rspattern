package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/teukumulya-ichsan/go-loan/config"
	"github.com/teukumulya-ichsan/go-loan/src/loan/models"
	"github.com/teukumulya-ichsan/go-loan/src/loan/repositories"
)

// Loan struct to repo
type Loan struct {
	repo repositories.LoanRepository
}

// NewLoanController ...
func NewLoanController(db *config.DB) *Loan {
	return &Loan{
		repo: repositories.NewLoanRepositoryPg(db.SQL),
	}
}

// FindAll request to get all loan
func (p *Loan) FindAll(w http.ResponseWriter, r *http.Request) {
	payload, _ := p.repo.FindAll()

	respondwithJSON(w, http.StatusOK, payload)
}

// FindByID request to Get loan by ID
func (p *Loan) FindByID(w http.ResponseWriter, r *http.Request) {
	payload, _ := p.repo.FindByID(chi.URLParam(r, "id"))

	respondwithJSON(w, http.StatusOK, payload)
}

// CreateLoan request to create loan
func (p *Loan) CreateLoan(w http.ResponseWriter, r *http.Request) {
	loan := models.Loan{}

	err := json.NewDecoder(r.Body).Decode(&loan)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Bad Request")
	} else {
		_, err = p.repo.CreateLoan(r.Context(), &loan)
		respondwithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
	}

}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {

	response, _ := json.Marshal(msg)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
