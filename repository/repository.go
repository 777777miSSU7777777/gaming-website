package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/777777miSSU7777777/gaming-website/model"
)

var UserNotFoundError = errors.New("user not found error")
var TournamentNotFoundError = errors.New("tournament not found error")

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return Repository{db}
}

func (r Repository) StartTx() (*sql.Tx, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("sql transaction error: %v", err)
	}

	return tx, nil
}

func (r Repository) Commit(tx *sql.Tx) error {
	err := tx.Commit()
	if err != nil {
		return fmt.Errorf("sql transaction commit error: %v", err)
	}

	return nil
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
		if err == sql.ErrNoRows {
			return model.User{}, err
		}
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
		return UserNotFoundError
	}

	return nil
}

func (r Repository) TakeUserBalanceByID(ctx context.Context, id int64, points int64) error {
	res, err := r.db.Exec("UPDATE USERS SET BALANCE=BALANCE-? WHERE USER_ID=?", points, id)
	if err != nil {
		return fmt.Errorf("take user balance error: %v", err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("take user balance error: %v", err)
	}
	if count == 0 {
		return UserNotFoundError
	}

	return nil
}

func (r Repository) AddUserBalanceByID(ctx context.Context, id int64, points int64) error {
	res, err := r.db.Exec("UPDATE USERS SET BALANCE=BALANCE+? WHERE USER_ID=?", points, id)
	if err != nil {
		return fmt.Errorf("add user balance error: %v", err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("add user balance error: %v", err)
	}
	if count == 0 {
		return UserNotFoundError
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
	row := r.db.QueryRow("SELECT * FROM TOURNAMENTS WHERE TOURNAMENT_ID=?", id)
	tournament := model.Tournament{}
	err := row.Scan(&tournament.ID, &tournament.TournamentName, &tournament.Status, &tournament.Deposit, &tournament.Prize, &tournament.WinnerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Tournament{}, TournamentNotFoundError
		}
		return model.Tournament{}, fmt.Errorf("get tournament error: %v", err)
	}

	return tournament, nil
}

func (r Repository) AddUserToTournament(ctx context.Context, tournamentID int64, userID int64) error {
	_, err := r.db.Query("SELECT * FROM TOURNAMENTS WHERE TOURNAMENT_ID=?", tournamentID)
	if err == sql.ErrNoRows {
		return TournamentNotFoundError
	}

	_, err = r.db.Query("SELECT * FROM USERS WHERE USER_ID=?", userID)
	if err == sql.ErrNoRows {
		return UserNotFoundError
	}

	var mtm model.MTMUserTournament
	row := r.db.QueryRow("SELECT * FROM MTM_USER_TOURNAMENT WHERE TOURNAMENT_ID=? AND USER_ID=?", tournamentID, userID)
	err = row.Scan(&mtm.ID, &mtm.TournamentID, &mtm.UserID)
	if err == nil {
		return fmt.Errorf("user already joined this tournament error")
	}

	_, err = r.db.Exec("INSERT INTO MTM_USER_TOURNAMENT(TOURNAMENT_ID, USER_ID) VALUES (?,?)", tournamentID, userID)
	if err != nil {
		return fmt.Errorf("add user to tournament error: %v", err)
	}

	return nil
}

func (r Repository) IncreaseTournamentPrize(ctx context.Context, id int64, points int64) error {
	res, err := r.db.Exec("UPDATE TOURNAMENTS SET PRIZE=PRIZE + ? WHERE TOURNAMENT_ID=?", points, id)
	if err != nil {
		return fmt.Errorf("increase prize error: %v", err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("increase prize error: %v", err)
	}
	if count == 0 {
		return TournamentNotFoundError
	}

	return nil
}

func (r Repository) SetTournamentStatusByID(ctx context.Context, id int64, status string) error {
	res, err := r.db.Exec("UPDATE TOURNAMENTS SET TOURNAMENT_STATUS=? WHERE TOURNAMENT_ID=?", status, id)
	if err != nil {
		return fmt.Errorf("set tournament status error: %v", err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("set tournament status error: %v", err)
	}
	if count == 0 {
		return TournamentNotFoundError
	}

	return nil
}

func (r Repository) SetTournamentWinner(ctx context.Context, tournamentID int64, winnerID int64) error {
	_, err := r.db.Query("SELECT * FROM MTM_USER_TOURNAMENT WHERE TOURNAMENT_ID=? AND USER_ID=?", tournamentID, winnerID)
	if err == sql.ErrNoRows {
		fmt.Errorf("user or tournament not found error: %v", err)
	}

	_, err = r.db.Exec("UPDATE TOURNAMENTS SET WINNER_ID=? WHERE TOURNAMENT_ID=?", winnerID, tournamentID)
	if err != nil {
		return fmt.Errorf("set tournament winner error: %v", err)
	}

	return nil
}

func (r Repository) GetTournamentUsers(ctx context.Context, id int64) ([]model.User, error) {
	_, err := r.db.Query("SELECT * FROM TOURNAMENTS WHERE TOURNAMENT_ID=?", id)
	if err == sql.ErrNoRows {
		return nil, TournamentNotFoundError
	}

	var count int
	row := r.db.QueryRow("SELECT COUNT(*) FROM USERS AS t1 JOIN MTM_USER_TOURNAMENT AS t2 ON t1.USER_ID = t2.USER_ID WHERE t2.TOURNAMENT_ID = ?", id)
	err = row.Scan(&count)
	if err != nil {
		return nil, fmt.Errorf("error while getting count of users: %v", err)
	}
	if count == 0 {
		return nil, errors.New("users not found error")
	}

	rows, err := r.db.Query("SELECT t1.* FROM USERS AS t1 JOIN MTM_USER_TOURNAMENT AS t2 ON t1.USER_ID = t2.USER_ID WHERE t2.TOURNAMENT_ID = ?", id)
	if err != nil {
		return nil, fmt.Errorf("error while getting users: %v", err)
	}
	defer rows.Close()

	users := make([]model.User, 0, count)
	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&user.ID, &user.Username, &user.Balance)
		if err != nil {
			return nil, fmt.Errorf("error while scanning users: %v", err)
		}
		users = append(users, user)
	}

	return users, nil
}
