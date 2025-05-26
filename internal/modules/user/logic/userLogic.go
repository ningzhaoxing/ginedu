package logic

import (
	"fmt"
	"gindeu/internal/modules/user/models"
	"gindeu/pkg/errors"
	"gindeu/pkg/globals"
	"gindeu/pkg/subscribes"
	"github.com/gin-gonic/gin"
)

type UserLogic struct {
	appCtx *globals.AppCtx
	ctx    *gin.Context
	//ctx        string
	//LabelLogic string
	//Jwt string
}

func NewUserLogic(appCtx *globals.AppCtx, ctx *gin.Context) *UserLogic {
	return &UserLogic{
		appCtx: appCtx,
		ctx:    ctx,
	}
}

func (user *UserLogic) Login() error {
	return &errors.DemoError{
		Code: 30,
		User: "张珊",
	}
}

func (user *UserLogic) setUserInfo(User models.User, t string) {

}

func (user *UserLogic) SetUserIntegral(userName string, num int) {
	fmt.Println("修改了用户积分", num, userName)
	user.appCtx.GetEventBus().Publish(globals.UserIntegralTopic, subscribes.UserIntegralTopic{
		User:           userName,
		BeforeIntegral: 20,
		AfterIntegral:  30,
	})
}
