package initialize

import (
	"fmt"
	"gindeu/pkg/globals"
	"github.com/spf13/viper"
	"testing"
)

func InitConfig() (error, *globals.Config) {
	viper.SetConfigFile(getConfigPath())
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("config Load error %s \n", err.Error())
		return err, nil
	}

	var config globals.Config
	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("config bind error %s \n", err.Error())
		return err, nil
	}
	return nil, &config
}

func getConfigPath() string {

	var configPath string
	if testing.Testing() {
		configPath = "../configs/dev.yaml"
	} else {
		configPath = "./configs/dev.yaml"
	}

	return configPath
}
