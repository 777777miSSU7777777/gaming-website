package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/777777miSSU7777777/gaming-website/service"

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

type API struct {
	svc service.Service
}

func New(svc service.Service) API {
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

	_ = json.NewEncoder(w).Encode(DeleteUserResponse{})
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

func (a API) NewTournament(w http.ResponseWriter, r *http.Request) {
	var req NewTournamentRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, 400, BodyParseError, fmt.Errorf("error while parsing body: %v", err))
		return
	}

	t := model.Tournament{TournamentName: req.Name, Deposit: req.Deposit}
	err = t.Validate()
	if err != nil {
		writeError(w, 400, ValidationError, fmt.Errorf("tournament validation error: %v", err))
		return
	}

	resp, err := a.svc.NewTournament(t.TournamentName, t.Deposit)
	if err != nil {
		writeError(w, 400, ServiceError, fmt.Errorf("error while adding tournament: %v", err))
		return
	}

	_ = json.NewEncoder(w).Encode(NewTournamentResponse{ID: resp.ID, Name: resp.TournamentName, Deposit: resp.Deposit, Prize: resp.Prize})
}

func wrapNotFinishedUsers(users []model.User) []NotFinishedUser {
	wrapped := make([]NotFinishedUser, 0, len(users))
	for _, u := range users {
		wrap := NotFinishedUser{ID: u.ID, Name: u.Username}
		wrapped = append(wrapped, wrap)
	}

	return wrapped
}

func wrapFinishedUsers(users []model.User, winnerID int64) []FinishedUser {
	wrapped := make([]FinishedUser, 0, len(users))
	for _, u := range users {
		wrap := FinishedUser{ID: u.ID, Name: u.Username}
		if wrap.ID == winnerID {
			wrap.Winner = true
		}
		wrapped = append(wrapped, wrap)
	}

	return wrapped
}

func (a API) GetTournament(w http.ResponseWriter, r *http.Request) {
	var req GetTournamentRequest
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		writeError(w, 400, IDParseError, fmt.Errorf("error while parsing id: %v", err))
		return
	}
	req.ID = id

	t, u, err := a.svc.GetTournament(req.ID)
	if err != nil {
		if err.Error() == "tournament not found error" {
			writeError(w, 404, NotFoundError, fmt.Errorf("tournament not found error"))
		} else {
			writeError(w, 404, ServiceError, fmt.Errorf("error while getting tournament: %v", err))
		}
		return
	}

	wrappedUsers := wrapNotFinishedUsers(u)

	_ = json.NewEncoder(w).Encode(GetTournamentResponse{ID: t.ID, Name: t.TournamentName, Deposit: t.Deposit, Prize: t.Prize, Users: wrappedUsers})
}

func (a API) JoinTournament(w http.ResponseWriter, r *http.Request) {
	var req JoinTournamentRequest
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		writeError(w, 400, IDParseError, fmt.Errorf("error while parsing id: %v", err))
		return
	}
	req.TournamentID = id

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, 400, BodyParseError, fmt.Errorf("error while parsing body: %v", err))
		return
	}

	t, u, err := a.svc.JoinTournament(req.TournamentID, req.UserID)
	if err != nil {
		if err.Error() == "user not found error" || err.Error() == "tournament not found error" {
			writeError(w, 404, NotFoundError, err)
		} else {
			writeError(w, 400, ServiceError, fmt.Errorf("error while join user to tournament: %v", err))
		}
		return
	}

	wrappedUsers := wrapNotFinishedUsers(u)

	_ = json.NewEncoder(w).Encode(JoinTournamentResponse{ID: t.ID, Name: t.TournamentName, Deposit: t.Deposit, Prize: t.Prize, Users: wrappedUsers})
}

func (a API) FinishTournament(w http.ResponseWriter, r *http.Request) {
	var req FinishTournamentRequest
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		writeError(w, 400, IDParseError, fmt.Errorf("error while parsing id: %v", err))
		return
	}
	req.ID = id

	t, u, err := a.svc.FinishTournament(req.ID)
	if err != nil {
		if err.Error() == "user not found error" || err.Error() == "tournament not found error" {
			writeError(w, 404, NotFoundError, err)
		} else {
			writeError(w, 400, ServiceError, fmt.Errorf("error while finishing tournament: %v", err))
		}
		return
	}

	wrappedUsers := wrapFinishedUsers(u, t.WinnerID)

	_ = json.NewEncoder(w).Encode(FinishTournamentResponse{ID: t.ID, Name: t.TournamentName, Deposit: t.Deposit, Prize: t.Prize, Users: wrappedUsers})
}

func (a API) CancelTournament(w http.ResponseWriter, r *http.Request) {
	var req CancelTournamentRequest
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		writeError(w, 400, IDParseError, fmt.Errorf("error while parsing id: %v", err))
		return
	}
	req.ID = id

	err = a.svc.CancelTournament(req.ID)
	if err != nil {
		if err.Error() == "tournament not found error" {
			writeError(w, 404, NotFoundError, err)
		} else {
			writeError(w, 400, ServiceError, fmt.Errorf("error while canceling tournament: %v", err))
		}
		return
	}

	_ = json.NewEncoder(w).Encode(CancelTournamentResponse{})
}
