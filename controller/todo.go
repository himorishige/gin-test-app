package controller

import (
	"gin-app/model"
	"gin-app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TodoList(c *gin.Context) {
	todoService := service.TodoService{}
	TodoLists := todoService.GetTodoList()
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data": TodoLists,
	})
}

func TodoAdd(c *gin.Context) {
	todo := model.Todo{}
	err := c.Bind(todo)
	if err != nil{
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	todoService := service.TodoService{}
	err = todoService.SetTodo(&todo)
	c.JSON(http.StatusCreated, gin.H{
		"status" : "ok",
	})
}