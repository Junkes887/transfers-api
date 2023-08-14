package httperr

import (
	"encoding/json"
	"errors"
	"net/http"
)

type RequestErrorDto struct {
	Message string `json:"error_message"`
}

type RequestError struct {
	Error      error
	StatusCode int
}

func NewRequestError(text string, status int) RequestError {
	return RequestError{
		Error:      errors.New(text),
		StatusCode: status,
	}
}

func ErrorHttpStatusInternalServerError(requestError RequestError, w http.ResponseWriter) {
	message := RequestErrorDto{
		Message: requestError.Error.Error(),
	}

	w.WriteHeader(requestError.StatusCode)
	json.NewEncoder(w).Encode(message)
}
