// This file contains the repository implementation layer.
package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	Db *pgxpool.Pool
}

type NewRepositoryOptions struct {
	Dsn string
}

func NewRepository(opts NewRepositoryOptions) *Repository {
	db, err := pgxpool.New(context.Background(), opts.Dsn)
	if err != nil {
		log.Fatal("Error while creating connection to the database!!")
	}
	return &Repository{
		Db: db,
	}
}
