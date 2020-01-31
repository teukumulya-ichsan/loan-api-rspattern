package response

import "github.com/teukumulya-ichsan/loan-api-rspattern/models"

type Response struct {
	data     models.Loan
	paginate Paginate
}

type Paginate struct {
	Count   int32
	Summary int64
	Average float64
}
