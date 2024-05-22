package middlewares

import (
	"net/http"
	"strings"

	"github.com/IdaDanuartha/atv-backend-app/app/api/services"
	"github.com/IdaDanuartha/atv-backend-app/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authService *services.AuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			utils.ErrorJSON(ctx, http.StatusUnauthorized, "Unauthorized")
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			utils.ErrorJSON(ctx, http.StatusUnauthorized, "Unauthorized")
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			utils.ErrorJSON(ctx, http.StatusUnauthorized, "Unauthorized")
			return
		}

		userID := claim["user_id"].(string)

		user, err := authService.GetUserByID(userID)
		if err != nil {
			utils.ErrorJSON(ctx, http.StatusUnauthorized, "Unauthorized")
			return
		}

		ctx.Set("currentUser", user)
	}
}