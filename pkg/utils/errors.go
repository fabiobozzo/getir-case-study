package utils

import "errors"

var (
	ErrInvalidInput   = errors.New("the input provided is invalid")
	ErrInternalError  = errors.New("an internal error occurred")
	ErrDatabaseError  = errors.New("a database error occurred")
	ErrEmptyValue     = errors.New("it is forbidden to store empty values")
	ErrNonExistingKey = errors.New("the selected key does not exist")
)

var Errors = map[error]int{
	ErrInvalidInput:   1,
	ErrInternalError:  2,
	ErrDatabaseError:  3,
	ErrEmptyValue:     4,
	ErrNonExistingKey: 5,
}
