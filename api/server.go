package api

import (
	"net/http"
	"os"
	"fmt"

	"github.com/gorilla/mux"
)

func NewHttpServer(api API) http.Handler {
	r := mux.NewRouter()
	r.Use(jsonTypeMiddleware)

	r.Methods("POST").Path("/user").HandlerFunc(api.NewUser)
	r.Methods("GET").Path("/user/{id}").HandlerFunc(api.GetUser)
	r.Methods("DELETE").Path("/user/{id}").HandlerFunc(api.DeleteUser)
	r.Methods("POST").Path("/user/{id}/take").HandlerFunc(api.UserTake)
	r.Methods("POST").Path("/user/{id}/fund").HandlerFunc(api.UserFund)

	r.Methods("POST").Path("/tournament").HandlerFunc(api.NewTournament)
	r.Methods("GET").Path("/tournament/{id}").HandlerFunc(api.GetTournament)
	r.Methods("POST").Path("/tournament/{id}/join").HandlerFunc(api.JoinTournament)
	r.Methods("POST").Path("/tournament/{id}/finish").HandlerFunc(api.FinishTournament)
	r.Methods("DELETE").Path("/tournament/{id}").HandlerFunc(api.CancelTournament)

	r.Methods("GET").Path("/health-check").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	r.Methods("POST").Path("/panic").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		os.Exit(1)
	})
	r.Methods("GET").Path("/hello").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})

	return r
}

func jsonTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}
