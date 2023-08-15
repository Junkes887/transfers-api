package ports

import (
	"github.com/Junkes887/transfers-api/internal/domain/model"
	"github.com/Junkes887/transfers-api/pkg/httperr"
)

type AccountUseCase interface {
	GetAccount(id string) (*model.AccountModel, httperr.RequestError)
	GetAllAccount() ([]*model.AccountModel, httperr.RequestError)
	CreateAccount(*model.AccountModel) (*model.AccountModel, httperr.RequestError)
}

type AccountRepository interface {
	GetAccount(id string) (*model.AccountModel, error)
	GetAccountByCpf(cpf string) (*model.AccountModel, error)
	GetAllAccount() ([]*model.AccountModel, error)
	CreateAccount(*model.AccountModel) error
	UpdateAccount(id string, balance float64) error
}

type LoginUseCase interface {
	Login(model *model.LoginModel) (string, httperr.RequestError)
}

type TransferUseCase interface {
	CreateTransfer(accountOriginID string, model *model.TransferModel) (*model.TransferModel, httperr.RequestError)
	GetTransfer(cpf string) ([]*model.TransferModel, httperr.RequestError)
}

type TransferRepository interface {
	CreateTransfer(*model.TransferModel) error
	GetTransfer(accountOriginID string) ([]*model.TransferModel, error)
}
