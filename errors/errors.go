package errors

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/goadesign/goa"
	"runtime"
	"net/http"
)

type ServiceError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Line    int    `json:"-"`
	File    string `json:"-"`
}

type ServiceErrorLogger interface {
	goa.ServiceError
	Log()
}

func (se *ServiceError) Error() string {
	return se.Message
}

func (se *ServiceError) Log() {
	logrus.Error(fmt.Sprintf("%s context: File %s Line %d ", se.Message, se.File, se.Line))
}
func (se *ServiceError) Token() string {
	return "random"
}

func (se *ServiceError) ResponseStatus() int {
	return se.Code
}

func NewServiceError(msg string, code int) ServiceErrorLogger {
	_, f, n, _ := runtime.Caller(1)
	return &ServiceError{Message: msg, Code: code, Line: n, File: f}
}

func LogAndReturnErrorWithCode(err error, code int) ServiceErrorLogger {
	if e, ok := err.(ServiceErrorLogger); ok {
		e.Log()
		return e
	}
	_, f, n, _ := runtime.Caller(1)
	wErr := &ServiceError{Message: err.Error(), Code: code, Line: n, File: f}
	wErr.Log()
	return wErr
}
func LogAndReturnError(err error) ServiceErrorLogger {
	if e, ok := err.(ServiceErrorLogger); ok {
		e.Log()
		return e
	}
	_, f, n, _ := runtime.Caller(1)
	wErr := &ServiceError{Message: err.Error(), Code: http.StatusInternalServerError, Line: n, File: f}
	wErr.Log()
	return wErr
}
