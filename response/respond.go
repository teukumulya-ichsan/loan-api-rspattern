package response

import "github.com/teukumulya-ichsan/loan-api-rspattern/models"

//Respond ...
type Respond struct {
	Data    []models.Loan `json:"data"`
	Summary []Info        `json:"info"`
}

//Info struct ...
type Info struct {
	Count   int64   `json:"count"`
	Sum     int64   `json:"summary"`
	Average float64 `json:"7-day-avg"`
}
