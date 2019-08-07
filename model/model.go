package model

import (
	"fmt"
)

type User struct {
	ID       int64
	Username string
	Balance  int64
}

func (u User) Validate() error {
	if u.Username == "" {
		return fmt.Errorf("username is empty")
	}
	if u.Balance < 0 {
		return fmt.Errorf("user balance is negative")
	}
	return nil
}

type Tournament struct {
	ID             int64
	TournamentName string
	Status         string
	Deposit        int64
	Prize          int64
	WinnerID       int64
}

func (t Tournament) Validate() error {
	if t.TournamentName == "" {
		return fmt.Errorf("tournament name is empty")
	}
	if t.Deposit <= 0 {
		return fmt.Errorf("tournament deposit is zero or negative")
	}
	return nil
}
