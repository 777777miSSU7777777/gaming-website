package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/777777miSSU7777777/gaming-website/model"
	"github.com/gorilla/mux"
)

type ErrorResponse struct {
	Err string `json:"error"`
}

var BodyParseError error = errors.New("BODY PARSE ERROR")
var IDParseError error = errors.New("ID PARSE ERROR")
var PointsParseError error = errors.New("POINTS PARSE ERROR")
var UserNotFoundError error = errors.New("USER NOT FOUND ERROR")

type Service interface {
	NewUser(string, int64) (model.User, error)
	GetUser(int64) (model.User, error)
	DeleteUser(int64) error
	UserTake(int64, int64) (model.User, error)
	UserFund(int64, int64) (model.User, error)
}

type API struct {
	svc Service
}

func New(svc Service) API {
	return API{svc}
}

func (a API) NewUser(rw http.ResponseWriter, r *http.Request) {
	var req NewUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(ErrorResponse{BodyParseError.Error()})
		return
	}

	u := model.User{Username: req.Name, Balance: req.Balance}
	err = u.Validate()
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(ErrorResponse{err.Error()})
		return
	}

	resp, err := a.svc.NewUser(u.Username, u.Balance)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(ErrorResponse{err.Error()})
		return
	}

	err = json.NewEncoder(rw).Encode(NewUserResponse{resp.ID, resp.Username, resp.Balance})
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(ErrorResponse{err.Error()})
	}
}

func (a API) GetUser(rw http.ResponseWriter, r *http.Request) {
	var req GetUserRequest
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(ErrorResponse{IDParseError.Error()})
		return
	}
	req.ID = id

	resp, err := a.svc.GetUser(req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			rw.WriteHeader(http.StatusNotFound)
			_ = json.NewEncoder(rw).Encode(ErrorResponse{UserNotFoundError.Error()})
		} else {
			rw.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(rw).Encode(ErrorResponse{err.Error()})
		}
		return
	}

	err = json.NewEncoder(rw).Encode(GetUserResponse{resp.ID, resp.Username, resp.Balance})
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(ErrorResponse{err.Error()})
	}
}

func (a API) DeleteUser(rw http.ResponseWriter, r *http.Request) {
	var req DeleteUserRequest
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(ErrorResponse{IDParseError.Error()})
		return
	}
	req.ID = id

	err = a.svc.DeleteUser(req.ID)
	if err != nil {
		if err.Error() == UserNotFoundError.Error() {
			rw.WriteHeader(http.StatusNotFound)
			_ = json.NewEncoder(rw).Encode(ErrorResponse{UserNotFoundError.Error()})
		} else {
			rw.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(rw).Encode(ErrorResponse{err.Error()})
		}
		return
	}

	err = json.NewEncoder(rw).Encode(DeleteUserRequest{})
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(ErrorResponse{err.Error()})
	}
}

func (a API) UserTake(rw http.ResponseWriter, r *http.Request) {
	var req UserTakeRequest
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(ErrorResponse{IDParseError.Error()})
		return
	}
	req.ID = id

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(ErrorResponse{BodyParseError.Error()})
		return
	}

	resp, err := a.svc.UserTake(req.ID, req.Points)
	if err != nil {
		if err == sql.ErrNoRows {
			rw.WriteHeader(http.StatusNotFound)
			_ = json.NewEncoder(rw).Encode(ErrorResponse{UserNotFoundError.Error()})
		} else {
			rw.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(rw).Encode(ErrorResponse{err.Error()})
		}
		return
	}

	err = json.NewEncoder(rw).Encode(UserTakeResponse{resp.ID, resp.Username, resp.Balance})
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(ErrorResponse{err.Error()})
	}
}

func (a API) UserFund(rw http.ResponseWriter, r *http.Request) {
	var req UserFundRequest
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(ErrorResponse{IDParseError.Error()})
		return
	}
	req.ID = id

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(ErrorResponse{BodyParseError.Error()})
		return
	}

	resp, err := a.svc.UserFund(req.ID, req.Points)
	if err != nil {
		if err == sql.ErrNoRows {
			rw.WriteHeader(http.StatusNotFound)
			_ = json.NewEncoder(rw).Encode(ErrorResponse{UserNotFoundError.Error()})
		} else {
			rw.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(rw).Encode(ErrorResponse{err.Error()})
		}
		return
	}

	err = json.NewEncoder(rw).Encode(UserFundResponse{resp.ID, resp.Username, resp.Balance})
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(ErrorResponse{err.Error()})
	}
}
