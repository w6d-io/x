package errorx

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type Interface interface {
	// Error type representing an error condition, with the nil value representing no error.
	Error() string

	// GetCause type representing an error condition, with the nil value representing no error.
	GetCause() error

	// GetMessage returns the Message
	GetMessage() string

	// GetStatusCode returns http status code.
	GetStatusCode() int

	// ShowStack prints the stack if exist
	ShowStack()
}

type Error struct {
	// Cause original error
	Cause error

	// Code for http response code
	Code int

	// Message details for this error
	Message string
}

var _ Interface = &Error{}

var (
	// ErrTokenInvalid is raised when token is invalid
	ErrTokenInvalid = errors.New("invalid token")

	// ErrTokenCheck is raised when check validity token failed
	ErrTokenCheck = errors.New("check validity token failed")

	// ErrTokenNotFound is raise when HTTP header does not contain Authorization
	ErrTokenNotFound = errors.New("authorization not found in http header")

	// ErrUnAuthorized is raised by check token on kratos with a non authorized token
	ErrUnAuthorized = errors.New("401 Unauthorized")

	// ErrMethod is raised when a internal call failed
	ErrMethod = errors.New("resource can't be empty")

	// ErrServiceUnavailable is raised when a remote service fail to response
	ErrServiceUnavailable = errors.New("service unavailable")
)

func (e *Error) GetMessage() string {
	return e.Message
}

func (e *Error) GetStatusCode() int {
	return e.Code
}

func (e *Error) GetCause() error {
	return errors.Cause(e.Cause)
}

func (e *Error) ShowStack() {
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}

	if err, ok := errors.Cause(e.Cause).(stackTracer); ok {
		for _, f := range err.StackTrace() {
			fmt.Printf("%+s:%d\n", f, f)
		}
	}
}

func (e *Error) Error() string {
	if e.Cause == nil {
		return e.Message
	}
	if e.Message == "" {
		return fmt.Sprintf("%s", e.GetCause())
	}
	return fmt.Sprintf("%s: %s", e.GetMessage(), e.GetCause())
}

func New(err error, message string) error {
	return &Error{Cause: err, Message: message}
}

func NewHTTP(err error, code int, message string) error {
	return &Error{Cause: err, Code: code, Message: message}
}

func Wrap(err error, message string) error {
	return &Error{Cause: errors.Wrap(err, message)}
}

func GetError(err error) *Error {
	if err == nil {
		return nil
	}
	e, ok := err.(*Error)
	if ok {
		return e
	}
	return &Error{Cause: err}
}

// Error2code return the http code according the Error code field or try to deduce it from the error itself
func Error2code(err error) int {
	if e, ok := err.(*Error); ok {
		code := http.StatusInternalServerError
		if e.Code != 0 {
			code = e.Code
		}
		return code
	}
	switch err {
	case ErrTokenNotFound, ErrTokenInvalid, ErrUnAuthorized:
		return http.StatusUnauthorized
	case ErrMethod:
		return http.StatusBadRequest
	case ErrTokenCheck, ErrServiceUnavailable:
		return http.StatusServiceUnavailable
	}
	return http.StatusInternalServerError
}

// ErrorEncoder writes the error into http.ResponseWriter
func ErrorEncoder(err error, w http.ResponseWriter) {
	w.WriteHeader(Error2code(err))
	_ = json.NewEncoder(w).Encode(err.Error())
}
