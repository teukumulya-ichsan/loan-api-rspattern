package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type chiRouter struct{}

var (
	chiDispatcher = chi.NewRouter()
)

//NewChiRouter ...
func NewChiRouter() Router {
	return &chiRouter{}
}

// GET Method Chi
func (*chiRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Get(uri, f)
}

// POST Method Chi
func (*chiRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Post(uri, f)
}

// SERVE Method Chi
func (*chiRouter) SERVE(port string) {
	fmt.Printf("Chi HTPP Server running on port : %v", port)
	_ = http.ListenAndServe(port, chiDispatcher)
}
