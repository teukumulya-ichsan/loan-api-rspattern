package response

import "github.com/teukumulya-ichsan/loan-api-rspattern/models"

//Respond ...
type Respond struct {
	Message string        `json:"message"`
	Data    []models.Loan `json:"data"`
}

//summary struct ...
type summary struct {
	Count   int32
	Sum     int64
	Average float64
}
