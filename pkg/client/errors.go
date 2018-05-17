package client

import "fmt"

type baseError struct {
	code    int
	message string
}

func (b baseError) Error() string {
	return fmt.Sprintf("[Code: %d] %s", b.code, b.message)
}

func (b baseError) Message() string {
	return b.message
}

func (b baseError) Code() int {
	return b.code
}

type ErrorServerUnknownError struct {
	baseError
}

type ErrorBadRequest struct {
	baseError
}

type ErrorAlreadyExists struct {
	baseError
}

type ErrorNotFound struct {
	baseError
}

type ErrorForbidden struct {
	baseError
}
