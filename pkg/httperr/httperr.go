package httperr

import (
	"encoding/json"
	"errors"
	"log"
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

// func ErrorHttpServerError(requestError RequestError, w http.ResponseWriter) {
// 	message := RequestErrorDto{
// 		Message: requestError.Error.Error(),
// 	}

// 	w.WriteHeader(requestError.StatusCode)
// 	json.NewEncoder(w).Encode(message)
// }

// func ErrorHttpStatusInternalServerError(err error, w http.ResponseWriter) {
// 	message := RequestErrorDto{
// 		Message: err.Error(),
// 	}

// 	w.WriteHeader(http.StatusInternalServerError)
// 	json.NewEncoder(w).Encode(message)

// }

func ErrorHttpServerError(requestError RequestError, w http.ResponseWriter) {
	message := RequestErrorDto{
		Message: requestError.Error.Error(),
	}

	jsonBytes, err := json.Marshal(message)
	if err != nil {
		log.Fatalf("Error marshaling struct: %v", err)
	}

	body := string(jsonBytes)
	http.Error(w, body, requestError.StatusCode)
}

func ErrorHttpStatusInternalServerError(err error, w http.ResponseWriter) {
	message := RequestErrorDto{
		Message: err.Error(),
	}

	jsonBytes, err := json.Marshal(message)
	if err != nil {
		log.Fatalf("Error marshaling struct: %v", err)
	}

	body := string(jsonBytes)
	http.Error(w, body, http.StatusInternalServerError)
}
