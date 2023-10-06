package repository

import "database/sql"

type Repository struct {
	Db *sql.DB
}

func New(dsn string) *Repository {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	return &Repository{
		Db: db,
	}
}
