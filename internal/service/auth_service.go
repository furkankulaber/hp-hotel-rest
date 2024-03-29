package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"hp-hotel-rest/internal/model"
	"time"

	"hp-hotel-rest/internal/repository"
	"hp-hotel-rest/pkg/config"
)

type AuthService interface {
	Login(request model.LoginRequest) (string, error)
	Register(request model.LoginRequest) error
}

type authService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) Login(request model.LoginRequest) (string, error) {
	user, err := s.repo.FindByCredentials(request.Email)
	if err != nil {
		return "", errors.New("user not found or invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	day := time.Hour * 24
	claims := jwt.MapClaims{
		"ID":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(day * 1).Unix(),
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return "", errors.New("error generating token")
	}

	return t, nil
}

func (s *authService) Register(user model.LoginRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	userModel := &model.User{
		Email:    user.Email,
		Password: string(hashedPassword), // Convert byte slice to string
	}

	err = s.repo.CreateUser(userModel)
	if err != nil {
		return err
	}

	return nil
}
