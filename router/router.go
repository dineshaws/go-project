package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter function
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/albums", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "not yet")
	}).Methods("GET")

	return r
}
