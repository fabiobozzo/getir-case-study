package utils

import (
	"errors"
	"net/http"
)

var (
	ErrInvalidInput   = errors.New("the input provided is invalid")
	ErrInternalError  = errors.New("an internal error occurred")
	ErrStorageError   = errors.New("a storage error occurred")
	ErrEmptyValue     = errors.New("it is forbidden to store empty values")
	ErrKeyNotFound    = errors.New("the selected key cannot be found")
	ErrInvalidHeaders = errors.New("request headers are invalid")
)

var ErrorCodeMap = map[error]int{
	ErrInvalidInput:   1,
	ErrInternalError:  2,
	ErrStorageError:   3,
	ErrEmptyValue:     4,
	ErrKeyNotFound:    5,
	ErrInvalidHeaders: 6,
}

var ErrorStatusMap = map[error]int{
	ErrInvalidInput:   http.StatusBadRequest,
	ErrInternalError:  http.StatusInternalServerError,
	ErrStorageError:   http.StatusInternalServerError,
	ErrEmptyValue:     http.StatusBadRequest,
	ErrKeyNotFound:    http.StatusNotFound,
	ErrInvalidHeaders: http.StatusBadRequest,
}
