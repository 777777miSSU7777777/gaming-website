package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/777777miSSU7777777/gaming-website/model"
	"context"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/777777miSSU7777777/gaming-website/service"
)

type ErrorResponse struct {
	Err error `json:"error"`
}

const BodyParseError error = fmt.Errorf("BODY PARSE ERROR")
const IDParseError error = fmt.Errorf("ID PARSE ERROR")
const PointsParseError = fmt.Errorf("POINTS PARSE ERROR")


func MakeNewUserHandler(svc service.UserService, logger *log.Logger) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var req NewUserRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			logger.Error(BodyParseError)
			json.NewEncoder(rw).Encode(ErrorResponse{BodyParseError})
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		
		u := model.User{Name: req.Name, req.Balance}
		err = u.Validate()
		if err != nil {
			logger.Error(err)
			json.NewEncoder(rw).Encode(ErrorResponse{err})
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		logger.Debug(u)

		resp, err := svc.NewUser(u.Name, u.Balance)
		if err != nil {
			logger.Error(err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		err := json.NewEncoder(rw).Encode(resp)
		if err != nil {
			logger.Error(err)
			json.NewEncoder(rw).Encode(err)
			rw.WriteHeader(http.StatusBadRequest)
		} 
	}
}

func MakeGetUserHandler(svc userservice.UserService, logger *log.Logger) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var req GetUserRequest
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {

		}

		resp, err := svc.GetUser(req.ID)
		if err != nil {
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
