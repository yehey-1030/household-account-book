package errmodel

import (
	"errors"
	"fmt"
	"github.com/yehey-1030/household-account-book/go/util/ioutil"
	"net/http"
)

var (
	ErrBadRequest            = NewBadRequestError("", nil)
	ErrUnauthorized          = NewUnauthorizedError("", nil)
	ErrPasswordMismatch      = NewPasswordMismatchError("", nil)
	ErrForbidden             = NewForbiddenError("", nil)
	ErrNotFound              = NewNotFoundError("", nil)
	ErrConflict              = NewConflictError("", nil)
	ErrInternalServer        = NewInternalServerError("", nil)
	ErrNotExist              = NewNotExistError("", nil)
	ErrRequestEntityTooLarge = NewRequestEntityTooLargeError("", nil)
	ErrUnprocessableEntity   = NewUnprocessableEntityError("", nil)
	ErrTooManyRequests       = NewTooManyRequestsError("", nil)
	ErrNotAcceptable         = NewNotAcceptableError("", nil)
	ErrUnsupportedType       = NewUnsupportedTypeError("", nil)
	ErrLocked                = NewLockedError("", nil)
	ErrFailedDependencyError = NewFailedDependencyError("", nil)
)

type baseError struct {
	msg     string
	wrapped error
}

func newError(errType error, msg string, wrapped error) *baseError {
	message := fmt.Sprintf("%T", errType)
	if msg != "" {
		message = message + ": " + msg
	}
	return &baseError{message, wrapped}
}
func (n *baseError) Error() string {
	if n.wrapped != nil {
		return fmt.Sprintf("%s, caused by %v", n.msg, n.wrapped)
	} else {
		return n.msg
	}
}
func (n *baseError) Unwrap() error { return n.wrapped }

type NotExistError struct {
	*baseError
}

func NewNotExistError(msg string, wrapped error) *NotExistError {
	err := NotExistError{}
	err.baseError = newError(err, msg, wrapped)
	return &err
}

type UnauthorizedError struct {
	*baseError
}

func NewUnauthorizedError(msg string, wrapped error) *UnauthorizedError {
	err := UnauthorizedError{}
	err.baseError = newError(err, msg, wrapped)
	return &err
}

type PasswordMismatchError struct {
	*baseError
}

func NewPasswordMismatchError(msg string, wrapped error) *PasswordMismatchError {
	err := PasswordMismatchError{}
	err.baseError = newError(err, msg, wrapped)
	return &err
}

type NotFoundError struct {
	*baseError
}

func NewNotFoundError(msg string, wrapped error) *NotFoundError {
	err := NotFoundError{}
	err.baseError = newError(err, msg, wrapped)
	return &err
}

type ForbiddenError struct {
	*baseError
}

func NewForbiddenError(msg string, wrapped error) *ForbiddenError {
	err := ForbiddenError{}
	err.baseError = newError(err, msg, wrapped)
	return &err
}

type BadRequestError struct {
	*baseError
}

func NewBadRequestError(msg string, wrapped error) *BadRequestError {
	err := BadRequestError{}
	err.baseError = newError(err, msg, wrapped)
	return &err
}

type ConflictError struct {
	*baseError
}

func NewConflictError(msg string, wrapped error) *ConflictError {
	err := ConflictError{}
	err.baseError = newError(err, msg, wrapped)
	return &err
}

type InternalServerError struct {
	*baseError
}

func NewInternalServerError(msg string, wrapped error) *InternalServerError {
	err := InternalServerError{}
	err.baseError = newError(err, msg, wrapped)
	return &err
}

type RequestEntityTooLargeError struct {
	*baseError
}

func NewRequestEntityTooLargeError(msg string, wrapped error) *RequestEntityTooLargeError {
	err := RequestEntityTooLargeError{}
	err.baseError = newError(err, msg, wrapped)
	return &err
}

type UnprocessableEntityError struct {
	*baseError
}

func NewUnprocessableEntityError(msg string, wrapped error) *UnprocessableEntityError {
	err := UnprocessableEntityError{}
	err.baseError = newError(err, msg, wrapped)
	return &err
}

type TooManyRequestsError struct {
	*baseError
}

func NewTooManyRequestsError(msg string, wrapped error) *TooManyRequestsError {
	err := TooManyRequestsError{}
	err.baseError = newError(err, msg, wrapped)
	return &err
}

