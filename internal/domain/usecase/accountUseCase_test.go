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
	"github.com/stretchr/testify/mock"
)

func TestGetAccountOk(t *testing.T) {
	model := model.NewAccountModel("teste", "cpf", "secret", float64(10))

	mockRepository := mocks.NewMockRepository(t)
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.GetAccount), "teste").Return(model, nil)

	usecase := NewUseCase(mockRepository, mockRepository)

	resp, err := usecase.GetAccount("teste")

	assert.Equal(t, httperr.RequestError{}, err)
	assert.Equal(t, resp.ID, model.ID)
}

func TestGetAccountError(t *testing.T) {
	mockRepository := mocks.NewMockRepository(t)
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.GetAccount), "teste").Return(nil, errors.New("teste"))

	usecase := NewUseCase(mockRepository, mockRepository)

	_, err := usecase.GetAccount("teste")

	assert.Equal(t, http.StatusInternalServerError, err.StatusCode)
}

func TestGetAccountNil(t *testing.T) {
	model := &model.AccountModel{}

	mockRepository := mocks.NewMockRepository(t)
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.GetAccount), "teste").Return(model, nil)

	usecase := NewUseCase(mockRepository, mockRepository)

	resp, _ := usecase.GetAccount("teste")

	assert.Empty(t, resp)
}

func TestGetAllAccountOk(t *testing.T) {
	var models []*model.AccountModel
	model := model.NewAccountModel("teste", "cpf", "secret", float64(10))
	models = append(models, model)

	mockRepository := mocks.NewMockRepository(t)
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.GetAllAccount)).Return(models, nil)

	usecase := NewUseCase(mockRepository, mockRepository)

	resp, err := usecase.GetAllAccount()

	assert.Equal(t, httperr.RequestError{}, err)
	assert.Equal(t, len(models), len(resp))
}

func TestGetAllAccountError(t *testing.T) {
	mockRepository := mocks.NewMockRepository(t)
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.GetAllAccount)).Return(nil, errors.New("teste"))

	usecase := NewUseCase(mockRepository, mockRepository)

	_, err := usecase.GetAllAccount()

	assert.Equal(t, http.StatusInternalServerError, err.StatusCode)
}

func TestGetAllAccountNil(t *testing.T) {
	var models []*model.AccountModel

	mockRepository := mocks.NewMockRepository(t)
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.GetAllAccount)).Return(models, nil)

	usecase := NewUseCase(mockRepository, mockRepository)

	resp, _ := usecase.GetAllAccount()

	assert.Empty(t, resp)
}

func TestCreateAccountOk(t *testing.T) {
	model1 := model.NewAccountModel("teste", "00487679059", "secret", float64(10))

	mockRepository := mocks.NewMockRepository(t)
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.CreateAccount), mock.Anything).Return(model1, nil)
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.GetAccountByCpf), "00487679059").Return(&model.AccountModel{}, nil)

	usecase := NewUseCase(mockRepository, mockRepository)

	resp, err := usecase.CreateAccount(model1)

	assert.Equal(t, httperr.RequestError{}, err)
	assert.Equal(t, resp.CPF, model1.CPF)
}

func TestCreateAccountValidateCPF(t *testing.T) {
	model := model.NewAccountModel("teste", "teste", "secret", float64(10))

	mockRepository := mocks.NewMockRepository(t)

	usecase := NewUseCase(mockRepository, mockRepository)

	_, err := usecase.CreateAccount(model)

	assert.Equal(t, http.StatusBadRequest, err.StatusCode)
	assert.Equal(t, "The CPF is invalid.", err.Error.Error())
}

func TestCreateAccountValidateBalance(t *testing.T) {
	model := model.NewAccountModel("teste", "00487679059", "secret", float64(-10))

	mockRepository := mocks.NewMockRepository(t)

	usecase := NewUseCase(mockRepository, mockRepository)

	_, err := usecase.CreateAccount(model)

	assert.Equal(t, http.StatusBadRequest, err.StatusCode)
	assert.Equal(t, "The balance cannot be negative", err.Error.Error())
}

func TestCreateAccountValidateCPFRegistered(t *testing.T) {
	model := model.NewAccountModel("teste", "00487679059", "secret", float64(10))

	mockRepository := mocks.NewMockRepository(t)
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.GetAccountByCpf), "00487679059").Return(model, nil)

	usecase := NewUseCase(mockRepository, mockRepository)

	_, err := usecase.CreateAccount(model)

	assert.Equal(t, http.StatusBadRequest, err.StatusCode)
	assert.Equal(t, "CPF already registered", err.Error.Error())
}

func TestCreateAccountError(t *testing.T) {
	model1 := model.NewAccountModel("teste", "00487679059", "secret", float64(10))

	mockRepository := mocks.NewMockRepository(t)
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.CreateAccount), mock.Anything).Return(nil, errors.New("teste"))
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.GetAccountByCpf), "00487679059").Return(&model.AccountModel{}, nil)

	usecase := NewUseCase(mockRepository, mockRepository)

	_, err := usecase.CreateAccount(model1)

	assert.Equal(t, http.StatusInternalServerError, err.StatusCode)
}
