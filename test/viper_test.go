package test

import (
	"fmt"
	"gindeu/initialize"
	errors2 "gindeu/pkg/errors"
	"gindeu/pkg/globals"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"testing"
)

func TestViper(t *testing.T) {

	var c globals.Config
	viper.SetConfigFile("../configs/dev.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&c)
	if err != nil {
		return
	}
	fmt.Println(c.DB.Config)
}

func TestInitConfig(t *testing.T) {
	initialize.InitConfig()
	fmt.Println(viper.GetString("db.types") == "")
}

func TestPath(t *testing.T) {
	d, _ := os.Getwd()
	fmt.Println(d, string(filepath.Separator), "configs")
}

func TestErr(t *testing.T) {

	myErr := errors.WithMessage(&errors2.AppError{Code: globals.CodeFailed, Message: "asdfa"}, "xxxx")

	var err *errors2.AppError
	if errors.As(myErr, &err) {
		fmt.Println(err.Code)

	}

	fmt.Printf("%v", myErr)

}
