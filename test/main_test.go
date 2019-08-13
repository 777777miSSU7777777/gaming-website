package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/777777miSSU7777777/gaming-website/model"

	"github.com/777777miSSU7777777/gaming-website/api"
	"github.com/stretchr/testify/require"
)

var baseURL = "http://127.0.0.1:8080"

// Runs user api tests with correct data
func TestFlow1(t *testing.T) {
	client := &http.Client{}

	// New user test
	newUserReq := api.NewUserRequest{Name: "test_user", Balance: 1000}
	body, _ := json.Marshal(newUserReq)
	req, _ := http.NewRequest("POST", baseURL+"/user", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := client.Do(req)
	var newUserResp api.NewUserResponse
	_ = json.NewDecoder(resp.Body).Decode(&newUserResp)

	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.NotEqual(t, int64(0), newUserResp.ID)
	require.Equal(t, newUserReq.Name, newUserResp.Name)
	require.Equal(t, newUserReq.Balance, newUserResp.Balance)

	_ = resp.Body.Close()

	// Get user test
	req, _ = http.NewRequest("GET", baseURL+fmt.Sprintf("/user/%v", newUserResp.ID), nil)
	req.Header.Set("Content-Type", "application/json")
	resp, _ = client.Do(req)
	var getUserResp api.GetUserResponse
	_ = json.NewDecoder(resp.Body).Decode(&getUserResp)

	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Equal(t, newUserResp.ID, getUserResp.ID)
	require.Equal(t, newUserResp.Name, getUserResp.Name)
	require.Equal(t, newUserResp.Balance, getUserResp.Balance)

	_ = resp.Body.Close()

	// Take user balance test
	userTakeReq := api.UserTakeRequest{Points: 500}
	body, _ = json.Marshal(userTakeReq)
	req, _ = http.NewRequest("POST", baseURL+fmt.Sprintf("/user/%v/take", newUserResp.ID), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ = client.Do(req)
	var userTakeResp api.UserTakeResponse
	_ = json.NewDecoder(resp.Body).Decode(&userTakeResp)

	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Equal(t, newUserResp.ID, userTakeResp.ID)
	require.Equal(t, newUserResp.Name, userTakeResp.Name)
	require.Equal(t, newUserResp.Balance-userTakeReq.Points, userTakeResp.Balance)

	_ = resp.Body.Close()

	// Fund user balance test
	userFundReq := api.UserFundRequest{Points: 1000}
	body, _ = json.Marshal(userFundReq)
	req, _ = http.NewRequest("POST", baseURL+fmt.Sprintf("/user/%v/fund", newUserResp.ID), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ = client.Do(req)
	var userFundResp api.UserFundResponse
	_ = json.NewDecoder(resp.Body).Decode(&userFundResp)

	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Equal(t, newUserResp.ID, userFundResp.ID)
	require.Equal(t, newUserResp.Name, userFundResp.Name)
	require.Equal(t, userTakeResp.Balance+userFundReq.Points, userFundResp.Balance)

	_ = resp.Body.Close()

	// Delete user balance test
	req, _ = http.NewRequest("DELETE", baseURL+fmt.Sprintf("/user/%v", newUserResp.ID), nil)
	req.Header.Set("Content-Type", "application/json")
	resp, _ = client.Do(req)

	require.Equal(t, http.StatusOK, resp.StatusCode)

	_ = resp.Body.Close()
}

// Runs user api tests with incorrect data
func TestFlow2(t *testing.T) {
	client := &http.Client{}

	// New user with empty name
	newUserReq := api.NewUserRequest{Name: "", Balance: 1000}
	body, _ := json.Marshal(newUserReq)
	req, _ := http.NewRequest("POST", baseURL+"/user", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := client.Do(req)
	var errResp api.ErrorResponse
	_ = json.NewDecoder(resp.Body).Decode(&errResp)

	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
	require.Equal(t, api.ValidationError, errResp.Type)

	_ = resp.Body.Close()

	// New user with negative balance
	newUserReq = api.NewUserRequest{Name: "tedt_user", Balance: -1000}
	body, _ = json.Marshal(newUserReq)
	req, _ = http.NewRequest("POST", baseURL+"/user", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ = client.Do(req)
	_ = json.NewDecoder(resp.Body).Decode(&errResp)

	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
	require.Equal(t, api.ValidationError, errResp.Type)

	_ = resp.Body.Close()

	// New user with incorrect body
	req, _ = http.NewRequest("POST", baseURL+"/user", bytes.NewBuffer([]byte(`{"name":123, "balance:"abc"}`)))
	req.Header.Set("Content-Type", "application/json")
	resp, _ = client.Do(req)
	_ = json.NewDecoder(resp.Body).Decode(&errResp)

	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
	require.Equal(t, api.BodyParseError, errResp.Type)

	_ = resp.Body.Close()

	// Get not existing user
	req, _ = http.NewRequest("GET", baseURL+"/user/-1", nil)
	req.Header.Set("Content-Type", "application/json")
	resp, _ = client.Do(req)
	_ = json.NewDecoder(resp.Body).Decode(&errResp)

	require.Equal(t, http.StatusNotFound, resp.StatusCode)
	require.Equal(t, api.NotFoundError, errResp.Type)

	_ = resp.Body.Close()

	// Get with invalid id
	req, _ = http.NewRequest("GET", baseURL+"/user/abc", nil)
	req.Header.Set("Content-Type", "application/json")
	resp, _ = client.Do(req)
	_ = json.NewDecoder(resp.Body).Decode(&errResp)

	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
	require.Equal(t, api.IDParseError, errResp.Type)

	_ = resp.Body.Close()

	// Delete not existing user
	req, _ = http.NewRequest("DELETE", baseURL+"/user/-1", nil)
	req.Header.Set("Content-Type", "application/json")
	resp, _ = client.Do(req)
	_ = json.NewDecoder(resp.Body).Decode(&errResp)

	require.Equal(t, http.StatusNotFound, resp.StatusCode)
	require.Equal(t, api.NotFoundError, errResp.Type)

	_ = resp.Body.Close()

	// Create correct user to test take and fund actions with incorrect data
	newUserReq = api.NewUserRequest{Name: "test_user", Balance: 1000}
	body, _ = json.Marshal(newUserReq)
	req, _ = http.NewRequest("POST", baseURL+"/user", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ = client.Do(req)
	var testUser api.NewUserResponse
	_ = json.NewDecoder(resp.Body).Decode(&testUser)

	// User take with 0 points
	userTakeReq := api.UserTakeRequest{Points: 0}
	body, _ = json.Marshal(userTakeReq)
	req, _ = http.NewRequest("POST", baseURL+fmt.Sprintf("/user/%v/take", testUser.ID), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ = client.Do(req)
	_ = json.NewDecoder(resp.Body).Decode(&errResp)

	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
	require.Equal(t, api.ServiceError, errResp.Type)

	_ = resp.Body.Close()

	// User take with negative points
	userTakeReq = api.UserTakeRequest{Points: -1000}
	body, _ = json.Marshal(userTakeReq)
	req, _ = http.NewRequest("POST", baseURL+fmt.Sprintf("/user/%v/take", testUser.ID), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ = client.Do(req)
	_ = json.NewDecoder(resp.Body).Decode(&errResp)

	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
	require.Equal(t, api.ServiceError, errResp.Type)

	_ = resp.Body.Close()

	// User take with points more than balance
	userTakeReq = api.UserTakeRequest{Points: 2000}
	body, _ = json.Marshal(userTakeReq)
	req, _ = http.NewRequest("POST", baseURL+fmt.Sprintf("/user/%v/take", testUser.ID), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ = client.Do(req)
	_ = json.NewDecoder(resp.Body).Decode(&errResp)

	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
	require.Equal(t, api.ServiceError, errResp.Type)

	_ = resp.Body.Close()

	// Take balance of not existing user
	userTakeReq = api.UserTakeRequest{Points: 1000}
	body, _ = json.Marshal(userTakeReq)
	req, _ = http.NewRequest("POST", baseURL+"/user/-1/take", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ = client.Do(req)
	_ = json.NewDecoder(resp.Body).Decode(&errResp)

	require.Equal(t, http.StatusNotFound, resp.StatusCode)
	require.Equal(t, api.NotFoundError, errResp.Type)

	_ = resp.Body.Close()

	// User take balance with incorrect body
	req, _ = http.NewRequest("POST", baseURL+fmt.Sprintf("/user/%v/take", testUser.ID), bytes.NewBuffer([]byte(`{"points":"abc"}`)))
	req.Header.Set("Content-Type", "application/json")
	resp, _ = client.Do(req)
	_ = json.NewDecoder(resp.Body).Decode(&errResp)

	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
	require.Equal(t, api.BodyParseError, errResp.Type)

	_ = resp.Body.Close()

	// User fund with 0 points
	userFundReq := api.UserFundRequest{Points: 0}
	body, _ = json.Marshal(userFundReq)
	req, _ = http.NewRequest("POST", baseURL+fmt.Sprintf("/user/%v/fund", testUser.ID), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ = client.Do(req)
	_ = json.NewDecoder(resp.Body).Decode(&errResp)

	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
	require.Equal(t, api.ServiceError, errResp.Type)

	_ = resp.Body.Close()

	// User fund with negative points
	userFundReq = api.UserFundRequest{Points: -1000}
	body, _ = json.Marshal(userFundReq)
	req, _ = http.NewRequest("POST", baseURL+fmt.Sprintf("/user/%v/fund", testUser.ID), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ = client.Do(req)
	_ = json.NewDecoder(resp.Body).Decode(&errResp)

	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
	require.Equal(t, api.ServiceError, errResp.Type)

	_ = resp.Body.Close()

	// Fund balance of not existing user
	userFundReq = api.UserFundRequest{Points: 1000}
	body, _ = json.Marshal(userTakeReq)
	req, _ = http.NewRequest("POST", baseURL+"/user/-1/fund", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ = client.Do(req)
	_ = json.NewDecoder(resp.Body).Decode(&errResp)

	require.Equal(t, http.StatusNotFound, resp.StatusCode)
	require.Equal(t, api.NotFoundError, errResp.Type)

	_ = resp.Body.Close()

	// User fund balance with incorrect body
	req, _ = http.NewRequest("POST", baseURL+fmt.Sprintf("/user/%v/fund", testUser.ID), bytes.NewBuffer([]byte(`{"points":"abc"}`)))
	req.Header.Set("Content-Type", "application/json")
	resp, _ = client.Do(req)
	_ = json.NewDecoder(resp.Body).Decode(&errResp)

	require.Equal(t, http.StatusBadRequest, resp.StatusCode)
	require.Equal(t, api.BodyParseError, errResp.Type)

	_ = resp.Body.Close()
}

func TestFlow3(t *testing.T) {
	client := &http.Client{}

	// New tournament test
	newTReq := api.NewTournamentRequest{Name: "T1", Deposit: 1000}
	body, _ := json.Marshal(newTReq)
	req, _ := http.NewRequest("POST", baseURL+"/tournament", bytes.NewBuffer(body))
	resp, _ := client.Do(req)
	var newTResp api.NewTournamentResponse
	_ = json.NewDecoder(resp.Body).Decode(&newTResp)

	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.NotEqual(t, int64(0), newTResp.ID)
	require.Equal(t, newTReq.Name, newTResp.Name)
	require.Equal(t, newTReq.Deposit, newTResp.Deposit)
	require.Equal(t, int64(0), newTResp.Prize)

	_ = resp.Body.Close()

	// Get tournament test
	req, _ = http.NewRequest("GET", baseURL+fmt.Sprintf("/tournament/%v", newTResp.ID), nil)
	resp, _ = client.Do(req)
	var getTResp api.GetTournamentResponse
	_ = json.NewDecoder(resp.Body).Decode(&getTResp)

	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Equal(t, newTResp.ID, getTResp.ID)
	require.Equal(t, newTResp.Name, getTResp.Name)
	require.Equal(t, newTResp.Deposit, getTResp.Deposit)

	_ = resp.Body.Close()

	// Create users to join tournament

	var testUsers [5]model.User

	for i := 0; i < len(testUsers); i++ {
		newUserReq := api.NewUserRequest{Name: "test_user", Balance: 1000}
		body, _ := json.Marshal(newUserReq)
		req, _ := http.NewRequest("POST", baseURL+"/user", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := client.Do(req)
		var newUserResp api.NewUserResponse
		_ = json.NewDecoder(resp.Body).Decode(&newUserResp)

		testUsers[i] = model.User{ID: newUserResp.ID, Username: newUserResp.Name, Balance: newUserResp.Balance}

		_ = resp.Body.Close()
	}

	// Join created users to tournament

	for _, u := range testUsers {
		joinUserReq := api.JoinTournamentRequest{UserID: u.ID}
		body, _ := json.Marshal(joinUserReq)
		req, _ := http.NewRequest("POST", baseURL+fmt.Sprintf("/tournament/%v/join", newTResp.ID), bytes.NewBuffer(body))
		resp, _ := client.Do(req)
		var joinTResp api.JoinTournamentResponse
		_ = json.NewDecoder(resp.Body).Decode(&joinTResp)

		require.Equal(t, http.StatusOK, resp.StatusCode)

		isUserJoined := false
		for _, j := range joinTResp.Users {
			if u.ID == j.ID {
				isUserJoined = true
				break
			}
		}

		require.True(t, isUserJoined)

		_ = resp.Body.Close()
	}

	// Check and update users data
	for i, u := range testUsers {
		req, _ = http.NewRequest("GET", baseURL+fmt.Sprintf("/user/%v", u.ID), nil)
		req.Header.Set("Content-Type", "application/json")
		resp, _ = client.Do(req)
		var user api.GetUserResponse
		_ = json.NewDecoder(resp.Body).Decode(&user)

		require.Equal(t, u.Balance-newTResp.Deposit, user.Balance)

		testUsers[i].Balance = user.Balance

		_ = resp.Body.Close()
	}

	// Check tournament prize
	var tWithUsers api.GetTournamentResponse
	req, _ = http.NewRequest("GET", baseURL+fmt.Sprintf("/tournament/%v", newTResp.ID), nil)
	resp, _ = client.Do(req)
	_ = json.NewDecoder(resp.Body).Decode(&tWithUsers)

	require.Equal(t, tWithUsers.Deposit*int64(len(testUsers)), tWithUsers.Prize)

	_ = resp.Body.Close()

	// Check finish tournament
	req, _ = http.NewRequest("POST", baseURL+fmt.Sprintf("/tournament/%v/finish", newTResp.ID), nil)
	resp, _ = client.Do(req)
	var finishedTResp api.FinishTournamentResponse
	_ = json.NewDecoder(resp.Body).Decode(&finishedTResp)

	require.Equal(t, http.StatusOK, resp.StatusCode)

	isWinnerFound := false
	winnerID := int64(-1)
	winnerIdx := -1

	for idx, u := range testUsers {
		for _, p := range finishedTResp.Users {
			if p.ID == u.ID && p.Winner {
				isWinnerFound = true
				winnerID = p.ID
				winnerIdx = idx
				break
			}
		}
		if isWinnerFound {
			break
		}
	}

	require.True(t, isWinnerFound)

	_ = resp.Body.Close()

	// Check winner
	req, _ = http.NewRequest("GET", baseURL+fmt.Sprintf("/user/%v", winnerID), nil)
	req.Header.Set("Content-Type", "application/json")
	resp, _ = client.Do(req)
	var user api.GetUserResponse
	_ = json.NewDecoder(resp.Body).Decode(&user)

	require.Equal(t, testUsers[winnerIdx].Balance+tWithUsers.Prize, user.Balance)

	_ = resp.Body.Close()
}
