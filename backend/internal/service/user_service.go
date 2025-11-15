package service

import (
	"backend/internal/app"
	"backend/internal/repository"
	"context"

	"gorm.io/gorm"
)

type UserService struct {
	db             *gorm.DB
	userRepository *repository.UserRepository
}

func NewUserService(db *gorm.DB, userRepository *repository.UserRepository) *UserService {
	return &UserService{
		db:             db,
		userRepository: userRepository,
	}
}

func (s *UserService) GetAll(ctx context.Context) ([]app.User, error) {
	tx := s.db.WithContext(ctx)
	users, err := s.userRepository.FindAll(tx)
	if err != nil {
		return nil, err
	}
	return users, nil
}
