package api

import (
	"net/http"

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

	return r
}

func jsonTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}
