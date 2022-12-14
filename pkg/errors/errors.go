package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	ErrBadRequest       = "Bad request"
	ErrWrongCredentials = "Wrong Credentials"
	ErrNotFound         = "Not Found"
	ErrUnauthorized     = "Unauthorized"
	ErrForbidden        = "Forbidden"
	ErrBadQueryParams   = "Invalid query params"
)

var (
	BadRequest          = errors.New("Bad request")
	NotFound            = errors.New("Not Found")
	Unauthorized        = errors.New("Unauthorized")
	Forbidden           = errors.New("Forbidden")
	PermissionDenied    = errors.New("Permission Denied")
	NotRequiredFields   = errors.New("No such required fields")
	BadQueryParams      = errors.New("Invalid query params")
	InternalServerError = errors.New("Internal Server Error")
	RequestTimeoutError = errors.New("Request Timeout")
	InvalidJWTToken     = errors.New("Invalid JWT token")
	InvalidJWTClaims    = errors.New("Invalid JWT claims")
)

// Error struct
type Error struct {
	ErrCode    int    `json:"code,omitempty"`
	ErrMessage string `json:"message,omitempty"`
	ErrCauses  any    `json:"-"`
}

// Error  Error() interface method
func (e Error) Error() string {
	return fmt.Sprintf("code: %d - message: %s - errors: %v", e.ErrCode, e.ErrMessage, e.ErrCauses)
}

// Error status
func (e Error) Status() int {
	return e.ErrCode
}

// Error Causes
func (e Error) Causes() any {
	return e.ErrCauses
}

// New Error
func NewError(code int, message string, causes any) *Error {
	return &Error{
		ErrCode:    code,
		ErrMessage: message,
		ErrCauses:  causes,
	}
}

// New Error With Message
func NewErrorWithMessage(code int, message string, causes any) *Error {
	return &Error{
		ErrCode:    code,
		ErrMessage: message,
		ErrCauses:  causes,
	}
}

// New Error From Bytes
func NewErrorFromBytes(bytes []byte) (*Error, error) {
	apiErr := &Error{}
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return apiErr, nil
}

// New Bad Request Error
func NewBadRequestError(causes any) *Error {
	return &Error{
		ErrCode:    http.StatusBadRequest,
		ErrMessage: BadRequest.Error(),
		ErrCauses:  causes,
	}
}

// New Not Found Error
func NewNotFoundError(causes any) *Error {
	return &Error{
		ErrCode:    http.StatusNotFound,
		ErrMessage: NotFound.Error(),
		ErrCauses:  causes,
	}
}

// New Unauthorized Error
func NewUnauthorizedError(causes any) *Error {
	return &Error{
		ErrCode:    http.StatusUnauthorized,
		ErrMessage: Unauthorized.Error(),
		ErrCauses:  causes,
	}
}

// New Forbidden Error
func NewForbiddenError(causes any) *Error {
	return &Error{
		ErrCode:    http.StatusForbidden,
		ErrMessage: Forbidden.Error(),
		ErrCauses:  causes,
	}
}

// New Internal Server Error
func NewInternalServerError(causes any) *Error {
	return &Error{
		ErrCode:    http.StatusInternalServerError,
		ErrMessage: InternalServerError.Error(),
		ErrCauses:  causes,
	}
}

type withMessage struct {
	cause error
	msg   string
}

func (w *withMessage) Error() string { return w.msg + ": " + w.cause.Error() }
func (w *withMessage) Cause() error  { return w.cause }

// WithMessage annotates err with a new message.
// If err is nil, WithMessage returns nil.
func WithMessage(err error, message string) error {
	if err == nil {
		return nil
	}
	return &withMessage{
		cause: err,
		msg:   message,
	}
}
