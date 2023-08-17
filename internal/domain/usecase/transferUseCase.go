package usecase

import (
	"net/http"

	"github.com/Junkes887/transfers-api/internal/domain/model"
	"github.com/Junkes887/transfers-api/pkg/httperr"
)

func (u *UseCase) CreateTransfer(cpf string, input *model.TransferModel) (*model.TransferModel, httperr.RequestError) {
	if input.Amount < float64(0) {
		return nil, httperr.NewRequestError("The amount for the transfer must be greater than zero", http.StatusBadRequest)
	}

	accountOrigin, err := u.AccountRepository.GetAccountByCpf(cpf)

	if err != nil {
		return nil, httperr.NewRequestError(err.Error(), http.StatusInternalServerError)
	}
	if *accountOrigin == (model.AccountModel{}) {
		return nil, httperr.NewRequestError("Account origin not found", http.StatusBadRequest)
	}

	if accountOrigin.Balance < input.Amount {
		return nil, httperr.NewRequestError("Source account without balance for transaction", http.StatusBadRequest)
	}

	if accountOrigin.ID == input.AccountDestinationID {
		return nil, httperr.NewRequestError("You cannot make a transfer to your account", http.StatusBadRequest)
	}

	accountDestination, err := u.AccountRepository.GetAccount(input.AccountDestinationID)
	if err != nil {
		return nil, httperr.NewRequestError(err.Error(), http.StatusInternalServerError)
	}
	if *accountDestination == (model.AccountModel{}) {
		return nil, httperr.NewRequestError("Account destination not found", http.StatusBadRequest)
	}

	transfer := model.NewTransferModel(accountOrigin.ID, input.AccountDestinationID, input.Amount)

	err = u.TransferRepository.CreateTransfer(transfer)
	if err != nil {
		return nil, httperr.NewRequestError(err.Error(), http.StatusInternalServerError)
	}

	accountOrigin.Balance = accountOrigin.Balance - transfer.Amount
	accountDestination.Balance = accountDestination.Balance + transfer.Amount

	err = u.AccountRepository.UpdateAccount(accountOrigin.ID, accountOrigin.Balance)
	if err != nil {
		return nil, httperr.NewRequestError(err.Error(), http.StatusInternalServerError)
	}
	err = u.AccountRepository.UpdateAccount(accountDestination.ID, accountDestination.Balance)
	if err != nil {
		return nil, httperr.NewRequestError(err.Error(), http.StatusInternalServerError)
	}

	return &model.TransferModel{
		ID:                   transfer.ID,
		AccountOriginID:      transfer.AccountOriginID,
		AccountDestinationID: transfer.AccountDestinationID,
		Amount:               transfer.Amount,
		CreatedAt:            transfer.CreatedAt,
	}, httperr.RequestError{}
}

func (u *UseCase) GetTransfer(cpf string) ([]*model.TransferModel, httperr.RequestError) {
	accountOrigin, err := u.AccountRepository.GetAccountByCpf(cpf)

	if err != nil {
		return nil, httperr.NewRequestError(err.Error(), http.StatusInternalServerError)
	}
	if *accountOrigin == (model.AccountModel{}) {
		return nil, httperr.NewRequestError("Account origin not found", http.StatusBadRequest)
	}

	models, err := u.TransferRepository.GetTransfer(accountOrigin.ID)
	if err != nil {
		return nil, httperr.NewRequestError(err.Error(), http.StatusInternalServerError)
	}

	return models, httperr.RequestError{}
}
