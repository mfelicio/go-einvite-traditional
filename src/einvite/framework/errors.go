package framework

import ()

type FrameworkError struct {
	ErrorType    ErrorType
	ErrorCode    FrameworkErrorCode "code"
	ErrorMessage string             "message"
	InnerError   error
}

func (this FrameworkError) Error() string {

	if this.ErrorMessage != "" {
		return this.ErrorMessage
	} else {
		return this.InnerError.Error()
	}
}

func ToError(errorCode FrameworkErrorCode, err error) *FrameworkError {

	return &FrameworkError{ErrorType_Framework, errorCode, err.Error(), err}
}

func NewError(errorCode FrameworkErrorCode, errorMessage string) *FrameworkError {

	return &FrameworkError{ErrorType_Framework, errorCode, errorMessage, nil}
}

type ErrorType int

const (
	ErrorType_Framework ErrorType = 1
	ErrorType_App       ErrorType = 2
)

type FrameworkErrorCode int

const (

	//Unknown errors range: 1-50
	Error_Generic FrameworkErrorCode = 1

	//Web errors range: 51-100
	Error_Web_SessionTampered       FrameworkErrorCode = 51
	Error_Web_SessionAlreadyHasUser FrameworkErrorCode = 52
	Error_Web_SessionAlreadyHasId   FrameworkErrorCode = 53
	Error_Web_SessionExpired        FrameworkErrorCode = 54
	Error_Web_SessionNotFound       FrameworkErrorCode = 55

	Error_Web_UnableToAuthenticate FrameworkErrorCode = 60

	//Db range: 101 - 200
	Error_Db_DuplicateId  FrameworkErrorCode = 101
	Error_Db_IdNotFound   FrameworkErrorCode = 102
	Error_Db_NetworkError FrameworkErrorCode = 103
)
