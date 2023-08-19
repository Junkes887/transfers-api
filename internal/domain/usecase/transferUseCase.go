package usecase

import (
	"net/http"

	"github.com/Junkes887/transfers-api/internal/domain/model"
	"github.com/Junkes887/transfers-api/internal/ports"
	"github.com/Junkes887/transfers-api/pkg/httperr"
)

func (u *UseCase) CreateTransfer(cpf string, input *model.TransferModel) (*model.TransferModel, httperr.RequestError) {
	accountOrigin, accountDestination, requestError := validateNewTransfer(cpf, input, u.AccountRepository)
	if requestError != (httperr.RequestError{}) {
		return nil, requestError
	}

	transfer := model.NewTransferModel(accountOrigin.ID, input.AccountDestinationID, input.Amount)

	err := u.TransferRepository.CreateTransfer(transfer)
	if err != nil {
		return nil, httperr.NewRequestError(err.Error(), http.StatusInternalServerError)
	}

	accountOrigin.Balance = accountOrigin.Balance - transfer.Amount
	accountDestination.Balance = accountDestination.Balance + transfer.Amount

	err = updateAccount(accountOrigin, u.AccountRepository)
	if err != nil {
		return nil, httperr.NewRequestError(err.Error(), http.StatusInternalServerError)
	}

	err = updateAccount(accountDestination, u.AccountRepository)
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

func validateNewTransfer(cpf string, input *model.TransferModel, rep ports.AccountRepository) (*model.AccountModel, *model.AccountModel, httperr.RequestError) {
	if input.Amount < float64(0) {
		return nil, nil, httperr.NewRequestError("The amount for the transfer must be greater than zero", http.StatusBadRequest)
	}

	accountOrigin, err := rep.GetAccountByCpf(cpf)

	if err != nil {
		return nil, nil, httperr.NewRequestError(err.Error(), http.StatusInternalServerError)
	}

	if *accountOrigin == (model.AccountModel{}) {
		return nil, nil, httperr.NewRequestError("Account origin not found", http.StatusBadRequest)
	}

	if accountOrigin.Balance < input.Amount {
		return nil, nil, httperr.NewRequestError("Source account without balance for transaction", http.StatusBadRequest)
	}

	if accountOrigin.ID == input.AccountDestinationID {
		return nil, nil, httperr.NewRequestError("You cannot make a transfer to your account", http.StatusBadRequest)
	}

	accountDestination, err := rep.GetAccount(input.AccountDestinationID)
	if err != nil {
		return nil, nil, httperr.NewRequestError(err.Error(), http.StatusInternalServerError)
	}
	if *accountDestination == (model.AccountModel{}) {
		return nil, nil, httperr.NewRequestError("Account destination not found", http.StatusBadRequest)
	}

	return accountOrigin, accountDestination, httperr.RequestError{}
}

func updateAccount(account *model.AccountModel, rep ports.AccountRepository) error {
	return rep.UpdateAccount(account.ID, account.Balance)
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
