package service

import (
	"context"
	"errors"

	"github.com/777777miSSU7777777/gaming-website/model"
)

type Service interface {
	NewUser(string, int64) (model.User, error)
	GetUser(int64) (model.User, error)
	DeleteUser(int64) error
	UserTake(int64, int64) (model.User, error)
	UserFund(int64, int64) (model.User, error)
}

type Repository interface {
	NewUser(context.Context, string, int64) (int64, error)
	GetUserByID(context.Context, int64) (model.User, error)
	DeleteUserByID(context.Context, int64) error
	TakeUserBalanceByID(context.Context, int64, int64) error
	AddUserBalanceByID(context.Context, int64, int64) error
}

type ServiceImpl struct {
	repo Repository
}

func New(r Repository) ServiceImpl {
	return ServiceImpl{r}
}

func (s ServiceImpl) NewUser(username string, balance int64) (model.User, error) {
	id, err := s.repo.NewUser(context.Background(), username, balance)
	if err != nil {
		return model.User{}, err
	}

	user, err := s.repo.GetUserByID(context.Background(), id)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (s ServiceImpl) GetUser(id int64) (model.User, error) {
	user, err := s.repo.GetUserByID(context.Background(), id)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (s ServiceImpl) DeleteUser(id int64) error {
	err := s.repo.DeleteUserByID(context.Background(), id)
	if err != nil {
		return err
	}

	return nil
}

func (s ServiceImpl) UserTake(id int64, points int64) (model.User, error) {
	if points <= 0 {
		return model.User{}, errors.New("Can't take zero or negative points")
	}

	user, err := s.repo.GetUserByID(context.Background(), id)
	if err != nil {
		return model.User{}, err
	}

	if points > user.Balance {
		return model.User{}, errors.New("Balance isn't enough for taking points")
	}

	err = s.repo.TakeUserBalanceByID(context.Background(), id, points)
	if err != nil {
		return model.User{}, err
	}

	user, err = s.repo.GetUserByID(context.Background(), id)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (s ServiceImpl) UserFund(id int64, points int64) (model.User, error) {
	if points <= 0 {
		return model.User{}, errors.New("Can't fund zero or negative points")
	}

	user, err := s.repo.GetUserByID(context.Background(), id)
	if err != nil {
		return model.User{}, err
	}

	err = s.repo.AddUserBalanceByID(context.Background(), id, points)
	if err != nil {
		return model.User{}, err
	}

	user.Balance += points
	return user, nil
}
