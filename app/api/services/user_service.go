package services

import (
	"errors"

	"github.com/IdaDanuartha/atv-backend-app/app/api/inputs"
	"github.com/IdaDanuartha/atv-backend-app/app/api/repositories"
	// "github.com/IdaDanuartha/atv-backend-app/app/enums"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(input inputs.RegisterInput) (models.Customer, error)
	Login(input inputs.LoginInput) (models.User, error)
	GetUserByID(ID string) (models.User, error)
	GetAdminByUserID(ID string) (models.Admin, error)
	GetStaffByUserID(ID string) (models.Staff, error)
	GetInstructorByUserID(ID string) (models.Instructor, error)
	GetCustomerByUserID(ID string) (models.Customer, error)
	SaveAvatar(ID string, fileLocation string) (models.User, error)
	UpdateUser(input inputs.UpdateProfileInput, ctx *gin.Context) (models.User, error)
}

type userService struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) RegisterUser(input inputs.RegisterInput) (models.Customer, error) {
	customer := models.Customer{}
	customer.Name = input.Name
	customer.User.Username = input.Username
	customer.User.Email = input.Email
	// customer.User.Role = enums.Role(enums.Customer)
	customer.User.Role = "customer"

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return customer, err
	}

	customer.User.Password = string(passwordHash)

	newUser, err := s.repository.Save(customer)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *userService) Login(input inputs.LoginInput) (models.User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == "" {
		return models.User{}, errors.New("no user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *userService) GetUserByID(ID string) (models.User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == "" {
		return models.User{}, errors.New("no user found on with that ID")
	}

	return user, nil
}

func (s *userService) GetAdminByUserID(ID string) (models.Admin, error) {
	user, err := s.repository.FindAdminByUserID(ID, true)
	if err != nil {
		return user, err
	}

	if user.ID == "" {
		return models.Admin{}, errors.New("no admin found on with that ID")
	}

	return user, nil
}

func (s *userService) GetStaffByUserID(ID string) (models.Staff, error) {
	user, err := s.repository.FindStaffByUserID(ID, true)
	if err != nil {
		return user, err
	}

	if user.ID == "" {
		return models.Staff{}, errors.New("no staff found on with that ID")
	}

	return user, nil
}

func (s *userService) GetInstructorByUserID(ID string) (models.Instructor, error) {
	user, err := s.repository.FindInstructorByUserID(ID, true)
	if err != nil {
		return user, err
	}

	if user.ID == "" {
		return models.Instructor{}, errors.New("no instructor found on with that ID")
	}

	return user, nil
}

func (s *userService) GetCustomerByUserID(ID string) (models.Customer, error) {
	user, err := s.repository.FindCustomerByUserID(ID, true)
	if err != nil {
		return user, err
	}

	if user.ID == "" {
		return models.Customer{}, errors.New("no customer found on with that ID")
	}

	return user, nil
}

func (s *userService) SaveAvatar(ID string, fileLocation string) (models.User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	user.ProfilePath = &fileLocation

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *userService) UpdateUser(input inputs.UpdateProfileInput, ctx *gin.Context) (models.User, error) {
	userID := ctx.MustGet("currentUser").(models.User).ID

	user, err := s.repository.FindByID(userID)
	if err != nil {
		return user, err
	}

	if user.Role == "admin" {
		admin, _ := s.repository.FindAdminByUserID(userID, false)
		admin.Name = input.Name

		_, err := s.repository.UpdateAdmin(admin)
		if err != nil {
			return user, err
		}
	} else if user.Role == "staff" {
		staff, _ := s.repository.FindStaffByUserID(userID, false)
		staff.Name = input.Name
		staff.EmployeeCode = input.EmployeeCode

		_, err := s.repository.UpdateStaff(staff)
		if err != nil {
			return user, err
		}
	} else if user.Role == "instructor" {
		instructor, _ := s.repository.FindInstructorByUserID(userID, false)
		instructor.Name = input.Name
		instructor.EmployeeCode = input.EmployeeCode

		_, err := s.repository.UpdateInstructor(instructor)
		if err != nil {
			return user, err
		}
	} else if user.Role == "customer" {
		customer, _ := s.repository.FindCustomerByUserID(userID, false)
		customer.Name = input.Name
		customer.PhoneNumber = input.PhoneNumber

		_, err := s.repository.UpdateCustomer(customer)
		if err != nil {
			return user, err
		}
	}

	password := user.Password
	
	user.Username = input.Username
	user.Email = input.Email
	user.Password = password

	if input.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
		if err != nil {
			return user, err
		}

		user.Password = string(password)
	}

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}
