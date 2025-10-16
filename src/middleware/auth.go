package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct {
	jwtSecretToken string
}

func NewAuthMiddleware(jwtSecretToken string) AuthMiddleware {
	return AuthMiddleware{
		jwtSecretToken: jwtSecretToken,
	}
}

func (ref *AuthMiddleware) Auth(ctx *gin.Context) {
	if gin.Mode() == gin.TestMode {
		return
	}

	tokenStr := ctx.GetHeader("Authorization")

	if tokenStr == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token not provided"})
		return
	}

	splittedToken := strings.Split(tokenStr, " ")

	if len(splittedToken) <= 1 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token is invalid"})
		return
	}

	tokenStr = splittedToken[1]

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
		return []byte(ref.jwtSecretToken), nil
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if !token.Valid {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token is invalid"})
		return
	}

	claims := token.Claims.(jwt.MapClaims)

	ctx.Set("user_id", claims["user_id"])

	ctx.Next()
}
