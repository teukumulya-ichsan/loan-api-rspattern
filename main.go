package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	ph "github.com/teukumulya-ichsan/go-loan/src/loan/controllers"
)

func main() {

	// connection, err := config.ConnectDB()
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(-1)
	// }
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	pHandler := ph.NewLoanController()
	r.Get("/loan", pHandler.FindAll)
	r.Post("/loan", pHandler.CreateLoan)
	r.Get("/loan/{id}", pHandler.FindByID)

	http.ListenAndServe(":3000", r)
}
