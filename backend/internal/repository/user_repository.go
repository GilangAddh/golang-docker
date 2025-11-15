package repository

import (
	"backend/internal/app"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) FindAll(db *gorm.DB) ([]app.User, error) {
	var users []app.User
	if err := db.Where("deleted_at IS NULL").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) FindDetail(db *gorm.DB, dataID uint) (app.User, error) {
	var user app.User
	if err := db.Where("deleted_at IS NULL").First(&user, dataID).Error; err != nil {
		return app.User{}, err
	}

	return user, nil
}
