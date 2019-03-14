package main

import (
	"flag"
	"net/http"

	"github.com/gorilla/mux"
	v1 "github.com/kwtucker/bitlyapi/api/v1"
)

func main() {
	port := flag.String("port", "8080", "The port the server will be running on.")
	flag.Parse()

	var router = mux.NewRouter()

	var api = router.PathPrefix("/api").Subrouter()
	api.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	// Routes for the version one of the api.
	v1.Router(api.PathPrefix("/v1").Subrouter())

	api.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("Authorization") == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	})

	http.ListenAndServe("localhost:"+*port, router)
}
