package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"goiot/internal/global"
	"goiot/internal/pkg/errcode"
	"goiot/internal/pkg/jwtAuth"
	"strings"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := GetTokenFromHeader(c)
		if tokenString == "" {
			c.Set(global.ErrCtx, errcode.TokenAuthFail)
			c.Abort()
			return
		}

		uc, err := jwtAuth.ParseToken(tokenString)
		if err != nil {
			var validErr *jwt.ValidationError
			if errors.As(err, &validErr) && validErr.Errors&jwt.ValidationErrorExpired != 0 {
				c.Set(global.ErrCtx, errcode.AccessTokenExpiredError)
			} else {
				c.Set(global.ErrCtx, errcode.TokenAuthFail)
			}
			c.Abort()
			return
		}
		c.Set(global.UserIDCtx, uc.UserId)
		c.Next()
	}
}

func GetTokenFromHeader(c *gin.Context) string {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return ""
	}

	return parts[1]
}
