package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/teukumulya-ichsan/go-loan/config"
	ph "github.com/teukumulya-ichsan/go-loan/src/loan/controllers"
)

func main() {

	connection, err := config.ConnectDB()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	pHandler := ph.NewLoanController(connection)
	r.Get("/loan", pHandler.Fetch)
	http.ListenAndServe(":3000", r)
}
