package errorwrapper

import (
	"log"
	"net/http"
	"runtime"
)

const (
	ErrResourceAlreadyExists = iota + 1
	ErrInternalServer
	ErrBadRequest
	ErrResourceNotFound
	ErrForbidden
)

var (
	httpCodeMap = map[int]int{
		ErrResourceAlreadyExists: http.StatusConflict,
		ErrForbidden:             http.StatusForbidden,
		ErrBadRequest:            http.StatusBadRequest,
		ErrResourceNotFound:      http.StatusNotFound,
		ErrInternalServer:        http.StatusInternalServerError,
	}
)

type ErrorWrapper interface {
	error
}

type errorWrapper struct {
	Code       int
	Message    string
	DevMessage string
}

func (e *errorWrapper) Error() string {
	return e.Message
}

func WrapErr(code int, message string) error {
	// Get the name of the calling function.
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		// Handle error.
		return &errorWrapper{
			Code:       code,
			Message:    message,
			DevMessage: "[Unhandled]",
		}
	}

	// Get the function object.
	fn := runtime.FuncForPC(pc)

	// Get the function name.
	name := fn.Name()

	return &errorWrapper{
		Code:       code,
		Message:    message,
		DevMessage: name,
	}
}

func ConvertToHTTPError(err error) (code int, response map[string]any) {
	if err == nil {
		return http.StatusOK, map[string]any{"message": "ok"}
	}
	ew, ok := err.(*errorWrapper)
	if !ok {
		return http.StatusInternalServerError, map[string]any{"error": err.Error()}
	}

	// log internal server errors
	if ew.Code == ErrInternalServer {
		log.Println(ew.DevMessage+" -> ", ew.Message)
	}

	return httpCodeMap[ew.Code], map[string]any{"error": ew.Message}
}

func IsErrorContainingCode(err error, code int) bool {
	if err == nil {
		return false
	}
	ew, ok := err.(*errorWrapper)
	if !ok {
		return false
	}
	return ew.Code == code
}
