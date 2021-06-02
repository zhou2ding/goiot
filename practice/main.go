package main

import (
	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
}

func main() {
	r := gin.Default()
	r.GET("/hello",hello)
	_ = r.Run("127.0.0.1:8080")
}
