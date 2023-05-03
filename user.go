package mini_social_app

import "time"

type User struct {
	ID           int64     `db:"id"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	FirstName    string    `db:"first_name"`
	SecondName   string    `db:"second_name"`
	RegisteredAt time.Time `db:"registered_at"`
}
