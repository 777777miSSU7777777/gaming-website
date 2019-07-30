package userservice

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DecodeNewUser(_ context.Context, r *http.Request) (NewUserRequest, error) {
	req := NewUserRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(req)
	if err != nil {
		return NewUserRequest{}, err
	}
	return req, nil
}

func DecodeGetUser(_ context.Context, r *http.Request) (GetUserRequest, error) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return GetUserRequest{}, err
	}
	return GetUserRequest{id}, nil
}

func DecodeDeleteUser(_ context.Context, r *http.Request) (DeleteUserRequest, error) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return DeleteUserRequest{}, err
	}
	return DeleteUserRequest{id}, nil
}

func DecodeUserTake(_ context.Context, r *http.Request) (UserTakeRequest, error) {
	req := UserTakeRequest{}
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return UserTakeRequest{}, err
	}
	req.ID = id
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(req)
	if err != nil {
		return UserTakeRequest{}, err
	}
	return req, nil
}

func DecodeUserFund(_ context.Context, r *http.Request) (UserFundRequest, error) {
	req := UserFundRequest{}
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return UserFundRequest{}, err
	}
	req.ID = id
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(req)
	if err != nil {
		return UserFundRequest{}, err
	}
	return req, nil
}

func EncodeResponse(_ context.Context, rw http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(rw).Encode(response)
}
