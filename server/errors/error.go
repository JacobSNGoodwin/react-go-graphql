package errors

import "fmt"

const (
	authenticationError = "UNAUTHENTICATED"
	forbiddenError      = "FORBIDDEN"
	inputError          = "INPUT"
	serverError         = "SERVER"
)

// Error holds custom error types for good error communication with client application
type Error struct {
	Message string
	Code    string
}

func (e Error) Error() string {
	return fmt.Sprintf("%v", e.Message)
}

// NewAuthenticationError tells the user that they are not able to access any
// resource which requires a user. Message provides any extra description
func NewAuthenticationError(msg string) Error {
	return Error{
		Message: msg,
		Code:    authenticationError,
	}
}

// NewForbiddenError provides an error for a user who is unauthenticated
func NewForbiddenError(msg string) Error {
	return Error{
		Message: msg,
		Code:    forbiddenError,
	}
}

// NewInputError can be used when input cannot be parsed or is not valid
func NewInputError(msg string) Error {
	return Error{
		Message: msg,
		Code:    inputError,
	}
}

// NewServerError can be used for data access or other server errors
func NewServerError(msg string) Error {
	return Error{
		Message: msg,
		Code:    serverError,
	}
}
