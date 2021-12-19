package utils

import "errors"

var (
	ErrInvalidInput  = errors.New("the input provided is invalid")
	ErrInternalError = errors.New("an internal error occurred")
	ErrStorageError  = errors.New("a storage error occurred")
	ErrEmptyValue    = errors.New("it is forbidden to store empty values")
	ErrKeyNotFound   = errors.New("the selected key cannot be found")
)

var Errors = map[error]int{
	ErrInvalidInput:  1,
	ErrInternalError: 2,
	ErrStorageError:  3,
	ErrEmptyValue:    4,
	ErrKeyNotFound:   5,
}
