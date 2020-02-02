package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/teukumulya-ichsan/loan-api-rspattern/models"
	"github.com/teukumulya-ichsan/loan-api-rspattern/repositories"
	resp "github.com/teukumulya-ichsan/loan-api-rspattern/response"
	"github.com/teukumulya-ichsan/loan-api-rspattern/services"
)

// loanController interface
type loanController interface {
	GetLoan(res http.ResponseWriter, req *http.Request)
	AddLoan(res http.ResponseWriter, req *http.Request)
	LoanTracked(res http.ResponseWriter, req *http.Request)
}

type controller struct {
	repositories.LoanRepository
}

var (
	service = services.NewLoanServices()
)

// NewLoanController Constructor
func NewLoanController(repo repositories.LoanRepository) loanController {
	return &controller{
		repo,
	}
}

// GetLoan ...
func (r *controller) GetLoan(res http.ResponseWriter, _ *http.Request) {
	loans, err := r.FindAll()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		//json.NewEncoder(res).Encode(resp.Error{message: "Error getting data"})
		json.NewEncoder(res).Encode(resp.Error{Message: "Error Getting Data"})
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(resp.Respond{Data: loans})
}

// AddLoan ...
func (r *controller) AddLoan(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	var loan models.Loan

	//decode from req.body  into loan with reference ways "&loan"
	err := json.NewDecoder(req.Body).Decode(&loan)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(resp.Error{Message: "Error marshalling the request"})

		return
	}

	isValid := service.Validate(&loan)
	if isValid != nil {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(resp.Error{Message: isValid.Error()})
	}

	_, err2 := r.Save(&loan)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(resp.Error{Message: "Error Create Loan"})

		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(resp.Respond{Message: "Loans is Created"})
}

// LoanTracked ...
func (r *controller) LoanTracked(res http.ResponseWriter, req *http.Request) {

	// getting date from URL params
	date := chi.URLParam(req, "date")

	// get last 7 days data by date
	data, err := r.GetLast7days(date)

	// processing data to get summary of data
	summary,_ := service.GetSummary(data)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(resp.Error{Message: "Error getting data"})
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	//json.NewEncoder(res).Encode(data)
	json.NewEncoder(res).Encode(resp.Respond{Data: data, Summary: summary})
}
