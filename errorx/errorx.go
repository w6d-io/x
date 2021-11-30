package errorx

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/transport"

	"github.com/w6d-io/x/logx"

	"github.com/pkg/errors"
)

// Interface of errorx
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

	// EditCause field from Error struct then return it
	EditCause(error) *Error

	// EditMessage field from Error struct then return it
	EditMessage(string) *Error

	// EditCode field from Error struct then return it
	EditCode(string) *Error

	// EditStatusCode field from Error struct then return it
	EditStatusCode(int) *Error
}

// Error ...
type Error struct {
	// Cause original error
	Cause error `json:"-"`

	// StatusCode for http response code
	StatusCode int `json:"-"`

	// Code is the string convention from https://www.notion.so/w6d/Project-Spec-d1e49d91046b4b61952ecfc135983faf
	Code string `json:"code"`

	// Message from this error
	Message string `json:"message"`
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

// EditCause returns Error by updating cause
func (e Error) EditCause(err error) *Error {
	e.Cause = err
	return &e
}

// EditMessage returns Error by updating message
func (e Error) EditMessage(msg string) *Error {
	e.Message = msg
	return &e
}

// EditCode returns Error by updating code
func (e Error) EditCode(code string) *Error {
	e.Code = code
	return &e
}

// EditStatusCode returns Error by updating status code
func (e Error) EditStatusCode(statusCode int) *Error {
	e.StatusCode = statusCode
	return &e
}

// GetMessage returns the message from Error
func (e *Error) GetMessage() string {
	return e.Message
}

// GetStatusCode returns the http status code from Error
func (e *Error) GetStatusCode() int {
	return e.StatusCode
}

// GetCause returns the cause from Error
func (e *Error) GetCause() error {
	return errors.Cause(e.Cause)
}

// ShowStack prints the stack trace if available
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

// Error return a custom message from Error
func (e *Error) Error() string {
	if e.Cause == nil {
		return e.Message
	}
	if e.Message == "" {
		return fmt.Sprintf("%s", e.GetCause())
	}
	return fmt.Sprintf("%s: %s", e.GetMessage(), e.GetCause())
}

// New returns Error with err in Cause and message
func New(err error, message string) error {
	return &Error{Cause: err, Message: message}
}

// NewHTTP returns Error with err in Cause, code in http.StatusCode and message
func NewHTTP(err error, code int, message string) error {
	return &Error{Cause: err, StatusCode: code, Message: message}
}

// Wrap the error with Error type
func Wrap(err error, message string) error {
	return &Error{Cause: errors.Wrap(err, message), Message: message}
}

// GetError builds or gets the Error type
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
		if e.StatusCode != 0 {
			code = e.StatusCode
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
func ErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	logx.WithName(ctx, "ErrorEncoder").Error(err, "")
	w.WriteHeader(Error2code(err))
	if e, ok := err.(*Error); ok {
		_ = json.NewEncoder(w).Encode(e)
		return
	}
	_ = json.NewEncoder(w).Encode(err.Error())
}

//
// GRPC stuff
//

// NewErrorHandler returns an instance of transport.ErrorHandler
func NewErrorHandler() transport.ErrorHandler {
	return &errorHandler{}
}

type errorHandler struct{}

// Handle write the error in log
func (errorHandler) Handle(ctx context.Context, err error) {
	logx.WithName(ctx, "ErrorHandler").Error(err, "")
}
