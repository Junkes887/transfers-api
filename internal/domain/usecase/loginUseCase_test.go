package usecase

import (
	"errors"
	"net/http"
	"testing"

	"github.com/Junkes887/transfers-api/internal/domain/model"
	"github.com/Junkes887/transfers-api/internal/ports"
	"github.com/Junkes887/transfers-api/pkg/httperr"
	mocks "github.com/Junkes887/transfers-api/tests/mocks"
	"github.com/sergicanet9/scv-go-tools/v3/testutils"
	"github.com/stretchr/testify/assert"
)

func TestLoginOk(t *testing.T) {
	login := &model.LoginModel{CPF: "65223068084", Secret: "123"}
	account := &model.AccountModel{CPF: "65223068084", Secret: "123"}

	mockRepository := mocks.NewMockRepository(t)
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.GetAccountByCpf), "65223068084").Return(account, nil)

	usecase := NewUseCase(mockRepository, mockRepository)

	token, err := usecase.Login(login)

	assert.Equal(t, httperr.RequestError{}, err)
	assert.NotNil(t, token)

}

func TestLoginValidCpf(t *testing.T) {
	login := &model.LoginModel{CPF: "65223068084", Secret: "123"}
	account := &model.AccountModel{CPF: "65223068084", Secret: "321"}

	mockRepository := mocks.NewMockRepository(t)
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.GetAccountByCpf), "65223068084").Return(account, nil)

	usecase := NewUseCase(mockRepository, mockRepository)

	_, err := usecase.Login(login)

	assert.Equal(t, http.StatusInternalServerError, err.StatusCode)
	assert.Equal(t, "Incorrect CPF or Secret", err.Error.Error())
}

func TestLoginError(t *testing.T) {
	login := &model.LoginModel{CPF: "65223068084", Secret: "123"}

	mockRepository := mocks.NewMockRepository(t)
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.GetAccountByCpf), "65223068084").Return(nil, errors.New("error"))

	usecase := NewUseCase(mockRepository, mockRepository)

	_, err := usecase.Login(login)

	assert.Equal(t, http.StatusInternalServerError, err.StatusCode)
	assert.Equal(t, "error", err.Error.Error())
}
