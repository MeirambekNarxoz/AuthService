package repository

import (
	"authService/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) LoadUserRole(user *models.User) error {
	return r.db.Preload("Role").First(user, user.ID).Error
}

func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Preload("Role").Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
