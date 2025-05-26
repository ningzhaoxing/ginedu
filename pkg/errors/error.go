package errors

import "gindeu/pkg/globals"

var (
	ErrUserNotFound = NewAppError(globals.CodeHttpBad, "user Not found")
)
