package stores

import (
	"cashflow/pkg/interfaces"
	"cashflow/pkg/models"

	"gorm.io/gorm"
)

type userStore struct {
	db *gorm.DB
}

func (us *userStore) GetByEmail(email string) (*models.User, error) {
	var user models.User
	result := us.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

func (us *userStore) Create(user *models.User) (*models.User, error) {
	result := us.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func CreateUserStore(db *gorm.DB) interfaces.UserStore {
	return &userStore{db: db}
}
