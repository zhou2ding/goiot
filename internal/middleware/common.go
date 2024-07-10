package middleware

import (
	"github.com/gin-gonic/gin"
	"goiot/internal/global"
	"goiot/internal/pkg/errcode"
	"golang.org/x/time/rate"
	"time"
)

func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		cost := time.Since(start).Milliseconds()
		c.Set(global.CostCtx, cost)
	}
}

var limiter = rate.NewLimiter(rate.Limit(1.667), 10) // 每秒允许1.667个请求（即每分钟最多允许100个请求），瞬间最多允许10个请求，
func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		if hasToken, exists := c.Get("hasToken"); exists && hasToken.(bool) {
			c.Next()
			return
		}
		if !limiter.Allow() {
			c.Set(global.ErrCtx, errcode.TooMandyRequests)
			c.Abort()
			return
		}
		c.Next()
	}
}
