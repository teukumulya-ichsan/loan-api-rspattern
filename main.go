package main

import (
	"fmt"
	"net/http"

	"github.com/teukumulya-ichsan/go-loan/controllers"
	router "github.com/teukumulya-ichsan/go-loan/http"
	"github.com/teukumulya-ichsan/go-loan/repositories"
)

var (
	loanRepository repositories.LoanRepository = repositories.NewPostgresRepository()
	loanController controllers.LoanController  = controllers.NewLoanController(loanRepository)
	httpRouter     router.Router               = router.NewChiRouter()
)

func main() {

	const port string = ":3000"
	httpRouter.GET("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "Up and Running...")
	})

	httpRouter.GET("/loans", loanController.GetLoan)
	httpRouter.POST("/loans", loanController.AddLoan)
	httpRouter.GET("/loans/tracked/{date}", loanController.LoanTracked)

	httpRouter.SERVE(port)
}
