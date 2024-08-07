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
	FindAdminByUserID(ID string, showUser bool) (models.Admin, error)
	FindStaffByUserID(ID string, showUser bool) (models.Staff, error)
	FindInstructorByUserID(ID string, showUser bool) (models.Instructor, error)
	FindCustomerByUserID(ID string, showUser bool) (models.Customer, error)
	Update(user models.User) (models.User, error)
	UpdateAdmin(admin models.Admin) (models.Admin, error)
	UpdateStaff(staff models.Staff) (models.Staff, error)
	UpdateInstructor(instructor models.Instructor) (models.Instructor, error)
	UpdateCustomer(customer models.Customer) (models.Customer, error)
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

func (r *userRepository) FindAdminByUserID(ID string, showUser bool) (models.Admin, error) {
	var user models.Admin

	if(showUser) {
		err := r.db.DB.Preload("User").Where("user_id = ?", ID).Find(&user).Error
		if err != nil {
			return user, err
		}
	} else {
		err := r.db.DB.Where("user_id = ?", ID).Find(&user).Error
		if err != nil {
			return user, err
		}
	}

	return user, nil
}

func (r *userRepository) FindStaffByUserID(ID string, showUser bool) (models.Staff, error) {
	var user models.Staff

	if(showUser) {
		err := r.db.DB.Preload("User").Where("user_id = ?", ID).Find(&user).Error
		if err != nil {
			return user, err
		}
	} else {
		err := r.db.DB.Where("user_id = ?", ID).Find(&user).Error
		if err != nil {
			return user, err
		}
	}

	return user, nil
}

func (r *userRepository) FindInstructorByUserID(ID string, showUser bool) (models.Instructor, error) {
	var user models.Instructor

	if(showUser) {
		err := r.db.DB.Preload("User").Where("user_id = ?", ID).Find(&user).Error
		if err != nil {
			return user, err
		}
	} else {
		err := r.db.DB.Where("user_id = ?", ID).Find(&user).Error
		if err != nil {
			return user, err
		}
	}

	return user, nil
}

func (r *userRepository) FindCustomerByUserID(ID string, showUser bool) (models.Customer, error) {
	var user models.Customer

	if(showUser) {
		err := r.db.DB.Preload("User").Where("user_id = ?", ID).Find(&user).Error
		if err != nil {
			return user, err
		}
	} else {
		err := r.db.DB.Where("user_id = ?", ID).Find(&user).Error
		if err != nil {
			return user, err
		}
	}

	return user, nil
}

func (r *userRepository) Update(user models.User) (models.User, error) {
	// err := r.db.DB.Order("created_at desc").First(&user).Unscoped().Delete(&user).Error
	// if err != nil {
	// 	return user, err
	// }

	err := r.db.DB.Save(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) UpdateAdmin(admin models.Admin) (models.Admin, error) {
	// var user models.User

	// err := r.db.DB.Order("created_at desc").First(&user).Unscoped().Delete(&user).Error
	// if err != nil {
	// 	return admin, err
	// }

	err := r.db.DB.Save(&admin).Error

	if err != nil {
		return admin, err
	}

	return admin, nil
}

func (r *userRepository) UpdateStaff(staff models.Staff) (models.Staff, error) {
	err := r.db.DB.Save(&staff).Error

	if err != nil {
		return staff, err
	}

	return staff, nil
}

func (r *userRepository) UpdateInstructor(instructor models.Instructor) (models.Instructor, error) {
	err := r.db.DB.Save(&instructor).Error

	if err != nil {
		return instructor, err
	}

	return instructor, nil
}

func (r *userRepository) UpdateCustomer(customer models.Customer) (models.Customer, error) {
	err := r.db.DB.Save(&customer).Error

	if err != nil {
		return customer, err
	}

	return customer, nil
}
