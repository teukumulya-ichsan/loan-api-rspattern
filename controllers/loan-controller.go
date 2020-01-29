package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/teukumulya-ichsan/go-loan/errors"
	"github.com/teukumulya-ichsan/go-loan/models"
	"github.com/teukumulya-ichsan/go-loan/services"
)

type controller struct {
	services.LoanService
}

// LoanController interface
type LoanController interface {
	GetLoan(res http.ResponseWriter, req *http.Request)
	AddLoan(res http.ResponseWriter, req *http.Request)
}

// NewLoanController Contructor
func NewLoanController(service services.LoanService) LoanController {
	return &controller{
		service,
	}
}

func (service *controller) GetLoan(res http.ResponseWriter, req *http.Request) {
	loans, err := service.FindAll()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error getting data"})
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(loans)
}

func (service *controller) AddLoan(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")

	var loan models.Loan

	//decode dari req.body ke dalam variable loan dengan cara reference "&loan"
	err := json.NewDecoder(req.Body).Decode(&loan)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error marshalling the request"})

		return
	}

	loan.ID = rand.Int63()

	// adding new data to the Loans Array
	err1 := service.Validate(&loan)
	if err1 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: err1.Error()})

		return
	}

	result, err2 := service.Create(&loan)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message: "Error Create Loan"})

		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(result)
}
