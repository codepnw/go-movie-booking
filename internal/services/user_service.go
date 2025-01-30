package services

import (
	"context"
	"errors"
	"log"

	"github.com/codepnw/go-movie-booking/internal/models"
	"github.com/codepnw/go-movie-booking/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	Register(ctx context.Context, u *models.UserRegisterReq) error
	Login(ctx context.Context, req *models.UserLoginReq) (*models.User, error)
	GetByID(ctx context.Context, id int64) (*models.User, error)
}

type userService struct {
	repo repositories.IUserRepository
}

func NewUserService(repo repositories.IUserRepository) IUserService {
	return &userService{repo: repo}
}

func (s *userService) Register(ctx context.Context, u *models.UserRegisterReq) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		Username: u.Username,
		Email:    u.Email,
		Password: string(hashedPassword),
	}

	return s.repo.Create(ctx, user)
}

func (s *userService) Login(ctx context.Context, req *models.UserLoginReq) (*models.User, error) {
	user, err := s.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		log.Println(err)
		return nil, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		log.Println(err)
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}

func (s *userService) GetByID(ctx context.Context, id int64) (*models.User, error) {
	return s.repo.GetByID(ctx, id)
}
