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
}
