package main

import (
	"gin-app/controller"
	"gin-app/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	engine:= gin.Default()
	// middleware
	engine.Use(middleware.RecordUaAndTime)
	todoEngine := engine.Group("/todo")
	{
		v1 := todoEngine.Group("/v1")
		{
			v1.GET("/list", controller.TodoList)
			v1.POST("/add", controller.TodoAdd)
		}
	}
	engine.Run(":3000")
}