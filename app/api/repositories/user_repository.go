package repositories

import (
	"github.com/IdaDanuartha/atv-backend-app/app/config"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

type UserRepository interface {
	Save(user models.User) (models.User, error)
	FindByEmail(email string) (models.User, error)
	FindByID(ID string) (models.User, error)
	Update(user models.User) (models.User, error)
}

type userRepository struct {
	db config.Database
}

func NewUserRepository(db config.Database) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Save(user models.User) (models.User, error) {
	err := r.db.DB.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindByEmail(email string) (models.User, error) {
	var user models.User

	err := r.db.DB.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindByID(ID string) (models.User, error) {
	var user models.User

	err := r.db.DB.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) Update(user models.User) (models.User, error) {
	err := r.db.DB.Save(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
