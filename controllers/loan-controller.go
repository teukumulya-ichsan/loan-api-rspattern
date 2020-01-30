package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/teukumulya-ichsan/go-loan/errors"
	"github.com/teukumulya-ichsan/go-loan/models"
	"github.com/teukumulya-ichsan/go-loan/repositories"
	"github.com/teukumulya-ichsan/go-loan/services"
)

// LoanController interface
type LoanController interface {
	GetLoan(res http.ResponseWriter, req *http.Request)
	AddLoan(res http.ResponseWriter, req *http.Request)
	LoanTracked(res http.ResponseWriter, req *http.Request)
}

type controller struct {
	repositories.LoanRepository
}

var (
	service services.LoanService = services.NewLoanServices()
)

// NewLoanController Constructor
func NewLoanController(repo repositories.LoanRepository) LoanController {
	return &controller{
		repo,
	}
}

func (r *controller) GetLoan(res http.ResponseWriter, req *http.Request) {
	loans, err := r.FindAll()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error getting data"})
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(loans)
}

func (r *controller) AddLoan(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	var loan models.Loan

	//decode from req.body  into loan with reference ways "&loan"
	err := json.NewDecoder(req.Body).Decode(&loan)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error marshalling the request"})

		return
	}

	isValid := service.Validate(&loan)
	if isValid != nil {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: isValid.Error()})
	}

	result, err2 := r.Save(&loan)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error Create Loan"})

		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(result)
}

func (r *controller) LoanTracked(res http.ResponseWriter, req *http.Request) {

	data := chi.URLParam(req, "date")

	loans, err := r.GetLast7days(data)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error getting data"})
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(loans)

}
