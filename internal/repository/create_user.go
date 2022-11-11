package repository

import (
	"github.com/jmoiron/sqlx"
	"rest_api_test"
)

type Postgres struct {
	db *sqlx.DB
}

func NewPostgres(db *sqlx.DB) *Postgres {
	return &Postgres{db: db}
}

func (r *Postgres) CreateUser(user rest_api_test.User) error {
	tx, err := r.db.Begin()
	sqlStr := `insert into "usr" values ($1,$2)`
	_, err = tx.Exec(sqlStr, user.Id, user.Balance)
	if err != nil {
		tx.Rollback()
		return err
	}
	return err
}
