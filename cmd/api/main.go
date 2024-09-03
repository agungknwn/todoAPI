package main

import (
	"github.com/gin-gonic/gin"
	"github.com/waksun0x00/todoAPI/internal/handler"
)

func main() {
	router := gin.Default()
	router.GET("/TodoList", handler.GetTodoList)

	router.Run("localhost:6666")
}
