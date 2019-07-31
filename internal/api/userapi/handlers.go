package userapi

import (
	"context"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/777777miSSU7777777/gaming-website/pkg/services"
	"github.com/777777miSSU7777777/gaming-website/pkg/services/userservice"
)

func MakeNewUserHandler(svc userservice.UserService, logger *log.Logger) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		req, err := userservice.DecodeNewUser(context.Background(), r)
		if err != nil {
			logger.Warningln(err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		respVal, err := svc.NewUser(req.Name, req.Balance)
		if err != nil {
			logger.Warningln(err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		err = services.EncodeResponse(context.Background(), rw, respVal)
		if err != nil {
			logger.Warningln(err)
			rw.WriteHeader(http.StatusBadRequest)
		}
	}
}

func MakeGetUserHandler(svc userservice.UserService, logger *log.Logger) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		req, err := userservice.DecodeGetUser(context.Background(), r)
		if err != nil {
			logger.Warningln(err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		respVal, err := svc.GetUser(req.ID)
		if err != nil {
			logger.Warningln(err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		err = services.EncodeResponse(context.Background(), rw, respVal)
		if err != nil {
			logger.Warningln(err)
			rw.WriteHeader(http.StatusBadRequest)
		}
	}
}

func MakeDeleteUserHandler(svc userservice.UserService, logger *log.Logger) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		req, err := userservice.DecodeDeleteUser(context.Background(), r)
		if err != nil {
			logger.Warningln(err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		err = svc.DeleteUser(req.ID)
		if err != nil {
			logger.Warningln(err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		err = services.EncodeResponse(context.Background(), rw, userservice.DeleteUserRequest{})
		if err != nil {
			logger.Warningln(err)
			rw.WriteHeader(http.StatusBadRequest)
		}
	}
}

func MakeUserTakeHandler(svc userservice.UserService, logger *log.Logger) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		req, err := userservice.DecodeUserTake(context.Background(), r)
		if err != nil {
			logger.Warningln(err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		respVal, err := svc.UserTake(req.ID, req.Points)
		if err != nil {
			logger.Warningln(err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		err = services.EncodeResponse(context.Background(), rw, respVal)
		if err != nil {
			logger.Warningln(err)
			rw.WriteHeader(http.StatusBadRequest)
		}
	}
}

func MakeUserFundHandler(svc userservice.UserService, logger *log.Logger) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		req, err := userservice.DecodeUserFund(context.Background(), r)
		if err != nil {
			logger.Warningln(err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		respVal, err := svc.UserFund(req.ID, req.Points)
		if err != nil {
			logger.Warningln(err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		err = services.EncodeResponse(context.Background(), rw, respVal)
		if err != nil {
			logger.Warningln(err)
			rw.WriteHeader(http.StatusBadRequest)
		}
	}
}
