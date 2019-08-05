package repository

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/777777miSSU7777777/gaming-website/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return UserRepository{db}
}

func (r UserRepository) New(ctx context.Context, username string, balance int64) (int64, error) {
	result, err := r.db.Exec("INSERT INTO USERS (USERNAME, BALANCE) VALUES(?, ?)", username, balance)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r UserRepository) GetByID(ctx context.Context, id int64) (model.User, error) {
	row := r.db.QueryRow("SELECT * FROM USERS WHERE USER_ID=?", id)
	user := model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Balance)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r UserRepository) DeleteByID(ctx context.Context, id int64) error {
	_, err := r.db.Exec("DELETE FROM USERS WHERE USER_ID=?", id)
	if err != nil {
		return err
	}

	return nil
}

func (r UserRepository) TakeBalanceByID(ctx context.Context, id int64, points int64) error {
	row := r.db.QueryRow("SELECT * FROM USERS WHERE USER_ID=?", id)
	user := model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Balance)
	if err != nil {
		return err
	}

	_, err = r.db.Exec("UPDATE USERS SET BALANCE=? WHERE USER_ID=?", user.Balance-points, id)
	if err != nil {
		return err
	}

	return nil
}

func (r UserRepository) AddBalanceByID(ctx context.Context, id int64, points int64) error {
	row := r.db.QueryRow("SELECT * FROM USERS WHERE USER_ID=?", id)
	user := model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Balance)
	if err != nil {
		return err
	}
	_, err = r.db.Exec("UPDATE USERS SET BALANCE=? WHERE USER_ID=?", user.Balance+points, id)
	if err != nil {
		return err
	}
	return nil
}
