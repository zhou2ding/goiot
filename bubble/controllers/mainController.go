package controllers

import (
	"bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/*
	url    --> controller --> logic  --> model(CRUD)
	请求来了 --> 控制器      --> 业务逻辑 --> 增删改查
*/

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func AddToDo(c *gin.Context) {
	var todo models.Todo
	_ = c.ShouldBindJSON(&todo)
	err := models.CrateTodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		//c.JSON(http.StatusOK,gin.H{
		//	"code":2001,
		//	"msg":"insert data success",
		//	"data":todo,
		//})
		c.JSON(http.StatusOK, todo)
	}
}

func ListAllToDo(c *gin.Context) {
	todoList, err := models.ListAllToDo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateToDo(c *gin.Context) {
	var toto models.Todo
	_ = c.ShouldBindJSON(&toto)
	toto.Id, _ = strconv.Atoi(c.Param("id"))
	err := models.UpdateToDO(&toto)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, toto)
	}
}

func DeleteTodo(c *gin.Context) {
	var todo models.Todo
	todo.Id, _ = strconv.Atoi(c.Param("id"))
	err := models.DeleteToDo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"deleted_id": todo.Id})
	}
}
