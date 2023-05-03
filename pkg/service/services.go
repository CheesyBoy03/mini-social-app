package service

import "github.com/CheesyBoy03/mini-social-app/pkg/repository"

type SignUpInput struct {
	FirstName  string
	SecondName string
	Email      string
	Password   string
}

type Auth interface {
	SignIn(email, password string) (string, error)
	SignUp(email, password, firstname, lastname string) error
	ParseToken(token string) error
}

type Service struct {
	Auth
}

type Dependencies struct {
	Repos      *repository.Repository
	HashSalt   string
	SigningKey []byte
}

func NewServices(deps Dependencies) *Service {
	return &Service{
		Auth: NewAuthService(deps.Repos.Auth, deps.HashSalt, deps.SigningKey),
	}
}
