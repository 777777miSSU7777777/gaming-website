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
