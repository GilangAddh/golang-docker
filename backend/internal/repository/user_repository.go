package repository

import (
	"backend/internal/app"
	"fmt"

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
	if err := db.Debug().Where("deleted_at IS NULL").Find(&users).Error; err != nil {
		return nil, err
	}
	fmt.Println(len(users))
	return users, nil
}
