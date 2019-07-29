package userservice

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
