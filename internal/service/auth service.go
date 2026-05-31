package service

import (
	"errors"

	"github.com/georgigeorgiev/cloudops/auth-service/internal/model"
	"github.com/georgigeorgiev/cloudops/auth-service/internal/repository"

	jwtpkg "github.com/georgigeorgiev/cloudops/auth-service/pkg/jwt"
	"github.com/georgigeorgiev/cloudops/auth-service/pkg/password"
)

type AuthService struct {
	Repo      repository.UserRepository
	JWTSecret string
}

func (s *AuthService) Register(
	email string,
	rawPassword string,
) error {

	existing, _ := s.Repo.FindByEmail(email)

	if existing != nil {
		return errors.New("user already exists")
	}

	hashed, err := password.Hash(rawPassword)

	if err != nil {
		return err
	}

	user := &model.User{
		Email:    email,
		Password: hashed,
		Role:     "USER",
	}

	return s.Repo.Create(user)
}

func (s *AuthService) Login(
	email string,
	rawPassword string,
) (string, error) {

	user, err := s.Repo.FindByEmail(email)

	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !password.Compare(
		user.Password,
		rawPassword,
	) {
		return "", errors.New("invalid credentials")
	}

	return jwtpkg.Generate(
		user.ID,
		user.Email,
		s.JWTSecret,
	)
}
