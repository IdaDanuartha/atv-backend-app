package repositories

import (
	"errors"

	"github.com/IdaDanuartha/atv-backend-app/app/config"
	// "github.com/IdaDanuartha/atv-backend-app/app/enums"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(customer models.Customer) (models.Customer, error)
	FindByEmail(email string) (models.User, error)
	FindByUsername(username string) (models.User, error)
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

	// if user.Role == enums.Admin {
	// 	var admin models.Admin

	// 	err := r.db.DB.Where("user_id =?", user.ID).Find(&admin).Error
	// 	if err != nil {
	// 		return user, err
	// 	}

	// 	admin.UserID = user.ID

	// 	err = r.db.DB.Save(&admin).Error
	// 	if err != nil {
	// 		return user, err
	// 	}
	// } else if user.Role == enums.Staff {
	// 	var staff models.Staff

	// 	err := r.db.DB.Where("user_id =?", user.ID).Find(&staff).Error
	// 	if err != nil {
	// 		return user, err
	// 	}

	// 	staff.UserID = user.ID

	// 	err = r.db.DB.Save(&staff).Error
	// 	if err != nil {
	// 		return user, err
	// 	}
	// } else if user.Role == enums.Instructor {
	// 	var instructor models.Instructor

	// 	err := r.db.DB.Where("user_id =?", user.ID).Find(&instructor).Error
	// 	if err != nil {
	// 		return user, err
	// 	}

	// 	instructor.UserID = user.ID

	// 	err = r.db.DB.Save(&instructor).Error
	// 	if err != nil {
	// 		return user, err
	// 	}
	// } else if user.Role == enums.Customer {
	// 	var customer models.Customer

	// 	err := r.db.DB.Where("user_id =?", user.ID).Find(&customer).Error
	// 	if err != nil {
	// 		return user, err
	// 	}

	// 	customer.UserID = user.ID

	// 	err = r.db.DB.Save(&customer).Error
	// 	if err != nil {
	// 		return user, err
	// 	}
	// }

	return user, nil
}
