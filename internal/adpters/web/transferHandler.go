package web

import (
	"encoding/json"
	"net/http"

	"github.com/Junkes887/transfers-api/internal/adpters/web/dtos"
	"github.com/Junkes887/transfers-api/pkg/httperr"
)

func (h *Handler) CreateTransfer(w http.ResponseWriter, r *http.Request) {
	var input *dtos.TransferInput
	requestError := httperr.RequestError{}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		httperr.ErrorHttpStatusInternalServerError(err, w)
		return
	}

	cpf := r.Header.Get("cpf")

	model := dtos.TransferInputToTransferModel(input)
	model, requestError = h.TransferUseCase.CreateTransfer(cpf, model)
	if requestError != (httperr.RequestError{}) {
		httperr.ErrorHttpServerError(requestError, w)
		return
	}

	output := dtos.TransferModelToTransferOutput(model)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)

}

func (h *Handler) GetTransfer(w http.ResponseWriter, r *http.Request) {
	requestError := httperr.RequestError{}
	cpf := r.Header.Get("cpf")

	models, requestError := h.TransferUseCase.GetTransfer(cpf)
	if requestError != (httperr.RequestError{}) {
		httperr.ErrorHttpServerError(requestError, w)
		return
	}

	output := dtos.TransferModelToTransferOutputList(models)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)

}
