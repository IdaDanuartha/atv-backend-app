package formatters

import (
	"github.com/IdaDanuartha/atv-backend-app/app/enums"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

type CustomerFormatter struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Token    string     `json:"token"`
	User 	 models.User	`json:"user"`
}

func FormatCustomer(customer models.Customer, token string) CustomerFormatter {
	formatter := CustomerFormatter{
		ID:       customer.ID,
		Name:     customer.Name,
		Token:    token,
		User: customer.User,
	}

	return formatter
}

type AuthFormatter struct {
	ID       string     `json:"id"`
	Username string     `json:"username"`
	Email    string     `json:"email"`
	Token    string     `json:"token"`
	Role     enums.Role `json:"role"`
}

type UserFormatter struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func FormatAuth(user models.User, token string) AuthFormatter {
	formatter := AuthFormatter{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
		Role:     user.Role,
	}

	return formatter
}

func FormatUser(user models.User) UserFormatter {
	formatter := UserFormatter{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	return formatter
}
