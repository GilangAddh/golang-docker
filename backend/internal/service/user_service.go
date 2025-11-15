package service

import (
	"backend/internal/app"
	"backend/internal/repository"
	"context"
	"fmt"

	"github.com/gofiber/fiber"
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

func (s *UserService) GetByID(ctx context.Context, dataID uint) (app.User, error) {
	tx := s.db.WithContext(ctx)
	user, err := s.userRepository.FindDetail(tx, dataID)
	if err != nil {
		return app.User{}, err
	}
	return user, nil
}

func (s *UserService) Create(ctx context.Context, input app.CreateUserDTO) (any, error) {
	tx := s.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	user := app.User{
		Name:     input.Name,
		Position: input.Position,
		Salary:   input.Salary,
	}

	if err := tx.Create(&user).Error; err != nil {
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}
	return fiber.Map{"user_id": user.ID}, nil
}

func (s *UserService) Update(ctx context.Context, dataID uint, input app.UpdateUserDTO) error {
	tx := s.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	user, err := s.userRepository.FindDetail(tx, dataID)
	if err != nil {
		return err
	}

	user.Name = input.Name
	user.Position = input.Position
	user.Salary = input.Salary

	if err := tx.Save(&user).Error; err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (s *UserService) Delete(ctx context.Context, dataID uint) error {
	tx := s.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	user, err := s.userRepository.FindDetail(tx, dataID)
	if err != nil {
		return err
	}

	if err := tx.Delete(&user).Error; err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}
