package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorType uint

var errorTypes = []string{"source", "trace"}

const (
	ErrSource ErrorType = iota
	ErrTrace
)

type Error struct {
	Err        bool      `json:"error,required"`
	Code       int       `json:"code,required"`
	Type       ErrorType `json:"type,required"`
	Message    string    `json:"message,required"`
	StackTrace []error   `json:"stack_trace,omitempty"`
}

func (e ErrorType) String() string {
	return errorTypes[e]
}

func Unique(errSlice []error) []error {
	keys := make(map[string]bool)
	var list []error
	for _, entry := range errSlice {
		if _, value := keys[entry.Error()]; !value {
			keys[entry.Error()] = true
			list = append(list, entry)
		}
	}
	return list
}

func New(msg string) Error {
	err := Error{
		Err:     true,
		Code:    1,
		Message: msg,
		Type:    ErrSource,
	}
	return err
}

func Trace(msg string, err error) error {
	if e, ok := err.(*Error); ok {
		return &Error{
			Err:        true,
			Code:       1,
			Message:    msg,
			Type:       ErrSource,
			StackTrace: append(e.StackTrace, err),
		}
	} else {
		e = &Error{
			Err:        true,
			Code:       1,
			Message:    msg,
			Type:       ErrTrace,
			StackTrace: []error{err},
		}
		return e
	}
}

func (e Error) Error() string {
	return fmt.Sprintf("%s error: %d %s %v", e.Type, e.Code, e.Message, e.StackTrace)
}

func (e Error) Timeout() bool {
	return false
}

func (e Error) Temporary() bool {
	return false
}

func HTTP(w http.ResponseWriter, error string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := New(error)
	err.Code = code
	json.NewEncoder(w).Encode(err)
}
