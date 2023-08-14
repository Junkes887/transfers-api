package web

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Junkes887/transfers-api/internal/adpters/web/dtos"
	"github.com/Junkes887/transfers-api/pkg/httperr"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var input *dtos.LoginDtoInput
	requestError := httperr.RequestError{}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		requestError = httperr.RequestError{
			Error:      errors.New(err.Error()),
			StatusCode: http.StatusInternalServerError,
		}
		httperr.ErrorHttpStatusInternalServerError(requestError, w)
		return
	}

	model := dtos.LoginDtoInputToLoginModel(input)
	token, requestError := h.LoginUseCase.Login(model)
	if requestError != (httperr.RequestError{}) {
		httperr.ErrorHttpStatusInternalServerError(requestError, w)
		return
	}

	output := dtos.LoginDtoOutput{
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
