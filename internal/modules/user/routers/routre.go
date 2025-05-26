package routers

import (
	"gindeu/internal/modules/user/handlers"
	"github.com/gin-gonic/gin"
)

func InitRouter(e *gin.Engine) {
	e.POST("/login", handlers.Login())
}
