package errors

import (
	"errors"
)

var (
	ErrNotFound   = errors.New("not found")
	ErrBadRequest = errors.New("bad request")
	ErrInternal   = errors.New("internal server error")

	ErrPending = errors.New("transaction in pending status")
)
