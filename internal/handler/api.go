package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/waksun0x00/todoAPI/internal/tools"
)

func GetTodoList(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tools.GetTodo())
}
