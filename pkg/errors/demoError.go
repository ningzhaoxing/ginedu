package errors

import "gindeu/pkg/globals"

type DemoError struct {
	Code globals.AppCode
	User string
}

func (e *DemoError) Error() string {
	return e.User
}
