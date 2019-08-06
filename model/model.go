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
	if u.Username == "" && u.Balance < 0 {
		return fmt.Errorf("Username can't be empty and balance can't be negative")
	}
	if u.Username == "" {
		return fmt.Errorf("Username can't be empty")
	}
	if u.Balance < 0 {
		return fmt.Errorf("User balance can't be negative")
	}
	return nil
}

type Tournament struct {
	ID             int64
	TournamentName string
	Deposit        int64
	Prize          int64
}

func (t Tournament) Validate() error {
	if t.TournamentName == "" && t.Deposit <= 0 {
		return fmt.Errorf("Tournament name can't be empty and deposit can't be zero or negative")
	}
	if t.TournamentName == "" {
		return fmt.Errorf("Tournament name can't be empty")
	}
	if t.Deposit <= 0 {
		return fmt.Errorf("Tournament deposit can't be zero or negative")
	}
	return nil
}
