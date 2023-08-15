package usecase

import (
	"net/http"

	"github.com/Junkes887/transfers-api/internal/domain/model"
	"github.com/Junkes887/transfers-api/internal/ports"
	"github.com/Junkes887/transfers-api/pkg/httperr"
	"github.com/klassmann/cpfcnpj"
)

func (u *UseCase) GetAccount(id string) (*model.AccountModel, httperr.RequestError) {
	requestError := httperr.RequestError{}
	account, err := u.AccountRepository.GetAccount(id)

	if err != nil {
		requestError = httperr.NewRequestError(err.Error(), http.StatusInternalServerError)
	}

	if *account == (model.AccountModel{}) {
		requestError = httperr.NewRequestError("Account not found", http.StatusNotFound)
	}

	return account, requestError
}

func (u *UseCase) GetAllAccount() ([]*model.AccountModel, httperr.RequestError) {
	requestError := httperr.RequestError{}
	list, err := u.AccountRepository.GetAllAccount()

	if err != nil {
		requestError = httperr.NewRequestError(err.Error(), http.StatusInternalServerError)
	}

	return list, requestError
}

func (u *UseCase) CreateAccount(input *model.AccountModel) (*model.AccountModel, httperr.RequestError) {
	requestError := validateAccount(input, u.AccountRepository)

	if requestError != (httperr.RequestError{}) {
		return nil, requestError
	}

	account := model.NewAccountModel(input.Name, input.CPF, input.Secret, input.Balance)

	err := u.AccountRepository.CreateAccount(account)
	if err != nil {
		requestError = httperr.NewRequestError(err.Error(), http.StatusInternalServerError)
	}

	return &model.AccountModel{
		ID:        account.ID,
		Name:      account.Name,
		CPF:       account.CPF,
		Secret:    account.Secret,
		Balance:   account.Balance,
		CreatedAt: account.CreatedAt,
	}, requestError
}

func validateAccount(input *model.AccountModel, accountRepository ports.AccountRepository) httperr.RequestError {
	requestError := httperr.RequestError{}

	if !cpfcnpj.ValidateCPF(input.CPF) {
		requestError = httperr.NewRequestError("The CPF is invalid.", http.StatusBadRequest)
	}

	if input.Balance < float64(0) {
		requestError = httperr.NewRequestError("The balance cannot be negative", http.StatusBadRequest)
	}

	accountCheck, _ := accountRepository.GetAccountByCpf(input.CPF)

	if *accountCheck != (model.AccountModel{}) {
		requestError = httperr.NewRequestError("CPF already registered", http.StatusBadRequest)
	}

	return requestError
}
