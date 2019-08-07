package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/777777miSSU7777777/gaming-website/model"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return Repository{db}
}

func (r Repository) NewUser(ctx context.Context, name string, balance int64) (int64, error) {
	result, err := r.db.Exec("INSERT INTO USERS (USERNAME, BALANCE) VALUES(?, ?)", name, balance)
	if err != nil {
		return -1, fmt.Errorf("add user error: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("add user error: %v", err)
	}

	return id, nil
}

func (r Repository) GetUserByID(ctx context.Context, id int64) (model.User, error) {
	row := r.db.QueryRow("SELECT * FROM USERS WHERE USER_ID=?", id)
	user := model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Balance)
	if err != nil {
		return model.User{}, fmt.Errorf("get user error: %v", err)
	}

	return user, nil
}

func (r Repository) DeleteUserByID(ctx context.Context, id int64) error {
	res, err := r.db.Exec("DELETE FROM USERS WHERE USER_ID=?", id)
	if err != nil {
		return fmt.Errorf("delete user error: %v", err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("delete user error: %v", err)
	}
	if count == 0 {
		return errors.New("user not found error")
	}

	return nil
}

func (r Repository) TakeUserBalanceByID(ctx context.Context, id int64, points int64) error {
	row := r.db.QueryRow("SELECT * FROM USERS WHERE USER_ID=?", id)
	user := model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Balance)
	if err != nil {
		return fmt.Errorf("take user balance error: %v", err)
	}

	_, err = r.db.Exec("UPDATE USERS SET BALANCE=? WHERE USER_ID=?", user.Balance-points, id)
	if err != nil {
		return fmt.Errorf("take user balance error: %v", err)
	}

	return nil
}

func (r Repository) AddUserBalanceByID(ctx context.Context, id int64, points int64) error {
	row := r.db.QueryRow("SELECT * FROM USERS WHERE USER_ID=?", id)
	user := model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Balance)
	if err != nil {
		return fmt.Errorf("add user balance error: %v", err)
	}
	_, err = r.db.Exec("UPDATE USERS SET BALANCE=? WHERE USER_ID=?", user.Balance+points, id)
	if err != nil {
		return fmt.Errorf("add user balance error: %v", err)
	}

	return nil
}

func (r Repository) NewTournament(ctx context.Context, name string, deposit int64) (int64, error) {
	result, err := r.db.Exec("INSERT INTO TOURNAMENTS(TOURNAMENT_NAME, DEPOSIT) VALUES(?,?)", name, deposit)
	if err != nil {
		return -1, fmt.Errorf("new tournament error: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("new tournament error: %v", err)
	}

	return id, nil
}

func (r Repository) GetTournamentByID(ctx context.Context, id int64) (model.Tournament, error) {
	row := r.db.QueryRow("SELECT * FROM TOURNAMENTS WHERE ID=?", id)
	tournament := model.Tournament{}
	err := row.Scan(&tournament.ID, &tournament.TournamentName, &tournament.Deposit, &tournament.Prize)
	if err != nil {
		return model.Tournament{}, fmt.Errorf("get tournament error: %v", err)
	}

	return tournament, nil
}

func (r Repository) AddUserToTournament(ctx context.Context, tournamentID int64, userID int64) error {
	_, err := r.db.Exec("INSERT INTO MTM_USER_TOURNAMENT(TOURNAMENT_ID, USER_ID) VALUES (?,?)", tournamentID, userID)
	if err != nil {
		return fmt.Errorf("add user to tournament error: %v", err)
	}

	return nil
}

func (r Repository) DeleteTournamentByID(ctx context.Context, id int64) error {
	res, err := r.db.Exec("DELETE FROM TOURNAMENTS WHERE ID=?", id)
	if err != nil {
		return fmt.Errorf("delete tournament error: %v", err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("delete tournament error %v", err)
	}
	if count == 0 {
		return errors.New("tournament not found error")
	}

	return nil
}
