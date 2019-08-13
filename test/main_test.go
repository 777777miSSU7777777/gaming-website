package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/777777miSSU7777777/gaming-website/api"
)

// Runs user api tests with correct data
func TestFlow1(t *testing.T) {
	client := &http.Client{}
	baseURL := "http://127.0.0.1:8080"

	// New user test
	newUserReq := api.NewUserRequest{Name: "test_user", Balance: 1000}
	body, _ := json.Marshal(newUserReq)
	req, _ := http.NewRequest("POST", baseURL+"/user", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := client.Do(req)
	var newUserResp api.NewUserResponse
	_ = json.NewDecoder(resp.Body).Decode(&newUserResp)

	if !(newUserResp.ID != 0 && newUserResp.Name == newUserReq.Name && newUserResp.Balance == newUserReq.Balance) {
		t.Errorf("[NewUser] request - %+v and response - %+v", newUserReq, newUserResp)
	}

	_ = resp.Body.Close()

	// Get user test
	req, _ = http.NewRequest("GET", baseURL+fmt.Sprintf("/user/%v", newUserResp.ID), nil)
	req.Header.Set("Content-Type", "application/json")
	resp, _ = client.Do(req)
	var getUserResp api.GetUserResponse
	_ = json.NewDecoder(resp.Body).Decode(&getUserResp)

	if !(getUserResp.ID == newUserResp.ID && getUserResp.Name == newUserResp.Name && getUserResp.Balance == newUserResp.Balance) {
		t.Errorf("[GetUser] request id - %v and response - %+v", newUserResp.ID, getUserResp)
	}

	_ = resp.Body.Close()

	// Take user balance test
	userTakeReq := api.UserTakeRequest{Points: 500}
	body, _ = json.Marshal(userTakeReq)
	req, _ = http.NewRequest("POST", baseURL+fmt.Sprintf("/user/%v/take", newUserResp.ID), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ = client.Do(req)
	var userTakeResp api.UserTakeResponse
	_ = json.NewDecoder(resp.Body).Decode(&userTakeResp)

	if !(userTakeResp.ID == newUserResp.ID && userTakeResp.Name == newUserResp.Name && userTakeResp.Balance == newUserResp.Balance-userTakeReq.Points) {
		t.Errorf("[UserTake] request - id: %v, body: %+v and response - %+v", newUserResp.ID, userTakeReq, userTakeResp)
	}

	_ = resp.Body.Close()

	// Fund user balance test
	userFundReq := api.UserFundRequest{Points: 1000}
	body, _ = json.Marshal(userFundReq)
	req, _ = http.NewRequest("POST", baseURL+fmt.Sprintf("/user/%v/fund", newUserResp.ID), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ = client.Do(req)
	var userFundResp api.UserFundResponse
	_ = json.NewDecoder(resp.Body).Decode(&userFundResp)

	if !(userFundResp.ID == newUserResp.ID && userFundResp.Name == newUserResp.Name && userFundResp.Balance == userTakeResp.Balance+userFundReq.Points) {
		t.Errorf("[UserFund] request - id: %v, body: %+v and response - %+v", newUserResp.ID, userFundReq, userFundResp)
	}

	_ = resp.Body.Close()

	// Delete user balance test
	req, _ = http.NewRequest("DELETE", baseURL+fmt.Sprintf("/user/%v", newUserResp.ID), nil)
	req.Header.Set("Content-Type", "application/json")
	resp, _ = client.Do(req)

	if !(resp.StatusCode == http.StatusOK) {
		t.Errorf("[DeleteUser] status code - %v", resp.StatusCode)
	}

	_ = resp.Body.Close()
}
