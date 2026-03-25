package errors

import "fmt"

var (
	WRONG_PASSWORD Error = Error{
		Code:    "401",
		Message: "WRONG_PASSWORD",
	}
	USER_NOT_FOUND = Error{
		Code:    "404",
		Message: "USER_NOT_FOUND",
	}
)

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Cause   error  `json:"cause"`
}

func (e *Error) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %s (cause: %v)", e.Code, e.Message, e.Cause)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func (e *Error) Unwrap() error {
	return e.Cause
}
