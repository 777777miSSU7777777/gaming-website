package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/777777miSSU7777777/gaming-website/model"
	"github.com/gorilla/mux"
)

type ErrorResponse struct {
	Type  string `json:"type"`
	Error string `json:"error"`
}

var BodyParseError = "BODY PARSE ERROR"
var IDParseError = "ID PARSE ERROR"
var PointsParseError = "POINTS PARSE ERROR"
var ValidationError = "VALIDATION ERROR"
var NotFoundError = "NOT FOUND ERROR"
var ServiceError = "SERVICE ERROR"

func writeError(w http.ResponseWriter, statusCode int, errType string, err error) {
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(ErrorResponse{Type: errType, Error: err.Error()})
}

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

func (a API) NewUser(w http.ResponseWriter, r *http.Request) {
	var req NewUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, 400, BodyParseError, fmt.Errorf("error while parsing body: %v", err))
		return
	}

	u := model.User{Username: req.Name, Balance: req.Balance}
	err = u.Validate()
	if err != nil {
		writeError(w, 400, ValidationError, fmt.Errorf("user validation error: %v", err))
		return
	}

	resp, err := a.svc.NewUser(u.Username, u.Balance)
	if err != nil {
		writeError(w, 400, ServiceError, fmt.Errorf("error while adding user: %v", err))
		return
	}

	_ = json.NewEncoder(w).Encode(NewUserResponse{resp.ID, resp.Username, resp.Balance})
}

func (a API) GetUser(w http.ResponseWriter, r *http.Request) {
	var req GetUserRequest
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		writeError(w, 400, IDParseError, fmt.Errorf("error while parsing id: %v", err))
		return
	}
	req.ID = id

	resp, err := a.svc.GetUser(req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			writeError(w, 404, NotFoundError, fmt.Errorf("user not found error: %v", err))
		} else {
			writeError(w, 400, ServiceError, fmt.Errorf("error while getting user: %v", err))
		}
		return
	}

	err = json.NewEncoder(w).Encode(GetUserResponse{resp.ID, resp.Username, resp.Balance})
}

func (a API) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var req DeleteUserRequest
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		writeError(w, 400, IDParseError, fmt.Errorf("error while parsing id: %v", err))
		return
	}
	req.ID = id

	err = a.svc.DeleteUser(req.ID)
	if err != nil {
		if err.Error() == "user not found error" {
			writeError(w, 404, NotFoundError, fmt.Errorf("user not found error"))
			return
		} else {
			writeError(w, 400, ServiceError, fmt.Errorf("error while deleting user: %v", err))
		}
		return
	}

	_ = json.NewEncoder(w).Encode(DeleteUserRequest{})
}

func (a API) UserTake(w http.ResponseWriter, r *http.Request) {
	var req UserTakeRequest
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		writeError(w, 400, IDParseError, fmt.Errorf("error while parsing id: %v", err))
		return
	}
	req.ID = id

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, 400, BodyParseError, fmt.Errorf("error while parsing body: %v", err))
		return
	}

	resp, err := a.svc.UserTake(req.ID, req.Points)
	if err != nil {
		if err == sql.ErrNoRows {
			writeError(w, 404, NotFoundError, fmt.Errorf("user not found error: %v", err))
		} else {
			writeError(w, 400, ServiceError, fmt.Errorf("error while taking user balance: %v", err))
		}
		return
	}

	_ = json.NewEncoder(w).Encode(UserTakeResponse{resp.ID, resp.Username, resp.Balance})
}

func (a API) UserFund(w http.ResponseWriter, r *http.Request) {
	var req UserFundRequest
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		writeError(w, 400, IDParseError, fmt.Errorf("error while parsing id: %v", err))
		return
	}
	req.ID = id

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, 400, BodyParseError, fmt.Errorf("error while parsing body: %v", err))
		return
	}

	resp, err := a.svc.UserFund(req.ID, req.Points)
	if err != nil {
		if err == sql.ErrNoRows {
			writeError(w, 404, NotFoundError, fmt.Errorf("user not found error: %v", err))
		} else {
			writeError(w, 400, ServiceError, fmt.Errorf("error while adding user balance: %v", err))
		}
		return
	}

	_ = json.NewEncoder(w).Encode(UserFundResponse{resp.ID, resp.Username, resp.Balance})
}
