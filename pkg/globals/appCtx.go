package globals

import (
	EventBus2 "github.com/asaskevich/EventBus"
	"gorm.io/gorm"
)

type DbType string

var (
	Es        DbType = "es"
	Mysql     DbType = "mysql"
	Sqlserver DbType = "Sqlserver"
)

type AppCtx struct {
	db       *gorm.DB
	config   *Config
	eventBus EventBus2.Bus
}

type Options func(a *AppCtx)

func NewAppCtx(op ...Options) *AppCtx {
	appCtx := new(AppCtx)
	for _, options := range op {
		options(appCtx)
	}
	return appCtx
}

func NewDefaultAppCtx() *AppCtx {
	return NewAppCtx(WithOptionDb(Db), WithOptionConfig(C), WithOptionEventBus(EventBus))
}

func WithOptionDb(db *gorm.DB) Options {
	return func(a *AppCtx) {
		a.db = db
	}
}

func WithOptionConfig(c *Config) Options {
	return func(a *AppCtx) {
		a.config = c
	}
}

func WithOptionEventBus(e EventBus2.Bus) Options {
	return func(a *AppCtx) {
		a.eventBus = e
	}
}

func (a *AppCtx) GetEventBus() EventBus2.Bus {
	return a.eventBus
}

func (a *AppCtx) SetDb(db *gorm.DB) {
	a.db = db
}

func (a *AppCtx) SetConfig() {

}

func (a *AppCtx) GetDb(dbType DbType) *gorm.DB {
	switch dbType {
	case Es:
		return a.db
	case Mysql:
		return a.db
	case Sqlserver:
		return a.db
	}
	return a.db
}
