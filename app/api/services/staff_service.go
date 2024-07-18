package services

import (
	"github.com/IdaDanuartha/atv-backend-app/app/api/inputs"
	"github.com/IdaDanuartha/atv-backend-app/app/api/repositories"
	// "github.com/IdaDanuartha/atv-backend-app/app/enums"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"golang.org/x/crypto/bcrypt"
)

type StaffService interface {
	FindAll(model models.Staff, search string, currentPage int, pageSize int) ([]models.Staff, int64, int, error)
	Find(input inputs.GetStaffDetailInput) (models.Staff, error)
	Save(input inputs.StaffInput) (models.Staff, error)
	Update(inputID inputs.GetStaffDetailInput, input inputs.EditStaffInput) (models.Staff, error)
	Delete(inputID inputs.GetStaffDetailInput) (models.Staff, error)
}

// StaffService StaffService struct
type staffService struct {
	repository repositories.StaffRepository
	userRepository repositories.UserRepository
}

// NewStaffService : returns the StaffService struct instance
func NewStaffService(repository repositories.StaffRepository, userRepository repositories.UserRepository) staffService {
	return staffService{repository, userRepository}
}

// FindAll -> calls Staff repo find all method
func (s staffService) FindAll(model models.Staff, search string, currentPage int, pageSize int) ([]models.Staff, int64, int, error) {
	staffs, total, currentPage, err := s.repository.FindAll(model, search, currentPage, pageSize)
	if err != nil {
		return staffs, total, currentPage, err
	}

	return staffs, total, currentPage, nil
}

// Find -> calls Staff repo find method
func (s staffService) Find(input inputs.GetStaffDetailInput) (models.Staff, error) {
	staff, err := s.repository.Find(input.ID, true)

	if err != nil {
		return staff, err
	}

	return staff, nil
}

// Save -> calls Staff repository save method
func (s staffService) Save(input inputs.StaffInput) (models.Staff, error) {
	staff := models.Staff{}
	staff.Name = input.Name
	staff.EmployeeCode = input.EmployeeCode
	staff.User.Username = input.Username
	staff.User.Email = input.Email
	staff.User.Password = input.Password

	// staff.User.Role = enums.Role(enums.Staff)
	staff.User.Role = "staff"

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return staff, err
	}

	staff.User.Password = string(passwordHash)

	newStaff, err := s.repository.Save(staff)
	if err != nil {
		return newStaff, err
	}

	return newStaff, nil
}

// Update -> calls Staff repo update method
func (s staffService) Update(inputID inputs.GetStaffDetailInput, input inputs.EditStaffInput) (models.Staff, error) {
	staff, err := s.repository.Find(inputID.ID, false)
	if err != nil {
		return staff, err
	}

	staff.Name = input.Name
	staff.EmployeeCode = input.EmployeeCode
	
	user, _ := s.userRepository.FindByID(staff.UserID)
	password := user.Password

	user.Username = input.Username
	user.Email = input.Email
	user.Password = password
	user.Role = "staff"

	if input.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
		if err != nil {
			return staff, err
		}

		user.Password = string(password)
	}

	_, err = s.userRepository.Update(user)
	if err != nil {
		return staff, err
	}

	updatedStaff, err := s.repository.Update(staff)
	if err != nil {
		return updatedStaff, err
	}

	return updatedStaff, nil
}

// Delete -> calls Staff repo delete method
func (s staffService) Delete(inputID inputs.GetStaffDetailInput) (models.Staff, error) {
	staff, err := s.repository.Find(inputID.ID, true)
	if err != nil {
	return staff, err
	}

	deletedStaff, err := s.repository.Delete(staff)
	if err != nil {
		return deletedStaff, err
	}

	return deletedStaff, nil
}
