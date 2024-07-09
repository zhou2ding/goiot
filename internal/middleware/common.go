package middleware

import (
	"github.com/gin-gonic/gin"
	"goiot/internal/global"
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
