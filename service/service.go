package service

import (
	"context"
	"fmt"

	"github.com/777777miSSU7777777/gaming-website/model"
)

type UserService interface {
	NewUser(string, int64) (model.User, error)
	GetUser(int64) (model.User, error)
	DeleteUser(int64) error
	UserTake(int64, int64) (model.User, error)
	UserFund(int64, int64) (model.User, error)
}

type UserRepository interface {
	New(context.Context, string, int64) (int64, error)
	GetByID(context.Context, int64) (model.User, error)
	DeleteByID(context.Context, int64) error
	TakeBalanceByID(context.Context, int64, int64) error
	AddBalanceByID(context.Context, int64, int64) error
}

type service struct {
	repo UserRepository
}

func New(r UserRepository) UserService {
	return &service{r}
}

func (s service) NewUser(username string, balance int64) (model.User, error) {
	id, err := s.repo.New(context.Background(), username, balance)
	if err != nil {
		return model.User{}, err
	}

	user, err := s.repo.GetByID(context.Background(), id)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (s service) GetUser(id int64) (model.User, error) {
	user, err := s.repo.GetByID(context.Background(), id)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (s service) DeleteUser(id int64) error {
	err := s.repo.DeleteByID(context.Background(), id)
	if err != nil {
		return err
	}

	return nil
}

func (s service) UserTake(id int64, points int64) (model.User, error) {
	if points <= 0 {
		return model.User{}, fmt.Errorf("Cant take zero or negative points")
	}

	user, err := s.repo.GetByID(context.Background(), id)
	if err != nil {
		return model.User{}, err
	}

	if points > user.Balance {
		return model.User{}, fmt.Errorf("Balance isnt enough for taking points")
	}

	err = s.repo.TakeBalanceByID(context.Background(), id, points)
	if err != nil {
		return model.User{}, err
	}

	user, err = s.repo.GetByID(context.Background(), id)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (s service) UserFund(id int64, points int64) (model.User, error) {
	if points <= 0 {
		return model.User{}, fmt.Errorf("Cant fund zero or negative points")
	}

	user, err := s.repo.GetByID(context.Background(), id)
	if err != nil {
		return model.User{}, err
	}

	err = s.repo.AddBalanceByID(context.Background(), id, points)
	if err != nil {
		return model.User{}, err
	}

	user.Balance += points
	return user, nil
}
