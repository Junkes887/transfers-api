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

func TestGetTransferOk(t *testing.T) {
	var list []*model.TransferModel
	transfer := &model.TransferModel{ID: "123"}
	list = append(list, transfer)

	account := &model.AccountModel{ID: "123"}

	mockRepository := mocks.NewMockRepository(t)
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.GetAccountByCpf), "123").Return(account, nil)
	mockRepository.On(testutils.FunctionName(t, ports.TransferRepository.GetTransfer), "123").Return(list, nil)

	usecase := NewUseCase(mockRepository, mockRepository)

	resp, err := usecase.GetTransfer("123")

	assert.Equal(t, httperr.RequestError{}, err)
	assert.Equal(t, len(resp), len(list))
}

func TestGetTransferErrorGetAccountByCpf(t *testing.T) {
	mockRepository := mocks.NewMockRepository(t)
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.GetAccountByCpf), "123").Return(nil, errors.New("error"))

	usecase := NewUseCase(mockRepository, mockRepository)

	_, err := usecase.GetTransfer("123")

	assert.Equal(t, http.StatusInternalServerError, err.StatusCode)
	assert.Equal(t, "error", err.Error.Error())
}

func TestGetTransferAccountNotFound(t *testing.T) {
	var list []*model.TransferModel
	transfer := &model.TransferModel{ID: "123"}
	list = append(list, transfer)

	mockRepository := mocks.NewMockRepository(t)
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.GetAccountByCpf), "123").Return(&model.AccountModel{}, nil)

	usecase := NewUseCase(mockRepository, mockRepository)

	_, err := usecase.GetTransfer("123")

	assert.Equal(t, http.StatusBadRequest, err.StatusCode)
	assert.Equal(t, "Account origin not found", err.Error.Error())
}

func TestGetTransferErrorGetTransfer(t *testing.T) {
	var list []*model.TransferModel
	account := &model.AccountModel{ID: "123"}

	mockRepository := mocks.NewMockRepository(t)
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.GetAccountByCpf), "123").Return(account, nil)
	mockRepository.On(testutils.FunctionName(t, ports.TransferRepository.GetTransfer), "123").Return(list, errors.New("error"))

	usecase := NewUseCase(mockRepository, mockRepository)

	_, err := usecase.GetTransfer("123")

	assert.Equal(t, http.StatusInternalServerError, err.StatusCode)
	assert.Equal(t, "error", err.Error.Error())
}

func TestCreateTransferOk(t *testing.T) {
	transfer := &model.TransferModel{Amount: float64(100), AccountDestinationID: "321"}

	accountOrigin := &model.AccountModel{ID: "123", Balance: float64(1000)}
	accountDestination := &model.AccountModel{ID: "321", Balance: float64(1000)}

	mockRepository := mocks.NewMockRepository(t)
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.GetAccountByCpf), "123").Return(accountOrigin, nil)
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.GetAccount), "321").Return(accountDestination, nil)
	mockRepository.On(testutils.FunctionName(t, ports.TransferRepository.CreateTransfer), mock.Anything).Return(transfer, nil)
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.UpdateAccount), mock.Anything, mock.Anything).Return(nil)

	usecase := NewUseCase(mockRepository, mockRepository)

	resp, err := usecase.CreateTransfer("123", transfer)

	assert.Equal(t, httperr.RequestError{}, err)
	assert.Equal(t, transfer.Amount, resp.Amount)
}

func TestCreateTransferErrorUpdateAccountDest(t *testing.T) {
	transfer := &model.TransferModel{Amount: float64(100), AccountDestinationID: "321"}

	accountOrigin := &model.AccountModel{ID: "123", Balance: float64(1000)}
	accountDestination := &model.AccountModel{ID: "321", Balance: float64(1000)}

	mockRepository := mocks.NewMockRepository(t)
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.GetAccountByCpf), "123").Return(accountOrigin, nil)
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.GetAccount), "321").Return(accountDestination, nil)
	mockRepository.On(testutils.FunctionName(t, ports.TransferRepository.CreateTransfer), mock.Anything).Return(transfer, nil)
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.UpdateAccount), "123", mock.Anything).Return(nil)
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.UpdateAccount), "321", mock.Anything).Return(errors.New("error"))

	usecase := NewUseCase(mockRepository, mockRepository)

	_, err := usecase.CreateTransfer("123", transfer)

	assert.Equal(t, http.StatusInternalServerError, err.StatusCode)
	assert.Equal(t, "error", err.Error.Error())
}

func TestCreateTransferErrorUpdateAccountOri(t *testing.T) {
	transfer := &model.TransferModel{Amount: float64(100), AccountDestinationID: "321"}

	accountOrigin := &model.AccountModel{ID: "123", Balance: float64(1000)}
	accountDestination := &model.AccountModel{ID: "321", Balance: float64(1000)}

	mockRepository := mocks.NewMockRepository(t)
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.GetAccountByCpf), "123").Return(accountOrigin, nil)
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.GetAccount), "321").Return(accountDestination, nil)
	mockRepository.On(testutils.FunctionName(t, ports.TransferRepository.CreateTransfer), mock.Anything).Return(transfer, nil)
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.UpdateAccount), "123", mock.Anything).Return(errors.New("error"))

	usecase := NewUseCase(mockRepository, mockRepository)

	_, err := usecase.CreateTransfer("123", transfer)

	assert.Equal(t, http.StatusInternalServerError, err.StatusCode)
	assert.Equal(t, "error", err.Error.Error())
}

