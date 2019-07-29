package userservice

import (
	"context"
	"fmt"

	"github.com/777777miSSU7777777/gaming-website/pkg/entity"
	"github.com/777777miSSU7777777/gaming-website/pkg/repository/userrepository"
)

type UserService interface {
	NewUser(string, int64) (entity.User, error)
	GetUser(int64) (entity.User, error)
	DeleteUser(int64) error
	UserTake(int64, int64) (entity.User, error)
	UserFund(int64, int64) (entity.User, error)
}

type service struct {
	repo userrepository.UserRepository
}

func New(r userrepository.UserRepository) UserService {
	return &service{r}
}

func (s service) NewUser(username string, balance int64) (entity.User, error) {
	if balance < 0 {
		return entity.User{}, fmt.Errorf("Balance cant be negative")
	}
	id, err := s.repo.New(context.Background(), username, balance)
	if err != nil {
		return entity.User{}, err
	}
	user, err := s.repo.GetByID(context.Background(), id)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (s service) GetUser(id int64) (entity.User, error) {
	user, err := s.repo.GetByID(context.Background(), id)
	if err != nil {
		return entity.User{}, err
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

func (s service) UserTake(id int64, points int64) (entity.User, error) {
	user, err := s.repo.GetByID(context.Background(), id)
	if err != nil {
		return entity.User{}, err
	}
	if user.Balance < points {
		return entity.User{}, fmt.Errorf("User doesnt have enough balance")
	}
	err = s.repo.UpdateByID(context.Background(), id, user.Username, user.Balance-points)
	if err != nil {
		return entity.User{}, err
	}
	user, err = s.repo.GetByID(context.Background(), id)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (s service) UserFund(id int64, points int64) (entity.User, error) {
	user, err := s.repo.GetByID(context.Background(), id)
	if err != nil {
		return entity.User{}, err
	}
	err = s.repo.UpdateByID(context.Background(), id, user.Username, user.Balance+points)
	if err != nil {
		return entity.User{}, err
	}
	user, err = s.repo.GetByID(context.Background(), id)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}
