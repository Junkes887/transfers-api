package usecase

import (
	"net/http"

	"github.com/Junkes887/transfers-api/internal/domain/model"
	"github.com/Junkes887/transfers-api/pkg/httperr"
)

func (u *UseCase) CreateTransfer(cpf string, input *model.TransferModel) (*model.TransferModel, httperr.RequestError) {
	requestError := httperr.RequestError{}

	if input.Amount < float64(0) {
		requestError = httperr.NewRequestError("The amount for the transfer must be greater than zero", http.StatusBadRequest)
		return nil, requestError
	}

	accountOrigin, err := u.AccountRepository.GetAccountByCpf(cpf)

	if err != nil {
		requestError = httperr.NewRequestError(err.Error(), http.StatusInternalServerError)
		return nil, requestError
	}
	if *accountOrigin == (model.AccountModel{}) {
		requestError = httperr.NewRequestError("Account origin not found", http.StatusBadRequest)
		return nil, requestError
	}

	if accountOrigin.Balance < input.Amount {
		requestError = httperr.NewRequestError("Source account without balance", http.StatusBadRequest)
		return nil, requestError
	}

	accountDestination, err := u.AccountRepository.GetAccount(input.AccountDestinationID)
	if err != nil {
		requestError = httperr.NewRequestError(err.Error(), http.StatusInternalServerError)
		return nil, requestError
	}
	if *accountDestination == (model.AccountModel{}) {
		requestError = httperr.NewRequestError("Account destination not found", http.StatusBadRequest)
		return nil, requestError
	}

	transfer := model.NewTransferModel(accountOrigin.ID, input.AccountDestinationID, input.Amount)

	err = u.TransferRepository.CreateTransfer(transfer)
	if err != nil {
		requestError = httperr.NewRequestError(err.Error(), http.StatusInternalServerError)
	}

	accountOrigin.Balance = accountOrigin.Balance - transfer.Amount
	accountDestination.Balance = accountDestination.Balance + transfer.Amount

	err = u.AccountRepository.UpdateAccount(accountOrigin.ID, accountOrigin.Balance)
	if err != nil {
		requestError = httperr.NewRequestError(err.Error(), http.StatusInternalServerError)
	}
	err = u.AccountRepository.UpdateAccount(accountDestination.ID, accountDestination.Balance)
	if err != nil {
		requestError = httperr.NewRequestError(err.Error(), http.StatusInternalServerError)
	}

	return &model.TransferModel{
		ID:                   transfer.ID,
		AccountOriginID:      transfer.AccountOriginID,
		AccountDestinationID: transfer.AccountDestinationID,
		Amount:               transfer.Amount,
		CreatedAt:            transfer.CreatedAt,
	}, requestError
}

func (u *UseCase) GetTransfer(cpf string) ([]*model.TransferModel, httperr.RequestError) {
	requestError := httperr.RequestError{}

	accountOrigin, err := u.AccountRepository.GetAccountByCpf(cpf)

	if err != nil {
		requestError = httperr.NewRequestError(err.Error(), http.StatusInternalServerError)
		return nil, requestError
	}
	if *accountOrigin == (model.AccountModel{}) {
		requestError = httperr.NewRequestError("Account origin not found", http.StatusBadRequest)
		return nil, requestError
	}

	models, err := u.TransferRepository.GetTransfer(accountOrigin.ID)
	if err != nil {
		requestError = httperr.NewRequestError(err.Error(), http.StatusInternalServerError)
	}

	return models, requestError
}
