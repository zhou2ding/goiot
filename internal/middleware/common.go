package middleware

import (
	"github.com/gin-gonic/gin"
	"goiot/internal/global"
	"goiot/internal/pkg/errcode"
	"goiot/internal/pkg/utils"
	"golang.org/x/time/rate"
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
