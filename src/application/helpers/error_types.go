package helpers

import (
	"errors"
	"fmt"
)

const (
	InternalError     = "InternalError"
	BadRequestError   = "BadRequest"
	CompareCryptError = "CompareCryptError"
	ValidationError   = "ValidationError"
)

func ErrorMessage(msg string, value interface{}) error {
	errMsg := fmt.Sprintf(msg, value)
	return errors.New(errMsg)
}
