package api

type NewUserResponse struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Balance int64  `json:"balance"`
}

type GetUserResponse NewUserResponse

type DeleteUserResponse struct {
}

type UserTakeResponse NewUserResponse

type UserFundResponse NewUserResponse

type NotFinishedUser struct {
	ID   int64  `json:"userId"`
	Name string `json:"name"`
}

type NewTournamentResponse struct {
	ID      int64             `json:"id"`
	Name    string            `json:"name"`
	Deposit int64             `json:"deposit"`
	Prize   int64             `json:"prize"`
	Users   []NotFinishedUser `json:"users"`
}

type GetTournamentResponse NewTournamentResponse

type JoinTournamentResponse NewTournamentResponse

type FinishedUser struct {
	ID     int64  `json:"userId"`
	Name   string `json:"name"`
	Winner bool   `json:"winner"`
}

type FinishTournamentResponse struct {
	ID      int64          `json:"id"`
	Name    string         `json:"name"`
	Deposit int64          `json:"deposit"`
	Prize   int64          `json:"prize"`
	Users   []FinishedUser `json:"users"`
}

type CancelTournamentResponse struct {
}
