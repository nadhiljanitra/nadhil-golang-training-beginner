package common

import (
	"encoding/json"
	"errors"
	"net/http"
)

type HTTPErrorResponse struct {
	ErrorCode string `json:"error_code"`
	Message   string `json:"message"`
}

type HTTPError struct {
	StatusCode int
	ErrorCode  string
}

var (
	ErrNotFound    = errors.New("not found")
	ErrIllegalArg  = errors.New("illegal arg")
	ErrUnsupported = errors.New("unsupported")
	ErrPersistence = errors.New("persistence issue")
	ErrUnexpected  = errors.New("unexpected")
)

var errorMapping = map[error]HTTPError{
	ErrNotFound:    NewHTTPErrorResponse(http.StatusNotFound, "DATA_NOT_FOUND_ERROR"),
	ErrIllegalArg:  NewHTTPErrorResponse(http.StatusBadRequest, "API_VALIDATION_ERROR"),
	ErrUnsupported: NewHTTPErrorResponse(http.StatusNotImplemented, "UNIMPLEMENTED"),

	ErrPersistence: NewHTTPErrorResponse(http.StatusInternalServerError, "SERVER_ERROR"),
	ErrUnexpected:  NewHTTPErrorResponse(http.StatusInternalServerError, "SERVER_ERROR"),
}

func ErrorHandler(err error, w http.ResponseWriter, r *http.Request) {
	for mapErr, res := range errorMapping {
		if errors.Is(err, mapErr) {
			returnError(w, r, res.StatusCode, HTTPErrorResponse{
				ErrorCode: res.ErrorCode,
				Message:   err.Error(),
			})
		}
	}
}

func returnError(w http.ResponseWriter, r *http.Request, statusCode int, httpError HTTPErrorResponse) {
	if httpError.Message == "unexpected" {
		httpError.Message = "Something unexpected happened, we are investigating this issue right now"
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(httpError)
}

func NewHTTPErrorResponse(StatusCode int, ErrorCode string) HTTPError {
	return HTTPError{
		StatusCode: StatusCode,
		ErrorCode:  ErrorCode,
	}
}
