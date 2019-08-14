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

type APIClient struct {
	client  *http.Client
	baseURL string
}

func (c APIClient) NewUser(name string, balance int64) (int, api.NewUserResponse) {
	body, _ := json.Marshal(api.NewUserRequest{Name: name, Balance: balance})
	req, _ := http.NewRequest("POST", c.baseURL+"/user", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := c.client.Do(req)
	var newUser api.NewUserResponse
	_ = json.NewDecoder(resp.Body).Decode(&newUser)
	_ = resp.Body.Close()

	return resp.StatusCode, newUser
}

func (c APIClient) GetUser(id int64) (int, api.GetUserResponse) {
	req, _ := http.NewRequest("GET", c.baseURL+fmt.Sprintf("/user/%v", id), nil)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := c.client.Do(req)
	var getUser api.GetUserResponse
	_ = json.NewDecoder(resp.Body).Decode(&getUser)
	_ = resp.Body.Close()

	return resp.StatusCode, getUser
}

func (c APIClient) DeleteUser(id int64) int {
	req, _ := http.NewRequest("DELETE", c.baseURL+fmt.Sprintf("/user/%v", id), nil)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := c.client.Do(req)
	_ = resp.Body.Close()

	return resp.StatusCode
}

func (c APIClient) UserTake(id int64, points int64) (int, api.UserTakeResponse) {
	body, _ := json.Marshal(api.UserTakeRequest{Points: points})
	req, _ := http.NewRequest("POST", c.baseURL+fmt.Sprintf("/user/%v/take", id), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := c.client.Do(req)
	var userTake api.UserTakeResponse
	_ = json.NewDecoder(resp.Body).Decode(&userTake)
	_ = resp.Body.Close()

	return resp.StatusCode, userTake
}

func (c APIClient) UserFund(id int64, points int64) (int, api.UserFundResponse) {
	body, _ := json.Marshal(api.UserFundRequest{Points: points})
	req, _ := http.NewRequest("POST", c.baseURL+fmt.Sprintf("/user/%v/fund", id), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := c.client.Do(req)
	var userFund api.UserFundResponse
	_ = json.NewDecoder(resp.Body).Decode(&userFund)
	_ = resp.Body.Close()

	return resp.StatusCode, userFund
}

func (c APIClient) UserErr(method string, url string, body []byte) (int, api.ErrorResponse) {
	req, _ := http.NewRequest(method, c.baseURL+url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := c.client.Do(req)
	var err api.ErrorResponse
	_ = json.NewDecoder(resp.Body).Decode(&err)
	_ = resp.Body.Close()

	return resp.StatusCode, err
}

func (c APIClient) NewTournament(name string, deposit int64) (int, api.NewTournamentResponse) {
	body, _ := json.Marshal(api.NewTournamentRequest{Name: name, Deposit: deposit})
	req, _ := http.NewRequest("POST", c.baseURL+"/tournament", bytes.NewBuffer(body))
	resp, _ := c.client.Do(req)
	var newTournament api.NewTournamentResponse
	_ = json.NewDecoder(resp.Body).Decode(&newTournament)
	_ = resp.Body.Close()

	return resp.StatusCode, newTournament
}

func (c APIClient) GetTournament(id int64) (int, api.GetTournamentResponse) {
	req, _ := http.NewRequest("GET", c.baseURL+fmt.Sprintf("/tournament/%v", id), nil)
	resp, _ := c.client.Do(req)
	var getTournament api.GetTournamentResponse
	_ = json.NewDecoder(resp.Body).Decode(&getTournament)
	_ = resp.Body.Close()

	return resp.StatusCode, getTournament
}

func (c APIClient) JoinUserTournament(tID int64, uID int64) (int, api.JoinTournamentResponse) {
	body, _ := json.Marshal(api.JoinTournamentRequest{UserID: uID})
	req, _ := http.NewRequest("POST", c.baseURL+fmt.Sprintf("/tournament/%v/join", tID), bytes.NewBuffer(body))
	resp, _ := c.client.Do(req)
	var joinTournament api.JoinTournamentResponse
	_ = json.NewDecoder(resp.Body).Decode(&joinTournament)
	_ = resp.Body.Close()

	return resp.StatusCode, joinTournament
}

func (c APIClient) FinishTournament(id int64) (int, api.FinishTournamentResponse) {
	req, _ := http.NewRequest("POST", c.baseURL+fmt.Sprintf("/tournament/%v/finish", id), nil)
	resp, _ := c.client.Do(req)
	var finishedTournament api.FinishTournamentResponse
	_ = json.NewDecoder(resp.Body).Decode(&finishedTournament)
	_ = resp.Body.Close()

	return resp.StatusCode, finishedTournament
}

func (c APIClient) CancelTournament(id int64) int {
	req, _ := http.NewRequest("DELETE", c.baseURL+fmt.Sprintf("/tournament/%v", id), nil)
	resp, _ := c.client.Do(req)
	_ = resp.Body.Close()

	return resp.StatusCode
}

func TestUserApiWithCorrectData(t *testing.T) {
	client := APIClient{client: &http.Client{}, baseURL: "http://127.0.0.1:8080"}

	// New user test
	code, newUser := client.NewUser("test_user", 1000)

	require.Equal(t, http.StatusOK, code)
	require.NotEqual(t, int64(0), newUser.ID)
	require.Equal(t, "test_user", newUser.Name)
	require.Equal(t, int64(1000), newUser.Balance)

	// Get user test
	code, getUser := client.GetUser(newUser.ID)

	require.Equal(t, http.StatusOK, code)
	require.Equal(t, newUser.ID, getUser.ID)
	require.Equal(t, newUser.Name, getUser.Name)
	require.Equal(t, newUser.Balance, getUser.Balance)

	// Take user balance test
	code, userTake := client.UserTake(newUser.ID, 500)

	require.Equal(t, http.StatusOK, code)
	require.Equal(t, newUser.ID, userTake.ID)
	require.Equal(t, newUser.Name, userTake.Name)
	require.Equal(t, newUser.Balance-500, userTake.Balance)

	// Fund user balance test
	code, userFund := client.UserFund(newUser.ID, 1000)

	require.Equal(t, http.StatusOK, code)
	require.Equal(t, newUser.ID, userFund.ID)
	require.Equal(t, newUser.Name, userFund.Name)
	require.Equal(t, userTake.Balance+1000, userFund.Balance)

	// Delete user test
	code = client.DeleteUser(newUser.ID)

	require.Equal(t, http.StatusOK, code)
}

func TestUserApiWithIncorrectData(t *testing.T) {
	client := APIClient{client: &http.Client{}, baseURL: "http://127.0.0.1:8080"}

	// New user with empty name
	code, err := client.UserErr("POST", "/user", []byte(`{"name": "", "balance": 1000}`))

	require.Equal(t, http.StatusBadRequest, code)
	require.Equal(t, api.ValidationError, err.Type)

	// New user with negative balance

	code, err = client.UserErr("POST", "/user", []byte(`{"name": "test_user", "balance": -1000}`))

	require.Equal(t, http.StatusBadRequest, code)
	require.Equal(t, api.ValidationError, err.Type)

	// New user with incorrect body
	code, err = client.UserErr("POST", "/user", []byte(`{"name": 123, "balance": "abc"}`))

	require.Equal(t, http.StatusBadRequest, code)
	require.Equal(t, api.BodyParseError, err.Type)

	// Get not existing user
	code, err = client.UserErr("GET", "/user/-1", nil)

	require.Equal(t, http.StatusNotFound, code)
	require.Equal(t, api.NotFoundError, err.Type)

	// Get with invalid id
	code, err = client.UserErr("GET", "/user/abc", nil)

	require.Equal(t, http.StatusBadRequest, code)
	require.Equal(t, api.IDParseError, err.Type)

	// Delete not existing user
	code, err = client.UserErr("DELETE", "/user/-1", nil)

	require.Equal(t, http.StatusNotFound, code)
	require.Equal(t, api.NotFoundError, err.Type)

	// Create correct user to test take and fund actions with incorrect data
	_, testUser := client.NewUser("test_user", 1000)

	// User take with 0 points
	code, err = client.UserErr("POST", fmt.Sprintf("/user/%v/take", testUser.ID), []byte(`{"points": 0}`))

	require.Equal(t, http.StatusBadRequest, code)
	require.Equal(t, api.ServiceError, err.Type)

	// User take with negative points
	code, err = client.UserErr("POST", fmt.Sprintf("/user/%v/take", testUser.ID), []byte(`{"points": -1000}`))

	require.Equal(t, http.StatusBadRequest, code)
	require.Equal(t, api.ServiceError, err.Type)

	// User take with points more than balance
	code, err = client.UserErr("POST", fmt.Sprintf("/user/%v/take", testUser.ID), []byte(`{"points": 2000}`))

	require.Equal(t, http.StatusBadRequest, code)
	require.Equal(t, api.ServiceError, err.Type)

	// Take balance of not existing user
	code, err = client.UserErr("POST", "/user/-1/take", []byte(`{"points": 1000}`))

	require.Equal(t, http.StatusNotFound, code)
	require.Equal(t, api.NotFoundError, err.Type)

	// User take balance with incorrect body
	code, err = client.UserErr("POST", fmt.Sprintf("/user/%v/take", testUser.ID), []byte(`{"points": "abc"}`))

	require.Equal(t, http.StatusBadRequest, code)
	require.Equal(t, api.BodyParseError, err.Type)

	// User fund with 0 points
	code, err = client.UserErr("POST", fmt.Sprintf("/user/%v/fund", testUser.ID), []byte(`{"points": 0}`))

	require.Equal(t, http.StatusBadRequest, code)
	require.Equal(t, api.ServiceError, err.Type)

	// User fund with negative points
	code, err = client.UserErr("POST", fmt.Sprintf("/user/%v/fund", testUser.ID), []byte(`{"points": -1000}`))

	require.Equal(t, http.StatusBadRequest, code)
	require.Equal(t, api.ServiceError, err.Type)

	// Fund balance of not existing user
	code, err = client.UserErr("POST", "/user/-1/fund", []byte(`{"points": 1000}`))

	require.Equal(t, http.StatusNotFound, code)
	require.Equal(t, api.NotFoundError, err.Type)

	// User fund balance with incorrect body
	code, err = client.UserErr("POST", fmt.Sprintf("/user/%v/fund", testUser.ID), []byte(`{"points": "abc"}`))

	require.Equal(t, http.StatusBadRequest, code)
	require.Equal(t, api.BodyParseError, err.Type)
}

func TestTournamentFinish(t *testing.T) {
	client := APIClient{client: &http.Client{}, baseURL: "http://127.0.0.1:8080"}

	// New tournament test
	code, newTournament := client.NewTournament("test_tournament", 1000)

	require.Equal(t, http.StatusOK, code)
	require.NotEqual(t, int64(0), newTournament.ID)
	require.Equal(t, "test_tournament", newTournament.Name)
	require.Equal(t, int64(1000), newTournament.Deposit)
	require.Equal(t, int64(0), newTournament.Prize)

	// Get tournament test
	code, getTournament := client.GetTournament(newTournament.ID)

	require.Equal(t, http.StatusOK, code)
	require.Equal(t, newTournament.ID, getTournament.ID)
	require.Equal(t, newTournament.Name, getTournament.Name)
	require.Equal(t, newTournament.Deposit, getTournament.Deposit)

	// Create users to join tournament
	var testUsers [5]model.User

	for i := 0; i < len(testUsers); i++ {
		_, newUser := client.NewUser("test_user", 1000)
		testUsers[i] = model.User{ID: newUser.ID, Username: newUser.Name, Balance: newUser.Balance}
	}

	// Join created users to tournament
	for _, u := range testUsers {
		code, joinTournament := client.JoinUserTournament(newTournament.ID, u.ID)

		require.Equal(t, http.StatusOK, code)

		isUserJoined := false
		for _, j := range joinTournament.Users {
			if u.ID == j.ID {
				isUserJoined = true
				break
			}
		}

		require.True(t, isUserJoined)
	}

	// Check and update users data
	for i, u := range testUsers {
		_, user := client.GetUser(u.ID)

		require.Equal(t, u.Balance-newTournament.Deposit, user.Balance)

		testUsers[i].Balance = user.Balance
	}

	// Check tournament prize
	code, tournamentWithUsers := client.GetTournament(newTournament.ID)

	require.Equal(t, tournamentWithUsers.Deposit*int64(len(testUsers)), tournamentWithUsers.Prize)

	// Check finish tournament
	code, finishedTournament := client.FinishTournament(newTournament.ID)

	require.Equal(t, http.StatusOK, code)

	isWinnerFound := false
	winnerID := int64(-1)
	winnerIdx := -1

	for idx, u := range testUsers {
		for _, p := range finishedTournament.Users {
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

	// Check winner
	_, winner := client.GetUser(winnerID)

	require.Equal(t, testUsers[winnerIdx].Balance+tournamentWithUsers.Prize, winner.Balance)
}

func TestTournamentCancel(t *testing.T) {
	client := APIClient{client: &http.Client{}, baseURL: "http://127.0.0.1:8080"}

	// Create new tournament
	code, newTournament := client.NewTournament("test_tournament", 1000)

	// Create users to join tournament
	var testUsers [5]model.User

	for i := 0; i < len(testUsers); i++ {
		_, newUser := client.NewUser("test_user", 1000)
		testUsers[i] = model.User{ID: newUser.ID, Username: newUser.Name, Balance: newUser.Balance}
	}

	// Join created users to tournament
	for _, u := range testUsers {
		_, _ = client.JoinUserTournament(newTournament.ID, u.ID)
	}

	// Check cancel tournament
	code = client.CancelTournament(newTournament.ID)

	require.Equal(t, http.StatusOK, code)

	// Check returned money from canceled tournament
	for _, u := range testUsers {
		_, user := client.GetUser(u.ID)

		require.Equal(t, u.Balance, user.Balance)
	}
}
