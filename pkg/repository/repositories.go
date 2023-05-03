package repository

import "database/sql"

type Auth interface {
	Authorize(email, password string) error
	Register(email, password, firstname, lastname string) error
}

type Repository struct {
	Auth
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Auth: NewAuthRepository(db),
	}
}
