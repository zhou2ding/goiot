package globalmw

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"goiot/pkg/cache"
	"goiot/pkg/defs"
	"goiot/pkg/errcode"
	"goiot/pkg/jwtAuth"
	"goiot/pkg/logger"
	"goiot/pkg/result"
	"goiot/pkg/utils"
	"golang.org/x/time/rate"
	"google.golang.org/grpc/metadata"
	"net/http"
	"time"
)

type response struct {
	Code      int    `json:"code"`
	Message   any    `json:"message,omitempty"`
	Data      any    `json:"data,omitempty"`
	RequestId string `json:"requestId"`
	Duration  int64  `json:"duration"`
}

func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		cost := time.Since(start).Milliseconds()
		c.Set(defs.CostCtx, cost)
	}
}

func CommonResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(defs.RequestIdCtx, utils.GetUUIDFull())
		c.Next()

		if c.Writer.Written() {
			return
		}

		var (
			cost  = c.GetInt64(defs.CostCtx)
			reqId = c.GetString(defs.RequestIdCtx)
			resp  = response{
				RequestId: reqId,
				Duration:  cost,
			}
		)
		if err, ok := c.Get(defs.ErrCtx); ok {
			ec := err.(errcode.ErrCode)
			resp.Code, resp.Message = int(ec), ec.String()

			if ec >= errcode.TokenAuthFail || ec <= errcode.RefreshTokenExpiredError {
				c.JSON(http.StatusUnauthorized, resp)
			} else {
				errMsg, _ := c.Get(defs.ErrMsgCtx)
				if errMsg != nil {
					resp.Message = errMsg
					c.JSON(http.StatusBadRequest, resp)
				} else {
					c.JSON(http.StatusInternalServerError, resp)
				}
			}
		} else {
			data, _ := c.Get(defs.DataCtx)
			resp.Code, resp.Data = http.StatusOK, data
			c.JSON(http.StatusOK, resp)
		}
	}
}

var limiter = rate.NewLimiter(rate.Limit(1.667), 10) // 每秒允许1.667个请求（即每分钟最多允许100个请求），瞬间最多允许10个请求，
func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.Set(defs.ErrCtx, errcode.TooMandyRequests)
			c.Abort()
			return
		}
		c.Next()
	}
}

func NotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
}

type ProcessReqRespMiddleware struct {
	localIP string
}

func NewProcessReqRespMiddleware() *ProcessReqRespMiddleware {
	return &ProcessReqRespMiddleware{utils.GetLocalIP()}
}

func (m *ProcessReqRespMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestId := utils.GetUUIDFull()
		logger.Logger.Infof("get requeest from %s, requestId: %s", r.RemoteAddr, requestId)
		// 传递给rpc服务的上下文
		ctx := metadata.AppendToOutgoingContext(r.Context(),
			defs.RequestIdCtx, requestId,
			defs.RemoteIpCtx, m.localIP,
		)
		// api层内部传递的上下文
		ctx = context.WithValue(ctx, defs.RequestIdCtx, requestId)
		ctx = context.WithValue(ctx, defs.StartTimeCtx, time.Now())
		next(w, r.WithContext(ctx))
	}
}

type AuthMiddleware struct {
	secret string
}

func NewAuthMiddleware(secret string) *AuthMiddleware {
	return &AuthMiddleware{secret}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		uc := new(jwtAuth.UserClaims)
		_, err := jwt.ParseWithClaims(tokenString, uc, func(token *jwt.Token) (any, error) {
			return []byte(m.secret), nil
		})
		if err != nil {
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

type ApiKeyMiddleware struct {
}

func NewApiKeyMiddleware() *ApiKeyMiddleware {
	return &ApiKeyMiddleware{}
}

func (m *ApiKeyMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("x-api-key")
		if apiKey == "" || !isValidApiKey(apiKey) {
			result.ErrorResultWithCode(r.Context(), w, http.StatusForbidden, errcode.APIKeyAuthFail)
			return
		}
		next(w, r)
	}
}

func isValidApiKey(apiKey string) bool {
	return cache.GetRedis(cache.PermissionDB).Exists(cache.APIKeyKey+apiKey).Val() > 0
}
