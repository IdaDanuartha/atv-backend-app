package repositories

import (
	"errors"

	"github.com/IdaDanuartha/atv-backend-app/app/config"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user models.Customer) (models.Customer, error)
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

func (r *userRepository) Save(customer models.Customer) (models.Customer, error) {
	// 1. Validate email existence (using FindByEmail)
	existingUser, err := r.FindByEmail(customer.User.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return customer, err // Return error if not a "not found" error
	}

	if existingUser.ID != "" {
		return customer, errors.New("email already exists") // Return error if email exists
	}

	// 2. Validate username existence (using FindByUsername)
	existingUser, err = r.FindByUsername(customer.User.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return customer, err // Return error if not a "not found" error
	}

	if existingUser.ID != "" {
		return customer, errors.New("username already exists") // Return error if username exists
	}

	// 3. Create and save the User (if email doesn't exist)
	err = r.db.DB.Create(&customer.User).Error
	if err != nil {
		return customer, err
	}

	// 4. Set the UserID in the Customer object
	customer.UserID = customer.User.ID

	// 5. Save the Customer
	err = r.db.DB.Create(&customer).Error
	if err != nil {
		return customer, err
	}

	return customer, nil
}

func (r *userRepository) FindByEmail(email string) (models.User, error) {
	var user models.User

	err := r.db.DB.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindByUsername(username string) (models.User, error) {
	var user models.User

	err := r.db.DB.Where("username = ?", username).Find(&user).Error
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
