package usecase

import (
	"net/http"

	"github.com/Junkes887/transfers-api/internal/domain/model"
	"github.com/Junkes887/transfers-api/pkg/httperr"
	"github.com/Junkes887/transfers-api/pkg/jwtToken"
)

func (u *UseCase) Login(login *model.LoginModel) (string, httperr.RequestError) {
	accout, err := u.AccountRepository.GetAccountByCpf(login.CPF)

	if err != nil {
		return "", httperr.NewRequestError(err.Error(), http.StatusInternalServerError)
	}

	if accout == (&model.AccountModel{}) || accout.Secret != login.Secret {
		return "", httperr.NewRequestError("Incorrect CPF or Secret", http.StatusInternalServerError)
	}

	token, err := jwtToken.GenerateJWT(login.CPF, login.Secret)

	if err != nil {
		return "", httperr.NewRequestError(err.Error(), http.StatusInternalServerError)
	}

	return token, httperr.RequestError{}
}
