package common

import (
	"net/http"
	"time"
)

// Error is a struct that represents an error
type Error struct {
	Code    int    		`json:"code"`
	Message string 		`json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Path string 		`json:"path"`
}

// NewError creates a new Error
func NewError(code int, message string, path string) *Error {
	return &Error{Code: code, Message: message, Timestamp: time.Now(), Path: path}
}

// NewErrorFromError creates a new Error from an error
func NewErrorFromError(err error, path string) *Error {
	return &Error{Code: http.StatusInternalServerError, Message: err.Error(), Timestamp: time.Now(), Path: path}
}