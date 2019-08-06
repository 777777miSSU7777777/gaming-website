package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/777777miSSU7777777/gaming-website/service"
)

func NewHttpServer(svc service.Service) http.Handler {
	r := mux.NewRouter()
	r.Use(jsonTypeMiddleware)

	r.Methods("POST").Path("/user").Handler(MakeNewUserHandler(svc))
	r.Methods("GET").Path("/user/{id}").Handler(MakeGetUserHandler(svc))
	r.Methods("DELETE").Path("/user/{id}").Handler(MakeDeleteUserHandler(svc))
	r.Methods("POST").Path("/user/{id}/take").Handler(MakeUserTakeHandler(svc))
	r.Methods("POST").Path("/user/{id}/fund").Handler(MakeUserFundHandler(svc))

	return r
}

func jsonTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}
