package repository

import (
	"context"
	"database/sql"
	"errors"

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
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (r Repository) GetUserByID(ctx context.Context, id int64) (model.User, error) {
	row := r.db.QueryRow("SELECT * FROM USERS WHERE USER_ID=?", id)
	user := model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Balance)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (r Repository) DeleteUserByID(ctx context.Context, id int64) error {
	res, err := r.db.Exec("DELETE FROM USERS WHERE USER_ID=?", id)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("USER NOT FOUND ERROR")
	}

	return nil
}

func (r Repository) TakeUserBalanceByID(ctx context.Context, id int64, points int64) error {
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

func (r Repository) AddUserBalanceByID(ctx context.Context, id int64, points int64) error {
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