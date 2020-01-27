package main

import (
	"fmt"
	"net/http"

	"github.com/teukumulya-ichsan/go-loan/controllers"
	router "github.com/teukumulya-ichsan/go-loan/http"
	"github.com/teukumulya-ichsan/go-loan/repositories"
	"github.com/teukumulya-ichsan/go-loan/services"
)

var (
	loanRepository repositories.LoanRepository = repositories.NewPostgresRepository()
	loanService    services.LoanService        = services.NewLoanServices(loanRepository)
	loanController controllers.LoanController  = controllers.NewLoanController(loanService)
	httpRouter     router.Router               = router.NewChiRouter()
)

func main() {

	const port string = ":3000"
	httpRouter.GET("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "Up and Running...")
	})

	httpRouter.GET("/loans", loanController.GetLoan)
	httpRouter.POST("/loans", loanController.AddLoan)

	httpRouter.SERVE(port)
}
