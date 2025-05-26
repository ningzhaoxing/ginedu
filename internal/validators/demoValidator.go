package validators

import (
	"github.com/go-playground/validator/v10"
	"github.com/sonhineboy/gsadminValidator/ginValidator"
)

type DemoValidator struct {
	ginValidator.BaseValidator
}

func (d *DemoValidator) TagName() string {
	return "demo"
}

func (d *DemoValidator) Messages() string {
	return "{0}长度必须超过6个"
}

func (d *DemoValidator) Validator(fl validator.FieldLevel) bool {
	return len(fl.Field().String()) > 6
}
