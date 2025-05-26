package server

import (
	user "gindeu/internal/modules/user/routers"
	"github.com/gin-gonic/gin"
)

func InitRouter(e *gin.Engine) {
	// 这种代码不应该再这里出现 START
	//e.GET("/success", func(context *gin.Context) {
	//	response.Success(context, "ok", "")
	//})
	//
	//e.GET("/error", func(context *gin.Context) {
	//	response.Failed(context, globals.CodeFailed, " err")
	//
	//})
	// 这种代码不应该再这里出现 END

	//这种正确的做法
	user.InitRouter(e)
	//sse.InitRouter(e)
	//xxx.InitRouter(e)
}
