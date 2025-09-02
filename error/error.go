package errors

import "errors"

var (
	ErrInvalidAuth     = errors.New("login require")
	ErrNotFound        = errors.New("not found")
	ErrTooManyRequests = errors.New("too many requests")
	ErrInternalAPICall = errors.New("internal server error with api call")
	ErrBadRequest      = errors.New("bad request")
)
