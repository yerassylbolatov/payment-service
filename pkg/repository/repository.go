package repository

import "database/sql"

type Authorization interface {
}

type Transactions interface {
}

type Repository struct {
	Authorization
	Transactions
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{}
}
