package router

import (
	"bubble/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", controllers.IndexHandler)

	v1Group := r.Group("/v1")
	{
		v1Group.POST("/todo", controllers.AddToDo)
		v1Group.GET("/todo", controllers.ListAllToDo)
		v1Group.PUT("/todo/:id", controllers.UpdateToDo)
		v1Group.DELETE("/todo/:id", controllers.DeleteTodo)
	}

	return r
}
