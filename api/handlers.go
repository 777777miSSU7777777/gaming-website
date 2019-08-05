package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/777777miSSU7777777/gaming-website/model"
	"github.com/gorilla/mux"

	log "github.com/sirupsen/logrus"

	"github.com/777777miSSU7777777/gaming-website/service"
)

type ErrorResponse struct {
	Err string `json:"error"`
}

var BodyParseError error = fmt.Errorf("BODY PARSE ERROR")
var IDParseError error = fmt.Errorf("ID PARSE ERROR")
var PointsParseError = fmt.Errorf("POINTS PARSE ERROR")

func MakeNewUserHandler(svc service.UserService, logger *log.Logger) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var req NewUserRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			logger.Error(BodyParseError)
			rw.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(rw).Encode(ErrorResponse{BodyParseError.Error()})
			return
		}

		u := model.User{Username: req.Name, Balance: req.Balance}
		err = u.Validate()
		if err != nil {
			logger.Error(err)
			rw.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(rw).Encode(ErrorResponse{err.Error()})
			return
		}
		logger.Debug(u)

		resp, err := svc.NewUser(u.Username, u.Balance)
		if err != nil {
			logger.Error(err)
			rw.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(rw).Encode(ErrorResponse{err.Error()})
			return
		}

		err = json.NewEncoder(rw).Encode(NewUserResponse{resp.ID, resp.Username, resp.Balance})
		if err != nil {
			logger.Error(err)
			rw.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(rw).Encode(ErrorResponse{err.Error()})
		}
	}
}

func MakeGetUserHandler(svc service.UserService, logger *log.Logger) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var req GetUserRequest
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			logger.Error(err)
			rw.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(rw).Encode(ErrorResponse{IDParseError.Error()})
			return
		}
		req.ID = id

		resp, err := svc.GetUser(req.ID)
		if err != nil {
			logger.Error(err)
			rw.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(rw).Encode(ErrorResponse{err.Error()})
			return
		}

		err = json.NewEncoder(rw).Encode(GetUserResponse{resp.ID, resp.Username, resp.Balance})
		if err != nil {
			logger.Error(err)
			rw.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(rw).Encode(ErrorResponse{err.Error()})
		}
	}
}

func MakeDeleteUserHandler(svc service.UserService, logger *log.Logger) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var req DeleteUserRequest
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			logger.Error(err)
			rw.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(rw).Encode(ErrorResponse{IDParseError.Error()})
			return
		}
		req.ID = id

		err = svc.DeleteUser(req.ID)
		if err != nil {
			logger.Error(err)
			rw.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(rw).Encode(ErrorResponse{err.Error()})
			return
		}

		err = json.NewEncoder(rw).Encode(DeleteUserRequest{})
		if err != nil {
			logger.Error(ErrorResponse{err.Error()})
			rw.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(rw).Encode(ErrorResponse{err.Error()})
		}
	}
}

func MakeUserTakeHandler(svc service.UserService, logger *log.Logger) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var req UserTakeRequest
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			logger.Error(err)
			rw.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(rw).Encode(ErrorResponse{IDParseError.Error()})
			return
		}
		req.ID = id

		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			logger.Error(BodyParseError)
			rw.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(rw).Encode(ErrorResponse{BodyParseError.Error()})
			return
		}

		resp, err := svc.UserTake(req.ID, req.Points)
		if err != nil {
			logger.Error(err)
			rw.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(rw).Encode(ErrorResponse{err.Error()})
			return
		}

		err = json.NewEncoder(rw).Encode(UserTakeResponse{resp.ID, resp.Username, resp.Balance})
		if err != nil {
			logger.Error(err)
			rw.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(rw).Encode(ErrorResponse{err.Error()})
		}
	}
}

func MakeUserFundHandler(svc service.UserService, logger *log.Logger) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var req UserFundRequest
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			logger.Error(err)
			rw.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(rw).Encode(ErrorResponse{IDParseError.Error()})
			return
		}
		req.ID = id

		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			logger.Error(BodyParseError)
			rw.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(rw).Encode(ErrorResponse{BodyParseError.Error()})
			return
		}

		resp, err := svc.UserFund(req.ID, req.Points)
		if err != nil {
			logger.Error(err)
			rw.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(rw).Encode(ErrorResponse{err.Error()})
			return
		}

		err = json.NewEncoder(rw).Encode(UserFundResponse{resp.ID, resp.Username, resp.Balance})
		if err != nil {
			logger.Warningln(err)
			rw.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(rw).Encode(ErrorResponse{err.Error()})
		}
	}
}
