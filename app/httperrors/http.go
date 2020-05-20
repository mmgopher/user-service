package httperrors

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// HTTPError represents generic http error response
type HTTPError struct {
	HTTPCode      int    `json:"-"`
	Code          int    `json:"code"`
	Message       string `json:"message"`
	OriginalError error  `json:"-"`
}

// WithCause adds original error object to http error.
func (e *HTTPError) WithCause(err error) *HTTPError {
	e.OriginalError = err
	return e
}

// Error displays code and message as JSON.
func (e *HTTPError) Error() string {
	return fmt.Sprintf(`{"code":%d,"message":"%s"}`, e.Code, e.Message)
}

// New creates new HTTP error with the given HTTP status code.
func New(httpCode int, code int, message string) *HTTPError {
	return &HTTPError{
		HTTPCode:      httpCode,
		Code:          code,
		Message:       message,
		OriginalError: errors.New(message),
	}
}

// NewBadRequest creates new HTTP error with status 400.
func NewBadRequest(code int, message string) *HTTPError {
	return New(http.StatusBadRequest, code, message)
}

// NewNotFound creates new HTTP error with status 404.
func NewNotFound(code int, message string) *HTTPError {
	return New(http.StatusNotFound, code, message)
}

// NewHTTPInternalServerError creates new HTTP error with status 500.
func NewHTTPInternalServerError(code int, message string) *HTTPError {
	return New(http.StatusInternalServerError, code, message)
}

// Emit sets the http error in Gin context and logs the stacktrace
func Emit(ctx *gin.Context, err error) {
	httpError, ok := err.(*HTTPError)
	if !ok {
		httpError = NewHTTPInternalServerError(http.StatusInternalServerError, "")
		httpError.OriginalError = err
	}

	if httpError.OriginalError != nil {
		log.Errorf("%+v", httpError.OriginalError)
	}

	ctx.JSON(httpError.HTTPCode, httpError)
}
