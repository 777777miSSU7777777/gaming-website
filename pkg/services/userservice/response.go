package userservice

type NewUserResponse struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Balance int64  `json:"balance"`
}

type GetUserResponse NewUserResponse

type DeleteUserResponse struct {
}

type UserTakeResponse NewUserResponse
