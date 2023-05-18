package repository

import (
	"database/sql"
	"errors"
	"fmt"

	miniSocialApp "github.com/CheesyBoy03/mini-social-app"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) Authorize(email, passwordHash string) error {
	var user *miniSocialApp.User
	query := "SELECT * FROM users WHERE email=$1 AND password_hash=$2"
	err := r.db.QueryRow(query, email, passwordHash).Scan(&user)

	if user == nil {
		return errors.New("Wrong email or password")
	}

	fmt.Println(user)

	return err
}

func (r *AuthRepository) Register(email, password, firstname, lastname string) error {
	var isExistUser bool
	q := "SELECT EXISTS(SELECT 1 FROM users WHERE email=$1 AND password_hash=$2)"
	err := r.db.QueryRow(q, email, password).Scan(&isExistUser)
	if err != nil {
		return err
	}

	if isExistUser {
		return errors.New("User is already registered")
	}

	query := "INSERT INTO users (email, password_hash, first_name, last_name) VALUES ($1, $2, $3, $4)"

	_, err = r.db.Exec(query, email, password, firstname, lastname)
	if err != nil {
		fmt.Printf("Error on creating user: %s\n", err.Error())
	}

	return err
}
