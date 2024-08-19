package middleware

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"goiot/pkg/defs"
	"goiot/pkg/errcode"
	"goiot/pkg/jwtAuth"
	"goiot/pkg/logger"
	"goiot/pkg/result"
	"google.golang.org/grpc/metadata"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	secret string
}

func NewAuthMiddleware(secret string) *AuthMiddleware {
	return &AuthMiddleware{secret}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		parts := strings.SplitN(tokenString, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			result.ErrorResultWithCode(r.Context(), w, http.StatusUnauthorized, errcode.TokenAuthFail)
		}

		uc := new(jwtAuth.UserClaims)
		_, err := jwt.ParseWithClaims(parts[1], uc, func(token *jwt.Token) (any, error) {
			return []byte(m.secret), nil
		})
		if err != nil {
			logger.Logger.Errorf("requestId: %v parse token error: %v", r.Context().Value(defs.RequestIdCtx), err)
			var validErr *jwt.ValidationError
			if errors.As(err, &validErr) && validErr.Errors&jwt.ValidationErrorExpired != 0 {
				result.ErrorResultWithCode(r.Context(), w, http.StatusUnauthorized, errcode.AccessTokenExpiredError)
			} else {
				result.ErrorResultWithCode(r.Context(), w, http.StatusUnauthorized, errcode.TokenAuthFail)
			}
			return
		}

		// 传递给rpc服务的上下文
		ctx := metadata.AppendToOutgoingContext(r.Context(),
			defs.UserIDCtx, uc.UserId,
		)
		// api层内部传递的上下文
		ctx = context.WithValue(ctx, defs.UserIDCtx, uc.UserId)
		next(w, r.WithContext(ctx))
	}
}
