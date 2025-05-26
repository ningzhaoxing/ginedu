package initialize

import "gindeu/pkg/globals"

func AppInit() {

	initEnv()

	var err error
	err, globals.C = InitConfig()
	if err != nil {
		return
	}

	globals.E = InitEngine()
	globals.Db = InitDb(globals.C)
	globals.EventBus = InitEventBus()
	globals.ValidatorManager = InitValidator()

}

func initEnv() {
	if len(globals.Env) == 0 {
		globals.Env = "local"
	}
}