type NotAcceptableError struct {
	*baseError
}

func NewNotAcceptableError(msg string, wrapped error) *NotAcceptableError {
	err := NotAcceptableError{}
	err.baseError = newError(err, msg, wrapped)
	return &err
}

type UnsupportedTypeError struct {
	*baseError
}

func NewUnsupportedTypeError(msg string, wrapped error) *UnsupportedTypeError {
	err := UnsupportedTypeError{}
	err.baseError = newError(err, msg, wrapped)
	return &err
}

type LockedError struct {
	*baseError
}

func NewLockedError(msg string, wrapped error) *LockedError {
	err := LockedError{}
	err.baseError = newError(err, msg, wrapped)
	return &err
}

type FailedDependencyError struct {
	*baseError
}

func NewFailedDependencyError(msg string, wrapped error) *FailedDependencyError {
	err := FailedDependencyError{}
	err.baseError = newError(err, msg, wrapped)
	return &err
}

func ErrResponseFrom(err error) (int, string, string) {
	var statusCode int

	if errors.As(err, &ErrNotExist) {
		statusCode = http.StatusNotFound
	} else if errors.As(err, &ErrBadRequest) {
		statusCode = http.StatusBadRequest
	} else if errors.As(err, &ErrUnauthorized) {
		statusCode = http.StatusUnauthorized
	} else if errors.As(err, &ErrPasswordMismatch) {
		statusCode = http.StatusPaymentRequired
	} else if errors.As(err, &ErrForbidden) {
		statusCode = http.StatusForbidden
	} else if errors.As(err, &ErrNotFound) {
		statusCode = http.StatusNotFound
	} else if errors.As(err, &ErrConflict) {
		statusCode = http.StatusConflict
	} else if errors.As(err, &ErrRequestEntityTooLarge) {
		statusCode = http.StatusRequestEntityTooLarge
	} else if errors.As(err, &ErrUnprocessableEntity) {
		statusCode = http.StatusUnprocessableEntity
	} else if errors.As(err, &ErrTooManyRequests) {
		statusCode = http.StatusTooManyRequests
	} else if errors.As(err, &ErrNotAcceptable) {
		statusCode = http.StatusNotAcceptable
	} else if errors.As(err, &ErrUnsupportedType) {
		statusCode = http.StatusUnsupportedMediaType
	} else if errors.As(err, &ErrLocked) {
		statusCode = http.StatusLocked
	} else if errors.As(err, &ErrFailedDependencyError) {
		statusCode = http.StatusFailedDependency
	} else if errors.As(err, &ErrInternalServer) {
		statusCode = http.StatusInternalServerError

	} else {
		if errors.Is(err, ioutil.ErrTimeout) {
			statusCode = http.StatusRequestTimeout
		} else {
			statusCode = http.StatusInternalServerError
		}
	}

	return statusCode, ErrorMsgFrom(statusCode), ErrorPageFrom(statusCode)
}

func ToDomainError(msg string, err error) error {
	switch err.(type) {
	case *UnauthorizedError:
		return NewUnauthorizedError(msg, err)
	case *NotFoundError:
		return NewNotFoundError(msg, err)
	case *ForbiddenError:
		return NewForbiddenError(msg, err)
	case *BadRequestError:
		return NewBadRequestError(msg, err)
	case *ConflictError:
		return NewConflictError(msg, err)
	default:
		return NewInternalServerError(msg, err)
	}
}

const (
	ERR_RES_MSG_4XX = "Please refer to the error code in the guide document"
	ERR_RES_MSG_5XX = "Please contact the administrator"
)

func ErrorMsgFrom(code int) string {
	if code >= 400 && code < 500 {
		return ERR_RES_MSG_4XX
	} else if code >= 500 {
		return ERR_RES_MSG_5XX
	}

	return ""
}

func ErrorPageFrom(code int) string {
	switch code {
	case http.StatusBadRequest:
		return "status400.html"
	case http.StatusUnauthorized:
		return "status401.html"
	case http.StatusForbidden:
		return "status403.html"
	case http.StatusNotFound:
		return "status404.html"
	case http.StatusConflict:
		return "status409.html"
	case http.StatusLocked:
		return "status423.html"
	default:
		return "status500.html"
	}
}

type Result interface {
	Object() interface{}
	ErrorObject() error
}
