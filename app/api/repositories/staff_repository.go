package repositories

import (
	"errors"

	"github.com/IdaDanuartha/atv-backend-app/app/config"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"gorm.io/gorm"
)

type StaffRepository interface {
	FindAll(staff models.Staff, search string, currentPage int, pageSize int) ([]models.Staff, int64, int, error)
	Find(ID string, showRelations bool) (models.Staff, error)
	Save(staff models.Staff) (models.Staff, error)
	Update(staff models.Staff) (models.Staff, error)
	Delete(staff models.Staff) (models.Staff, error)
}

type staffRepository struct {
	db config.Database
}

// NewStaffRepository : fetching database
func NewStaffRepository(db config.Database) staffRepository {
	return staffRepository{db}
}

// FindAll -> Method for fetching all Staff from database
func (r staffRepository) FindAll(staff models.Staff, search string, currentPage int, pageSize int) ([]models.Staff, int64, int, error) {
	var staffs []models.Staff
	var totalRows int64 = 0

	queryBuilder := r.db.DB.Order("created_at desc").Model(&models.Staff{})

	// Search parameter
	if search != "" {
		querySearch := "%" + search + "%"
		queryBuilder = queryBuilder.Joins("JOIN users ON users.id = staff.user_id").Where(
			r.db.DB.Where("staff.name LIKE ? ", querySearch).
				Or("staff.employee_code LIKE ? ", querySearch).
				Or("users.username LIKE ? ", querySearch).
				Or("users.email LIKE ? ", querySearch))
	}

	if pageSize > 0 {
		// count the total number of rows
		err := queryBuilder.
			Where(staff).
			Count(&totalRows).Error
		if err != nil {
			return nil, 0, 0, err
		}
			
		// Apply offset and limit to fetch paginated results
		err = queryBuilder.
			Preload("User").
			Where(staff).
			Offset((currentPage - 1) * pageSize).
			Limit(pageSize).
			Find(&staffs).Error
		return staffs, totalRows, currentPage, err
	} else {
		err := queryBuilder.
			Preload("User").
			Where(staff).
			Find(&staffs).
			Count(&totalRows).Error
		return staffs, 0, 0, err
	}
}

// Find -> Method for fetching Staff by id
func (r staffRepository) Find(ID string, showRelations bool) (models.Staff, error) {
	var staffs models.Staff
	
	if(showRelations) {
		err := r.db.DB.
			Preload("User").
			Debug().
			Model(&models.Staff{}).
			Where("id = ?", ID).
			Find(&staffs).Error
		return staffs, err
	} else {
		err := r.db.DB.
			Debug().
			Model(&models.Staff{}).
			Where("id = ?", ID).
			Find(&staffs).Error
		return staffs, err
	}
}

// Save -> Method for saving Staff to database
func (r staffRepository) Save(staff models.Staff) (models.Staff, error) {
	userRepository := NewUserRepository(r.db)

	// 1. Validate email existence (using FindByEmail)
	existingUser, err := userRepository.FindByEmail(staff.User.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return staff, err // Return error if not a "not found" error
	}

	if existingUser.ID != "" {
		return staff, errors.New("email already exists") // Return error if email exists
	}

	// 2. Validate username existence (using FindByUsername)
	existingUser, err = userRepository.FindByUsername(staff.User.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return staff, err // Return error if not a "not found" error
	}

	if existingUser.ID != "" {
		return staff, errors.New("username already exists") // Return error if username exists
	}

	err = r.db.DB.Preload("User").Create(&staff).Error
	if err != nil {
		return staff, err
	}

	return staff, nil
}

// Update -> Method for updating Staff
func (r *staffRepository) Update(staff models.Staff) (models.Staff, error) {
	err := r.db.DB.Save(&staff).Error
	if err != nil {
		return staff, err
	}

	return staff, nil
}

// Delete -> Method for deleting Staff
func (r staffRepository) Delete(staff models.Staff) (models.Staff, error) {
	err := r.db.DB.Delete(&staff.User).Error
	if err != nil {
		return staff, err
	}

	err = r.db.DB.Preload("User").Delete(&staff).Error
	if err != nil {
		return staff, err
	}

	return staff, nil
}
