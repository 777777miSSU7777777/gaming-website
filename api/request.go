package api

type NewUserRequest struct {
	Name    string `json:"name"`
	Balance int64  `json:"balance"`
}

type GetUserRequest struct {
	ID int64
}

type DeleteUserRequest GetUserRequest

type UserTakeRequest struct {
	ID     int64
	Points int64 `json:"points"`
}

type UserFundRequest UserTakeRequest

type NewTournamentRequest struct {
	Name    string `json:"name"`
	Deposit int64  `json:"deposit"`
}

type GetTournamentRequest struct {
	ID int64
}

type JoinTournamentRequest struct {
	TournamentID int64
	UserID       int64 `json:"userId"`
}

type FinishTournamentRequest struct {
	ID int64
}

type CancelTournamentRequest struct {
	ID int64
}
