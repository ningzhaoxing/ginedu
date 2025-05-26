package response

import (
	"gindeu/pkg/errors"
	"gindeu/pkg/globals"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(ctx *gin.Context, msg string, data interface{}) {
	if data == nil {
		data = ""
	}
	ctx.JSON(http.StatusOK, &AppData{
		Code: globals.CodeSuccess,
		Data: data,
		Msg:  msg,
	})
}

func Failed(ctx *gin.Context, err error) {
	appError := errors.ExceptionHandler(err)
	ctx.JSON(http.StatusOK, &AppData{
		Code: appError.Code,
		Msg:  appError.Message,
		Data: "",
	})
}
