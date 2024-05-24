package formatters

import "github.com/IdaDanuartha/atv-backend-app/app/models"

type UserFormatter struct {
	ID         string    `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

func FormatUser(user models.User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:         user.ID,
		Username: 	user.Username,
		Email:      user.Email,
		Token:      token,
	}

	return formatter
}
