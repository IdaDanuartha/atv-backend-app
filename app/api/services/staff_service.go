package services

import (
	"github.com/IdaDanuartha/atv-backend-app/app/api/inputs"
	"github.com/IdaDanuartha/atv-backend-app/app/api/repositories"
	// "github.com/IdaDanuartha/atv-backend-app/app/enums"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"golang.org/x/crypto/bcrypt"
)

type StaffService interface {
	FindAll(model models.Staff, search string) ([]models.Staff, int64, error)
	Find(input inputs.GetStaffDetailInput) (models.Staff, error)
	Save(input inputs.StaffInput) (models.Staff, error)
	Update(inputID inputs.GetStaffDetailInput, input inputs.EditStaffInput) (models.Staff, error)
	Delete(inputID inputs.GetStaffDetailInput) (models.Staff, error)
}

// StaffService StaffService struct
type staffService struct {
	repository repositories.StaffRepository
}

// NewStaffService : returns the StaffService struct instance
func NewStaffService(repository repositories.StaffRepository) staffService {
	return staffService{repository}
}

// FindAll -> calls Staff repo find all method
func (s staffService) FindAll(model models.Staff, search string) ([]models.Staff, int64, error) {
	staffs, total, err := s.repository.FindAll(model, search)
	if err != nil {
		return staffs, total, err
	}

	return staffs, total, nil
}

// Find -> calls Staff repo find method
func (s staffService) Find(input inputs.GetStaffDetailInput) (models.Staff, error) {
	staff, err := s.repository.Find(input.ID)

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
	staff, err := s.repository.Find(inputID.ID)
	if err != nil {
		return staff, err
	}

	staff.Name = input.Name
	staff.EmployeeCode = input.EmployeeCode
	staff.User.Username = input.Username
	staff.User.Email = input.Email
	staff.User.Password = input.Password

	// staff.User.Role = enums.Role(enums.Staff)
	staff.User.Role = "staff"

	if input.Password != "" {
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
		if err != nil {
			return staff, err
		}

		staff.User.Password = string(passwordHash)
	}

	updatedStaff, err := s.repository.Update(staff)
	if err != nil {
		return updatedStaff, err
	}

	return updatedStaff, nil
}

// Delete -> calls Staff repo delete method
func (s staffService) Delete(inputID inputs.GetStaffDetailInput) (models.Staff, error) {
	staff, err := s.repository.Find(inputID.ID)
	if err != nil {
		return staff, err
	}

	deletedStaff, err := s.repository.Delete(staff)
	if err != nil {
		return deletedStaff, err
	}

	return deletedStaff, nil
}
