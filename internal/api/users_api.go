package api

import (
	"context"
	"net/http"

	"github.com/777777miSSU7777777/gaming-website/pkg/services/userservice"
)

func MakeNewUserHandler(svc userservice.UserService) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		req, err := userservice.DecodeNewUser(context.Background(), r)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
		}
		respVal, err := svc.NewUser(req.Name, req.Balance)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
		}
		err = userservice.EncodeResponse(context.Background(), rw, respVal)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
		} else {
			rw.WriteHeader(http.StatusOK)
		}
	}
}

func MakeGetUserHandler(svc userservice.UserService) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		req, err := userservice.DecodeGetUser(context.Background(), r)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
		}
		respVal, err := svc.GetUser(req.ID)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
		}
		err = userservice.EncodeResponse(context.Background(), rw, respVal)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
		} else {
			rw.WriteHeader(http.StatusOK)
		}
	}
}

func MakeDeleteUserHandler(svc userservice.UserService) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		req, err := userservice.DecodeDeleteUser(context.Background(), r)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
		}
		err = svc.DeleteUser(req.ID)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
		}
		err = userservice.EncodeResponse(context.Background(), rw, userservice.DeleteUserRequest{})
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
		} else {
			rw.WriteHeader(http.StatusOK)
		}
	}
}

func MakeUserTakeHandler(svc userservice.UserService) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		req, err := userservice.DecodeUserTake(context.Background(), r)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
		}
		respVal, err := svc.UserTake(req.ID, req.Points)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
		}
		err = userservice.EncodeResponse(context.Background(), rw, respVal)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
		} else {
			rw.WriteHeader(http.StatusOK)
		}
	}
}

func MakeUserFundHandler(svc userservice.UserService) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		req, err := userservice.DecodeUserFund(context.Background(), r)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
		}
		respVal, err := svc.UserFund(req.ID, req.Points)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
		}
		err = userservice.EncodeResponse(context.Background(), rw, respVal)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
		} else {
			rw.WriteHeader(http.StatusOK)
		}
	}
}
