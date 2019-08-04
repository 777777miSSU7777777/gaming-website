package userservice

import (
	"context"

	"github.com/777777miSSU7777777/gaming-website/model"
	"github.com/777777miSSU7777777/gaming-website/repository"
)

type UserService interface {
	NewUser(string, int64) (entity.User, error)
	GetUser(int64) (entity.User, error)
	DeleteUser(int64) error
	UserTake(int64, int64) (entity.User, error)
	UserFund(int64, int64) (entity.User, error)
}

type UserRepository interface {
	New(context.Context, string, int64) (int64, error)
	GetByID(context.Context, int64) (model.User, error)
	DeleteByID(context.Context, int64) error
	UpdateByID(context.Context, int64, string, int64) error
}

type service struct {
	repo UserRepository
}

func New(r userrepository.UserRepository) UserService {
	return &service{r}
}

func (s service) NewUser(username string, balance int64) (model.User, error) {
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

func (s service) GetUser(id int64) (model.User, error) {
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

func (s service) UserTake(id int64, points int64) (model.User, error) {
	user, err := s.repo.GetByID(context.Background(), id)
	if err != nil {
		return model.User{}, err
	}

	err = s.repo.UpdateByID(context.Background(), id, user.Username, user.Balance-points)
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
	if points < 0 {
		return model.User{}, fmt.Errorf("")
	}

	user, err := s.repo.GetByID(context.Background(), id)
	if err != nil {
		return model.User{}, err
	}

	err = s.repo.UpdateByID(context.Background(), id, user.Username, user.Balance+points)
	if err != nil {
		return model.User{}, err
	}

	user.Balance += points
	return user, nil
}
