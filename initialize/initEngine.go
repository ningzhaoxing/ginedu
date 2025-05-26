package initialize

import (
	"github.com/gin-gonic/gin"
)

func InitEngine() *gin.Engine {
	e := gin.Default()
	e.Use(gin.Recovery())
	return e
}
