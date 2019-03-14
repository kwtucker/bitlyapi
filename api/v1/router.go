package v1

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Router(router *mux.Router) {
	router.HandleFunc("/groups/default/countries", HandleGroupsCountries).Methods("GET")
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	})
}
