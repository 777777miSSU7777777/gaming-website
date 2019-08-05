package api

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/777777miSSU7777777/gaming-website/service"
)

func NewHttpServer(svc service.UserService, logger *log.Logger) http.Handler {
	r := mux.NewRouter()
	r.Use(jsonTypeMiddleware)

	r.Methods("POST").Path("/user").Handler(MakeNewUserHandler(svc, logger))
	r.Methods("GET").Path("/user/{id}").Handler(MakeGetUserHandler(svc, logger))
	r.Methods("DELETE").Path("/user/{id}").Handler(MakeDeleteUserHandler(svc, logger))
	r.Methods("POST").Path("/user/{id}/take").Handler(MakeUserTakeHandler(svc, logger))
	r.Methods("POST").Path("/user/{id}/fund").Handler(MakeUserFundHandler(svc, logger))

	return r
}

func jsonTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}
