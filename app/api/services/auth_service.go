package services

import (
	"errors"

	"github.com/IdaDanuartha/atv-backend-app/app/api/inputs"
	"github.com/IdaDanuartha/atv-backend-app/app/api/repositories"
	"github.com/IdaDanuartha/atv-backend-app/app/enums"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// AuthService AuthService struct
type ServiceInterface interface {
	GenerateToken(userID string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
	GetUserByID(ID string) (models.User, error)
}

type AuthService struct {
	repository repositories.AuthRepository
}

// NewAuthService : returns the AuthService struct instance
func NewAuthService(r repositories.AuthRepository) AuthService {
	return AuthService{
		repository: r,
	}
}

var SECRET_KEY = []byte("atv_system_s3cr3T_k3Y")

func (s *AuthService) GenerateToken(userID string) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (s *AuthService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid token")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}

func (s *AuthService) RegisterUser(input inputs.RegisterInput) (models.Customer, error) {
	customer := models.Customer{}
	customer.User.Username = input.Username
	customer.User.Email = input.Email
	customer.User.Role = enums.Customer
	customer.Name = input.Name

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

func (s *AuthService) Login(input inputs.LoginInput) (models.User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	// if user.ID == 0 {
	// 	return models.User, errors.New("No user found on that email")
	// }

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *AuthService) GetUserByID(ID string) (models.User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	return user, nil
}
