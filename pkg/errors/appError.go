package errors

import "gindeu/pkg/globals"

type AppError struct {
	Code    globals.AppCode `json:"code"`  // HTTP状态码
	Message string          `json:"error"` // 错误信息
}

func NewAppError(code globals.AppCode, msg string) *AppError {
	return &AppError{
		Code:    code,
		Message: msg,
	}
}

func (e *AppError) Error() string {
	return e.Message
}
