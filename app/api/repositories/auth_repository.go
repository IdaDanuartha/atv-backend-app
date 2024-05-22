package repositories

import (
	"github.com/IdaDanuartha/atv-backend-app/app/config"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

//AuthRepository -> AuthRepository
type AuthRepository struct {
    db config.Database
}

// NewAuthRepository : fetching database
func NewAuthRepository(db config.Database) AuthRepository {
    return AuthRepository{
        db: db,
    }
}

func (r *AuthRepository) Save(user models.Customer) (models.Customer, error) {
	err := r.db.DB.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *AuthRepository) FindByEmail(email string) (models.User, error) {
	var user models.User

	err := r.db.DB.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *AuthRepository) FindByID(ID int) (models.User, error) {
	var user models.User

	err := r.db.DB.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}