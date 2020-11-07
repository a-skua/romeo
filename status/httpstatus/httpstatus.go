package httpstatus

import (
	"net/http"

	"github.com/a-skua/romeo/status"
)

const (
	StatusOK = iota
	StatusBadRequest
	StatusUnAuthorized
	StatusForbidden
	StatusNotFound
	StatusTeapot
	StatusInternalError
)

// OK return a httstatus 200
func OK() status.HttpStatus {
	return status.NewHttpStatus(
		StatusOK,
		http.StatusOK,
	)
}

// BadRequest return a httpstatus 400
func BadRequest() status.HttpStatus {
	return status.NewHttpStatus(
		StatusBadRequest,
		http.StatusBadRequest,
	)
}

// Unauthorized return a httpstatus 401
func UnAuthorized() status.HttpStatus {
	return status.NewHttpStatus(
		StatusUnAuthorized,
		http.StatusUnauthorized,
	)
}

// Forbidden return a httpstatus 403
func Forbidden() status.HttpStatus {
	return status.NewHttpStatus(
		StatusForbidden,
		http.StatusForbidden,
	)
}

// NotFound return a httpstatus 404
func NotFound() status.HttpStatus {
	return status.NewHttpStatus(
		StatusNotFound,
		http.StatusNotFound,
	)
}

// Teapot return a httpstatus 418
func Teapot() status.HttpStatus {
	return status.NewHttpStatus(
		StatusTeapot,
		http.StatusTeapot,
	)
}

// InternalError return a httpstatus 500
func InternalError() status.HttpStatus {
	return status.NewHttpStatus(
		StatusInternalError,
		http.StatusInternalServerError,
	)
}
