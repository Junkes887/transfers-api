package web

import (
	"encoding/json"
	"net/http"

	"github.com/Junkes887/transfers-api/internal/adpters/web/dtos"
	"github.com/Junkes887/transfers-api/pkg/httperr"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var input *dtos.LoginInput
	requestError := httperr.RequestError{}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		httperr.ErrorHttpStatusInternalServerError(err, w)
		return
	}

	model := dtos.LoginInputToLoginModel(input)
	token, requestError := h.LoginUseCase.Login(model)
	if requestError != (httperr.RequestError{}) {
		httperr.ErrorHttpServerError(requestError, w)
		return
	}

	output := dtos.LoginOutput{
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
