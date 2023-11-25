package repository

import "github.com/jackc/pgx/v5/pgxpool"

type Repo struct {
	db *pgxpool.Pool
}

func NewRepo(conn *pgxpool.Pool) (*Repo, error) {
	return &Repo{
		db: conn,
	}, nil
}
