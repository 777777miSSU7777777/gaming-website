package service

import (
	"context"
	"errors"

	"github.com/777777miSSU7777777/gaming-website/repository"

	"github.com/777777miSSU7777777/gaming-website/model"
)

type Service struct {
	repo repository.Repository
}

var TournamentFinishedError = errors.New("tournament already finished")
var TournamentCanceledError = errors.New("tournament already canceled")

func New(r repository.Repository) Service {
	return Service{r}
}

func (s Service) NewUser(username string, balance int64) (model.User, error) {
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

func (s Service) GetUser(id int64) (model.User, error) {
	user, err := s.repo.GetUserByID(context.Background(), id)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (s Service) DeleteUser(id int64) error {
	err := s.repo.DeleteUserByID(context.Background(), id)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) UserTake(id int64, points int64) (model.User, error) {
	if points <= 0 {
		return model.User{}, errors.New("cant take zero or negative points")
	}

	user, err := s.repo.GetUserByID(context.Background(), id)
	if err != nil {
		return model.User{}, err
	}

	if points > user.Balance {
		return model.User{}, errors.New("balance isnt enough for taking points")
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

func (s Service) UserFund(id int64, points int64) (model.User, error) {
	if points <= 0 {
		return model.User{}, errors.New("cant fund zero or negative points")
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

func (s Service) NewTournament(name string, deposit int64) (model.Tournament, error) {
	id, err := s.repo.NewTournament(context.Background(), name, deposit)
	if err != nil {
		return model.Tournament{}, err
	}

	tournament, err := s.repo.GetTournamentByID(context.Background(), id)
	if err != nil {
		return model.Tournament{}, err
	}

	return tournament, nil
}

func (s Service) GetTournament(id int64) (model.Tournament, []model.User, error) {
	tournament, err := s.repo.GetTournamentByID(context.Background(), id)
	if err != nil {
		return model.Tournament{}, nil, err
	}

	users, err := s.repo.GetTournamentUsers(context.Background(), id)
	if err != nil {
		if err.Error() != "users not found error" {
			return model.Tournament{}, nil, err
		}
	}

	return tournament, users, nil
}

func checkTournamentStatus(t model.Tournament) error {
	if t.Status == "Finished" {
		return TournamentFinishedError
	}
	if t.Status == "Canceled" {
		return TournamentCanceledError
	}
	return nil
}

func (s Service) JoinTournament(tournamentID int64, userID int64) (model.Tournament, []model.User, error) {
	tournament, err := s.repo.GetTournamentByID(context.Background(), tournamentID)
	if err != nil {
		return model.Tournament{}, nil, err
	}

	err = checkTournamentStatus(tournament)
	if err != nil {
		return model.Tournament{}, nil, err
	}

	user, err := s.repo.GetUserByID(context.Background(), userID)
	if err != nil {
		return model.Tournament{}, nil, err
	}

	err = s.repo.CheckUserJoinTournament(context.Background(), tournamentID, userID)
	if err != nil {
		return model.Tournament{}, nil, err
	}

	if user.Balance < tournament.Deposit {
		return model.Tournament{}, nil, errors.New("user balance isnt enough to join this tournament")
	}

	err = s.repo.JoinUserTournament(context.Background(), tournamentID, userID)
	if err != nil {
		return model.Tournament{}, nil, err
	}

	tournament, err = s.repo.GetTournamentByID(context.Background(), tournamentID)
	if err != nil {
		return model.Tournament{}, nil, err
	}

	users, err := s.repo.GetTournamentUsers(context.Background(), tournamentID)
	if err != nil {
		return model.Tournament{}, nil, err
	}

	return tournament, users, nil
}

func (s Service) FinishTournament(id int64) (model.Tournament, []model.User, error) {
	tournament, err := s.repo.GetTournamentByID(context.Background(), id)
	if err != nil {
		return model.Tournament{}, nil, err
	}

	err = checkTournamentStatus(tournament)
	if err != nil {
		return model.Tournament{}, nil, err
	}

	err = s.repo.FinishTournament(context.Background(), id)

	if err != nil {
		return model.Tournament{}, nil, err
	}

	tournament, err = s.repo.GetTournamentByID(context.Background(), id)
	if err != nil {
		return model.Tournament{}, nil, err
	}

	users, err := s.repo.GetTournamentUsers(context.Background(), id)
	if err != nil {
		return model.Tournament{}, nil, err
	}

	return tournament, users, nil
}

func (s Service) CancelTournament(id int64) error {
	tournament, err := s.repo.GetTournamentByID(context.Background(), id)
	if err != nil {
		return err
	}

	err = checkTournamentStatus(tournament)
	if err != nil {
		return err
	}

	err = s.repo.CancelTournament(context.Background(), id)
	if err != nil {
		return err
	}

	return nil
}
