package globalmw

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"goiot/pkg/defs"
	"goiot/pkg/errcode"
	"goiot/pkg/jwtAuth"
	"strings"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := GetTokenFromHeader(c)
		if tokenString == "" {
			c.Set(defs.ErrCtx, errcode.TokenAuthFail)
			c.Abort()
			return
		}

		uc, err := jwtAuth.ParseToken(tokenString)
		if err != nil {
			var validErr *jwt.ValidationError
			if errors.As(err, &validErr) && validErr.Errors&jwt.ValidationErrorExpired != 0 {
				c.Set(defs.ErrCtx, errcode.AccessTokenExpiredError)
			} else {
				c.Set(defs.ErrCtx, errcode.TokenAuthFail)
			}
			c.Abort()
			return
		}
		c.Set(defs.UserIDCtx, uc.UserId)
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
