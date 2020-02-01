package main

import (
	"fmt"
	"net/http"

	"github.com/teukumulya-ichsan/loan-api-rspattern/controllers"
	router "github.com/teukumulya-ichsan/loan-api-rspattern/http"
	"github.com/teukumulya-ichsan/loan-api-rspattern/repositories"
)

var (
	loanRepository = repositories.NewPostgresRepository()
	loanController = controllers.NewLoanController(loanRepository)
	httpRouter     = router.NewChiRouter()
)

func main() {

	const port string = ":3001"
	httpRouter.GET("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "Up and Running...")
	})

	httpRouter.GET("/loans", loanController.GetLoan)
	httpRouter.POST("/loans", loanController.AddLoan)
	httpRouter.GET("/loans/tracked/{date}", loanController.LoanTracked)

	httpRouter.SERVE(port)
}
