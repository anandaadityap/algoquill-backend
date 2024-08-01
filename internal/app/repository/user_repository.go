package repository

import (
	"github.com/anandaadityap/algoquill-backend/internal/app/models/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user entity.User) (entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (ur *userRepository) Save(user entity.User) (entity.User, error) {
	err := ur.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