func TestCreateTransferErrorCreateTransfer(t *testing.T) {
	transfer := &model.TransferModel{Amount: float64(100), AccountDestinationID: "321"}

	accountOrigin := &model.AccountModel{ID: "123", Balance: float64(1000)}
	accountDestination := &model.AccountModel{ID: "321", Balance: float64(1000)}

	mockRepository := mocks.NewMockRepository(t)
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.GetAccountByCpf), "123").Return(accountOrigin, nil)
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.GetAccount), "321").Return(accountDestination, nil)
	mockRepository.On(testutils.FunctionName(t, ports.TransferRepository.CreateTransfer), mock.Anything).Return(nil, errors.New("error"))

	usecase := NewUseCase(mockRepository, mockRepository)

	_, err := usecase.CreateTransfer("123", transfer)

	assert.Equal(t, http.StatusInternalServerError, err.StatusCode)
	assert.Equal(t, "error", err.Error.Error())
}

func TestCreateTransferErrorDestinationNF(t *testing.T) {
	transfer := &model.TransferModel{Amount: float64(100), AccountDestinationID: "321"}

	accountOrigin := &model.AccountModel{ID: "123", Balance: float64(1000)}

	mockRepository := mocks.NewMockRepository(t)
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.GetAccountByCpf), "123").Return(accountOrigin, nil)
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.GetAccount), "321").Return(&model.AccountModel{}, nil)

	usecase := NewUseCase(mockRepository, mockRepository)

	_, err := usecase.CreateTransfer("123", transfer)

	assert.Equal(t, http.StatusBadRequest, err.StatusCode)
	assert.Equal(t, "Account destination not found", err.Error.Error())
}

func TestCreateTransferErrorGetAccount(t *testing.T) {
	transfer := &model.TransferModel{Amount: float64(100), AccountDestinationID: "321"}

	accountOrigin := &model.AccountModel{ID: "123", Balance: float64(1000)}

	mockRepository := mocks.NewMockRepository(t)
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.GetAccountByCpf), "123").Return(accountOrigin, nil)
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.GetAccount), "321").Return(nil, errors.New("error"))

	usecase := NewUseCase(mockRepository, mockRepository)

	_, err := usecase.CreateTransfer("123", transfer)

	assert.Equal(t, http.StatusInternalServerError, err.StatusCode)
	assert.Equal(t, "error", err.Error.Error())
}

func TestCreateTransferErrorOriDesEquals(t *testing.T) {
	transfer := &model.TransferModel{Amount: float64(100), AccountDestinationID: "123"}

	accountOrigin := &model.AccountModel{ID: "123", Balance: float64(1000)}

	mockRepository := mocks.NewMockRepository(t)
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.GetAccountByCpf), "123").Return(accountOrigin, nil)

	usecase := NewUseCase(mockRepository, mockRepository)

	_, err := usecase.CreateTransfer("123", transfer)

	assert.Equal(t, http.StatusBadRequest, err.StatusCode)
	assert.Equal(t, "You cannot make a transfer to your account", err.Error.Error())
}

func TestCreateTransferErrorAmount(t *testing.T) {
	transfer := &model.TransferModel{Amount: float64(1000), AccountDestinationID: "321"}

	accountOrigin := &model.AccountModel{ID: "123", Balance: float64(100)}

	mockRepository := mocks.NewMockRepository(t)
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.GetAccountByCpf), "123").Return(accountOrigin, nil)

	usecase := NewUseCase(mockRepository, mockRepository)

	_, err := usecase.CreateTransfer("123", transfer)

	assert.Equal(t, http.StatusBadRequest, err.StatusCode)
	assert.Equal(t, "Source account without balance for transaction", err.Error.Error())
}

func TestCreateTransferErrorOriginNotFound(t *testing.T) {
	transfer := &model.TransferModel{Amount: float64(100), AccountDestinationID: "321"}

	mockRepository := mocks.NewMockRepository(t)
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.GetAccountByCpf), "123").Return(&model.AccountModel{}, nil)

	usecase := NewUseCase(mockRepository, mockRepository)

	_, err := usecase.CreateTransfer("123", transfer)

	assert.Equal(t, http.StatusBadRequest, err.StatusCode)
	assert.Equal(t, "Account origin not found", err.Error.Error())
}

func TestCreateTransferErrorGetAccountByCpf(t *testing.T) {
	transfer := &model.TransferModel{Amount: float64(100), AccountDestinationID: "321"}

	mockRepository := mocks.NewMockRepository(t)
	mockRepository.On(testutils.FunctionName(t, ports.AccountRepository.GetAccountByCpf), "123").Return(nil, errors.New("error"))

	usecase := NewUseCase(mockRepository, mockRepository)

	_, err := usecase.CreateTransfer("123", transfer)

	assert.Equal(t, http.StatusInternalServerError, err.StatusCode)
	assert.Equal(t, "error", err.Error.Error())
}

func TestCreateTransferErrorAmountNegative(t *testing.T) {
	transfer := &model.TransferModel{Amount: float64(-100), AccountDestinationID: "123"}

	mockRepository := mocks.NewMockRepository(t)

	usecase := NewUseCase(mockRepository, mockRepository)

	_, err := usecase.CreateTransfer("123", transfer)

	assert.Equal(t, http.StatusBadRequest, err.StatusCode)
	assert.Equal(t, "The amount for the transfer must be greater than zero", err.Error.Error())
}
