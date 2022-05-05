package errors

import (
	"fmt"
)

const (
	RequestTimeOut = iota
)

const (
	ClientError = iota
	ServerError
	UnknownError
)

var ErrorMap = map[int]string {
	0: "ClientError",
	1: "ServerError",
	2: "UnknownError",
}

type AppError struct {
	ErrorKind int
	AppErrID int
	HttpStatusCode int
	ClientErrMessage string
	ServerErrMessage string
	ErrMessage string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("kind=%d, appErrID=%s, httpStatusCode:%d, errMessage:%s", e.ErrorKind, ErrorMap[e.AppErrID], e.HttpStatusCode, e.ClientErrMessage)
}

func NewAppError(kind int, statusCode int, message string) *AppError {
	appErr := &AppError{
		ErrorKind:      kind,
		HttpStatusCode: statusCode,
	}
	switch kind {
	case ClientError:
		appErr.ClientErrMessage = message
	case ServerError:
		appErr.ServerErrMessage = message
	case UnknownError:
		appErr.ErrMessage = message
	}
	return appErr
}