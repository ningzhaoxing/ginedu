package globals

import (
	EventBus2 "github.com/asaskevich/EventBus"
	"github.com/gin-gonic/gin"
	"github.com/sonhineboy/gsadminValidator/ginValidator"
	"gorm.io/gorm"
)

var (
	E                *gin.Engine
	Db               *gorm.DB
	C                *Config
	Env              string
	EventBus         EventBus2.Bus
	ValidatorManager *ginValidator.CustomValidatorManager
)
