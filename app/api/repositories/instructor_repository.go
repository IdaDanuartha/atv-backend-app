package repositories

import (
	"errors"

	"github.com/IdaDanuartha/atv-backend-app/app/config"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"gorm.io/gorm"
)

type InstructorRepository interface {
	FindAll(instructor models.Instructor, search string, currentPage int, pageSize int) ([]models.Instructor, int64, int, error)
	Find(ID string, showRelations bool) (models.Instructor, error)
	FindByUserID(userID string) (models.Instructor, error)
	Save(instructor models.Instructor) (models.Instructor, error)
	Update(instructor models.Instructor) (models.Instructor, error)
	Delete(instructor models.Instructor) (models.Instructor, error)
}

type instructorRepository struct {
	db config.Database
}

// NewInstructorRepository : fetching database
func NewInstructorRepository(db config.Database) instructorRepository {
	return instructorRepository{db}
}

// FindAll -> Method for fetching all Instructor from database
func (r instructorRepository) FindAll(instructor models.Instructor, search string, currentPage int, pageSize int) ([]models.Instructor, int64, int, error) {
	var instructors []models.Instructor
	var totalRows int64 = 0

	queryBuilder := r.db.DB.Order("created_at desc").Model(&models.Instructor{})

	// Search parameter
	if search != "" {
		querySearch := "%" + search + "%"
		queryBuilder = queryBuilder.Joins("JOIN users ON users.id = instructors.user_id").Where(
			r.db.DB.Where("instructors.name LIKE ? ", querySearch).
				Or("instructors.employee_code LIKE ? ", querySearch).
				Or("users.username LIKE ? ", querySearch).
				Or("users.email LIKE ? ", querySearch))
	}

	if pageSize > 0 {
		// count the total number of rows
		err := queryBuilder.
			Where(instructor).
			Count(&totalRows).Error
		if err != nil {
			return nil, 0, 0, err
		}

		// Apply offset and limit to fetch paginated results
		err = queryBuilder.
			Preload("User").
			Preload("Services.EntertainmentService.EntertainmentCategory").
			Where(instructor).
			Offset((currentPage - 1) * pageSize).
			Limit(pageSize).
			Find(&instructors).Error
		return instructors, totalRows, currentPage, err
	} else {
		err := queryBuilder.
			Preload("User").
			Preload("Services.EntertainmentService.EntertainmentCategory").
			Where(instructor).
			Find(&instructors).
			Count(&totalRows).Error
		return instructors, 0, 0, err
	}
}

// Find -> Method for fetching Instructor by id
func (r instructorRepository) Find(ID string, showRelations bool) (models.Instructor, error) {
	var instructors models.Instructor
	
	if(showRelations) {
		err := r.db.DB.
		Preload("User").
		Preload("Services.EntertainmentService.EntertainmentCategory").
		Debug().
		Model(&models.Instructor{}).
		Where("id = ?", ID).
		Find(&instructors).Error

		return instructors, err
	} else {
		err := r.db.DB.
		Debug().
		Model(&models.Instructor{}).
		Where("id = ?", ID).
		Find(&instructors).Error

		return instructors, err
	}
}

// FindByUserID -> Method for fetching Instructor by user id
func (r instructorRepository) FindByUserID(userID string) (models.Instructor, error) {
	var instructors models.Instructor
	err := r.db.DB.
		Preload("User").
		Preload("Services.EntertainmentService.EntertainmentCategory").
		Debug().
		Model(&models.Instructor{}).
		Where("user_id = ?", userID).
		Find(&instructors).Error
	return instructors, err
}

// Save -> Method for saving Instructor to database
func (r instructorRepository) Save(instructor models.Instructor) (models.Instructor, error) {
	userRepository := NewUserRepository(r.db)

	// 1. Validate email existence (using FindByEmail)
	existingUser, err := userRepository.FindByEmail(instructor.User.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return instructor, err // Return error if not a "not found" error
	}

	if existingUser.ID != "" {
		return instructor, errors.New("email already exists") // Return error if email exists
	}

	// 2. Validate username existence (using FindByUsername)
	existingUser, err = userRepository.FindByUsername(instructor.User.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return instructor, err // Return error if not a "not found" error
	}

	if existingUser.ID != "" {
		return instructor, errors.New("username already exists") // Return error if username exists
	}

	err = r.db.DB.Create(&instructor).Error
	if err != nil {
		return instructor, err
	}

	return instructor, nil
}

// Update -> Method for updating Instructor
func (r *instructorRepository) Update(instructor models.Instructor) (models.Instructor, error) {
	err := r.db.DB.Save(&instructor).Error
	if err != nil {
		return instructor, err
	}

	return instructor, nil
}

// Delete -> Method for deleting Instructor
func (r instructorRepository) Delete(instructor models.Instructor) (models.Instructor, error) {
	err := r.db.DB.Delete(&instructor.User).Error
	if err != nil {
		return instructor, err
	}

	err = r.db.DB.Delete(&instructor).Error
	if err != nil {
		return instructor, err
	}

	return instructor, nil
}
