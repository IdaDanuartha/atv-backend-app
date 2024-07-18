package services

import (
	"github.com/IdaDanuartha/atv-backend-app/app/api/inputs"
	"github.com/IdaDanuartha/atv-backend-app/app/api/repositories"
	// "github.com/IdaDanuartha/atv-backend-app/app/enums"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"golang.org/x/crypto/bcrypt"
)

type InstructorService interface {
	FindAll(model models.Instructor, search string, currentPage int, pageSize int) ([]models.Instructor, int64, int, error)
	Find(input inputs.GetInstructorDetailInput) (models.Instructor, error)
	FindByUserID(input inputs.GetInstructorDetailInput) (models.Instructor, error)
	Save(input inputs.InstructorInput) (models.Instructor, error)
	Update(inputID inputs.GetInstructorDetailInput, input inputs.EditInstructorInput) (models.Instructor, error)
	Delete(inputID inputs.GetInstructorDetailInput) (models.Instructor, error)
}

// InstructorService InstructorService struct
type instructorService struct {
	repository repositories.InstructorRepository
	userRepository repositories.UserRepository
}

// NewInstructorService : returns the InstructorService struct instance
func NewInstructorService(repository repositories.InstructorRepository, userRepository repositories.UserRepository) instructorService {
	return instructorService{repository, userRepository}
}

// FindAll -> calls Instructor repo find all method
func (s instructorService) FindAll(model models.Instructor, search string, currentPage int, pageSize int) ([]models.Instructor, int64, int, error) {
	instructors, total, currentPage, err := s.repository.FindAll(model, search, currentPage, pageSize)
	if err != nil {
		return instructors, total, currentPage, err
	}

	return instructors, total, currentPage, nil
}

// Find -> calls Instructor repo find method
func (s instructorService) Find(input inputs.GetInstructorDetailInput) (models.Instructor, error) {
	instructor, err := s.repository.Find(input.ID, true)

	if err != nil {
		return instructor, err
	}

	return instructor, nil
}

// FindByUserID -> calls Instructor repo find method
func (s instructorService) FindByUserID(input inputs.GetInstructorDetailInput) (models.Instructor, error) {
	instructor, err := s.repository.Find(input.ID, true)

	if err != nil {
		return instructor, err
	}

	return instructor, nil
}

// Save -> calls Instructor repository save method
func (s instructorService) Save(input inputs.InstructorInput) (models.Instructor, error) {
	instructor := models.Instructor{}
	instructor.Name = input.Name
	instructor.EmployeeCode = input.EmployeeCode
	instructor.User.Username = input.Username
	instructor.User.Email = input.Email
	instructor.User.Password = input.Password

	// instructor.User.Role = enums.Role(enums.Instructor)
	instructor.User.Role = "instructor"

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return instructor, err
	}

	instructor.User.Password = string(passwordHash)

	newInstructor, err := s.repository.Save(instructor)
	if err != nil {
		return newInstructor, err
	}

	return newInstructor, nil
}

// Update -> calls Instructor repo update method
func (s instructorService) Update(inputID inputs.GetInstructorDetailInput, input inputs.EditInstructorInput) (models.Instructor, error) {
	instructor, err := s.repository.Find(inputID.ID, false)
	if err != nil {
		return instructor, err
	}

	instructor.Name = input.Name
	instructor.EmployeeCode = input.EmployeeCode

	user, _ := s.userRepository.FindByID(instructor.UserID)
	password := user.Password

	user.Username = input.Username
	user.Email = input.Email
	user.Password = password
	user.Role = "instructor"

	if input.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
		if err != nil {
			return instructor, err
		}

		user.Password = string(password)
	}

	_, err = s.userRepository.Update(user)
	if err != nil {
		return instructor, err
	}
	
	updatedInstructor, err := s.repository.Update(instructor)
	if err != nil {
		return updatedInstructor, err
	}

	return updatedInstructor, nil
}

// Delete -> calls Instructor repo delete method
func (s instructorService) Delete(inputID inputs.GetInstructorDetailInput) (models.Instructor, error) {
	instructor, err := s.repository.Find(inputID.ID, true)
	if err != nil {
		return instructor, err
	}

	deletedInstructor, err := s.repository.Delete(instructor)
	if err != nil {
		return deletedInstructor, err
	}

	return deletedInstructor, nil
}
