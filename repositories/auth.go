package repositories

import (
	"waysbuck/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
}

func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Register(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}
