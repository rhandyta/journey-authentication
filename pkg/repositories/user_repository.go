package repositories

import (
	"journey-user/pkg/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserById(userId uint) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) CreateUser(user *models.User) error {
	return ur.db.Create(user).Error
}

func (ur *userRepository) GetUserById(userId uint) (*models.User, error) {
	user := &models.User{}
	if err := ur.db.First(user, userId).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	if err := ur.db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
