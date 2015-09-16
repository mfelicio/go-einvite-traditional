package common

import (
	"einvite/framework"
)

func ToError(errorCode AppErrorCode, err error) *framework.FrameworkError {

	return &framework.FrameworkError{framework.ErrorType_App, int(errorCode), err.Error(), err}
}

func NewError(errorCode AppErrorCode, errorMessage string) *framework.FrameworkError {

	return &framework.FrameworkError{framework.ErrorType_App, int(errorCode), errorMessage, nil}
}

type AppErrorCode framework.FrameworkErrorCode

const (
//TODO: add business errors here

)
