package handlers

import (
	"gindeu/internal/modules/user/logic"
	request "gindeu/internal/modules/user/requests"
	"gindeu/pkg/globals"
	"gindeu/pkg/response"
	"github.com/gin-gonic/gin"
)

// Login 登陆
func Login() gin.HandlerFunc {
	appCtx := globals.NewDefaultAppCtx()
	return func(context *gin.Context) {
		userLogic := logic.NewUserLogic(appCtx, context)
		var login request.Login
		err := context.ShouldBind(&login)
		if err != nil {
			response.Failed(context, err)
			return
		}
		err = userLogic.Login()
		if err != nil {
			response.Failed(context, err)
			return
		}
		response.Success(context, "", nil)
	}
}
