package userrepository

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/777777miSSU7777777/gaming-website/pkg/entity"
)

type UserRepository interface {
	New(context.Context, string, int64) (int64, error)
	GetByID(context.Context, int64) (entity.User, error)
	DeleteByID(context.Context, int64) error
	UpdateByID(context.Context, int64, string, int64) error
}
type repository struct {
	db *sql.DB
}

func New(db *sql.DB) UserRepository {
	return &repository{db}
}

func (r repository) New(ctx context.Context, username string, balance int64) (int64, error) {
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

func (r repository) GetByID(ctx context.Context, id int64) (entity.User, error) {
	row := r.db.QueryRow("SELECT * FROM USERS WHERE USER_ID=?", id)
	user := entity.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Balance)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r repository) DeleteByID(ctx context.Context, id int64) error {
	_, err := r.db.Exec("DELETE FROM USERS WHERE USER_ID=?", id)
	if err != nil {
		return err
	}
	return nil
}

func (r repository) UpdateByID(ctx context.Context, id int64, username string, balance int64) error {
	_, err := r.db.Exec("UPDATE USERS SET USERNAME=?, BALANCE=? WHERE USER_ID=?", username, balance, id)
	if err != nil {
		return err
	}
	return nil
}
