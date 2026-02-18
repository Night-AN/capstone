package errors

import "fmt"

var ()

type DomainError struct {
	Code    string
	Message string
	Cause   error
}

func (e *DomainError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %s (cause: %v)", e.Code, e.Message, e.Cause)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func (e *DomainError) Unwrap() error {
	return e.Cause
}

func NewDomainWithError(code, message string, cause error) DomainError {
	return DomainError{
		Code:    code,
		Message: message,
		Cause:   cause,
	}
}

func NewDomainError(code, message string) DomainError {
	return DomainError{
		Code:    code,
		Message: message,
	}
}
