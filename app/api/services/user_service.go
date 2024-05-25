package services

import (
	"errors"

	"github.com/IdaDanuartha/atv-backend-app/app/api/inputs"
	"github.com/IdaDanuartha/atv-backend-app/app/api/repositories"
	"github.com/IdaDanuartha/atv-backend-app/app/enums"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(input inputs.RegisterInput) (models.Customer, error)
	Login(input inputs.LoginInput) (models.User, error)
	GetUserByID(ID string) (models.User, error)
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
	customer.User.Role = enums.Role(enums.Customer)

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

	user.Username = input.Username
	user.Email = input.Email
	user.Role = input.Role

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}
