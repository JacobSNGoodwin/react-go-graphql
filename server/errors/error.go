package errors

import "fmt"

// Type holds a type string and integer code for the error
type Type string

// Model error types after apollo server for now
const (
	Authentication Type = "UNAUTHENTICATED" // Authentication Failures
	Forbidden      Type = "FORBIDDEN"       // Authorization failutre
	Input          Type = "INPUT"           // Validation errors (already exists, not found)
	Internal       Type = "INTERNAL"        // Fallback for uncaught failures
)

// Error holds custom error types/codes using graphql-go error extensions interface
type Error struct {
	Message  string
	Type     Type
	original error
}

// Error satisfies standard error interface
func (e *Error) Error() string {
	return fmt.Sprintf("%v", e.Message)
}

// Cause of the original error
func (e *Error) Cause() string {
	if e.original != nil {
		return e.original.Error()
	}

	return ""
}

// Extensions for graphQL extension allows us to return error type
func (e *Error) Extensions() map[string]interface{} {
	return map[string]interface{}{
		"type": e.Type,
	}
}

// NewAuthentication tells the user that they are not able to access any
// resource which requires a user. Message provides any extra description
func NewAuthentication(msg string, err error) *Error {
	return &Error{
		Message:  msg,
		Type:     Authentication,
		original: err,
	}
}

// NewForbidden provides an error for a user who is unauthenticated
func NewForbidden(msg string, err error) *Error {
	return &Error{
		Message:  msg,
		Type:     Forbidden,
		original: err,
	}
}

// NewInput can be used when input cannot be parsed or is not valid
func NewInput(msg string, err error) *Error {
	return &Error{
		Message:  msg,
		Type:     Input,
		original: err,
	}
}

// NewInternal can be used for data access or other server errors
func NewInternal(msg string, err error) *Error {
	return &Error{
		Message:  msg,
		Type:     Internal,
		original: err,
	}
}
