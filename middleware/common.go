package middleware

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"goiot/pkg/errcode"
	"goiot/pkg/global"
	"goiot/pkg/jwtAuth"
	"goiot/pkg/logger"
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
		c.Set(global.CostCtx, cost)
	}
}

func CommonResponse() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(global.RequestIdCtx, utils.GetUUIDFull())
		c.Next()

		if c.Writer.Written() {
			return
		}

		var (
			cost  = c.GetInt64(global.CostCtx)
			reqId = c.GetString(global.RequestIdCtx)
			resp  = response{
				RequestId: reqId,
				Duration:  cost,
			}
		)
		if err, ok := c.Get(global.ErrCtx); ok {
			ec := err.(errcode.ErrCode)
			resp.Code, resp.Message = int(ec), ec.String()

			if ec >= errcode.TokenAuthFail || ec <= errcode.RefreshTokenExpiredError {
				c.JSON(http.StatusUnauthorized, resp)
			} else {
				errMsg, _ := c.Get(global.ErrMsgCtx)
				if errMsg != nil {
					resp.Message = errMsg
					c.JSON(http.StatusBadRequest, resp)
				} else {
					c.JSON(http.StatusInternalServerError, resp)
				}
			}
		} else {
			data, _ := c.Get(global.DataCtx)
			resp.Code, resp.Data = http.StatusOK, data
			c.JSON(http.StatusOK, resp)
		}
	}
}

var limiter = rate.NewLimiter(rate.Limit(1.667), 10) // 每秒允许1.667个请求（即每分钟最多允许100个请求），瞬间最多允许10个请求，
func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.Set(global.ErrCtx, errcode.TooMandyRequests)
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
			global.RequestIdCtx, requestId,
			global.RemoteIpCtx, m.localIP,
		)
		// api层内部传递的上下文
		ctx = context.WithValue(ctx, global.RequestIdCtx, requestId)
		ctx = context.WithValue(ctx, global.StartTimeCtx, time.Now())
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
				// todo 共用返回逻辑
			} else {
				// todo 共用返回逻辑
			}
			return
		}
		// 传递给rpc服务的上下文
		ctx := metadata.AppendToOutgoingContext(r.Context(),
			global.UserIDCtx, uc.UserId,
		)
		// api层内部传递的上下文
		ctx = context.WithValue(ctx, global.UserIDCtx, uc.UserId)
		next(w, r.WithContext(ctx))
	}
}
