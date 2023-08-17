package web

import (
	"encoding/json"
	"net/http"

	"github.com/Junkes887/transfers-api/internal/adpters/web/dtos"
	"github.com/Junkes887/transfers-api/pkg/httperr"
	"github.com/go-chi/chi"
)

func (h *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var input *dtos.AccountInput
	requestError := httperr.RequestError{}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		httperr.ErrorHttpStatusInternalServerError(err, w)
		return
	}
	model := dtos.AccountInputToAccountModel(input)
	model, requestError = h.AccountUseCase.CreateAccount(model)
	if requestError != (httperr.RequestError{}) {
		httperr.ErrorHttpServerError(requestError, w)
		return
	}

	output := dtos.AccountModelToAccountOutput(model)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)

}

func (h *Handler) GetAllAccount(w http.ResponseWriter, r *http.Request) {
	models, requestError := h.AccountUseCase.GetAllAccount()

	outputs := dtos.AccountModelToAccountOutputList(models)
	if requestError != (httperr.RequestError{}) {
		httperr.ErrorHttpServerError(requestError, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(outputs)
}

func (h *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "account_id")

	model, requestError := h.AccountUseCase.GetAccount(id)
	if requestError != (httperr.RequestError{}) {
		httperr.ErrorHttpServerError(requestError, w)
		return
	}
	output := dtos.AccountModelToBalanceOutput(model)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
