package services

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthService interface {
	GenerateToken(userID string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

var SECRET_KEY = []byte(os.Getenv("SECRET_KEY"))

func NewAuthService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID string) (string, error) {
	// Define the expiration time (one week from now)
    expirationTime := time.Now().Add(7 * 24 * time.Hour)

    // Create the JWT claims, including the user ID and expiration time
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     expirationTime.Unix(),
    }

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
