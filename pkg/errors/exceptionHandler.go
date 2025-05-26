package errors

import (
	"errors"
	"fmt"
	"gindeu/pkg/globals"
	"github.com/go-playground/validator/v10"
	"reflect"
)

func getErrHandler() map[reflect.Type]func(err error) *AppError {
	return map[reflect.Type]func(err error) *AppError{

		reflect.TypeOf(validator.ValidationErrors{}): func(err error) *AppError {
			var validateErr validator.ValidationErrors
			errors.As(err, &validateErr)
			return NewAppError(globals.CodeFailed, validateErr[0].Translate(globals.ValidatorManager.GetTrans()))
		},

		reflect.TypeOf(&DemoError{}): func(err error) *AppError {
			var demoErr *DemoError
			errors.As(err, &demoErr)
			globals.EventBus.Publish(globals.DemoErrTopic, demoErr)
			return NewAppError(demoErr.Code, fmt.Sprintf("%s很漂亮", demoErr.Error()))
		},
	}
}

func ExceptionHandler(err error) *AppError {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr
	}

	errHandlers := getErrHandler()
	if v, ok := errHandlers[reflect.TypeOf(err)]; ok {
		return v(err)
	}

	return &AppError{Code: globals.CodeFailed, Message: err.Error()}
}
