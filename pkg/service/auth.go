package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/CheesyBoy03/mini-social-app/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	tokenTKL = 12 * time.Hour
)

type AuthService struct {
	repo       repository.Auth
	salt       string
	signingKey []byte
}

func NewAuthService(repo repository.Auth, salt string, signingKey []byte) *AuthService {
	return &AuthService{
		repo:       repo,
		salt:       salt,
		signingKey: signingKey,
	}
}

func (s *AuthService) SignIn(email, password string) (string, error) {
	hashedPassword := s.getPasswordHash(password)

	err := s.repo.Authorize(email, hashedPassword)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(tokenTKL).Unix(),
		IssuedAt:  time.Now().Unix(),
	})

	return token.SignedString(s.signingKey)
}

func (s *AuthService) SignUp(email, password, firstname, lastname string) error {
	hashedPassword := s.getPasswordHash(password)
	err := s.repo.Register(email, hashedPassword, firstname, lastname)
	return err
}

func (s *AuthService) ParseToken(token string) error {
	t, err := jwt.Parse(token, func(token *jwt.Token) (i any, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return s.signingKey, nil
	})
	if err != nil {
		return err
	}

	_, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return fmt.Errorf("error get user claims from token")
	}

	return nil
}

func (s *AuthService) getPasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(s.salt))

	return fmt.Sprintf("%x", sha1.Sum([]byte(password)))
}
