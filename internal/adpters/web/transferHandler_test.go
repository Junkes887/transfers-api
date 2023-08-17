package web

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Junkes887/transfers-api/internal/adpters/web/dtos"
	"github.com/Junkes887/transfers-api/internal/domain/model"
	"github.com/Junkes887/transfers-api/internal/ports"
	"github.com/Junkes887/transfers-api/pkg/httperr"
	"github.com/Junkes887/transfers-api/tests/mocks"
	"github.com/sergicanet9/scv-go-tools/v3/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTransferOk(t *testing.T) {
	body := dtos.AccountInput{}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest("POST", "/tranfers", bodyReader)
	req.Header.Add("cpf", "123")

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	transfer := &model.TransferModel{Amount: float64(10)}
	mockUsecase := mocks.NewMockUseCase(t)
	mockUsecase.On(testutils.FunctionName(t, ports.TransferUseCase.CreateTransfer), mock.Anything, mock.Anything).Return(transfer, nil)
	mockHandler := NewHandler(mockUsecase, mockUsecase, mockUsecase)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockHandler.CreateTransfer)
	handler.ServeHTTP(rr, req)

	var output dtos.TransferOutput

	if err := json.NewDecoder(rr.Body).Decode(&output); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	assert.Equal(t, rr.Result().StatusCode, http.StatusOK)
	assert.Equal(t, output.Amount, transfer.Amount)
}

func TestCreateTransferErrorDecode(t *testing.T) {
	jsonBody, err := json.Marshal("{'id': 110.3}")
	if err != nil {
		t.Fatal(err)
	}
	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest("POST", "/tranfers", bodyReader)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	mockUsecase := mocks.NewMockUseCase(t)
	mockHandler := NewHandler(mockUsecase, mockUsecase, mockUsecase)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockHandler.CreateTransfer)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Result().StatusCode, http.StatusInternalServerError)
}

func TestCreateTransferErrorUseCase(t *testing.T) {
	body := dtos.AccountInput{}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest("POST", "/tranfers", bodyReader)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	mockUsecase := mocks.NewMockUseCase(t)
	mockUsecase.On(testutils.FunctionName(t, ports.TransferUseCase.CreateTransfer), mock.Anything, mock.Anything).Return(nil, httperr.RequestError{StatusCode: 500, Error: errors.New("error")})
	mockHandler := NewHandler(mockUsecase, mockUsecase, mockUsecase)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockHandler.CreateTransfer)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Result().StatusCode, http.StatusInternalServerError)
}

func TestGetTransferOk(t *testing.T) {
	req, err := http.NewRequest("GET", "/tranfers", nil)
	req.Header.Add("cpf", "123")

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	var list []*model.TransferModel
	transfer := &model.TransferModel{Amount: float64(10)}
	list = append(list, transfer)
	mockUsecase := mocks.NewMockUseCase(t)
	mockUsecase.On(testutils.FunctionName(t, ports.TransferUseCase.GetTransfer), mock.Anything).Return(list, nil)
	mockHandler := NewHandler(mockUsecase, mockUsecase, mockUsecase)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockHandler.GetTransfer)
	handler.ServeHTTP(rr, req)

	var output []dtos.TransferOutput

	if err := json.NewDecoder(rr.Body).Decode(&output); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	assert.Equal(t, rr.Result().StatusCode, http.StatusOK)
	assert.Equal(t, len(output), len(list))
}

func TestGetTransferErrorUseCase(t *testing.T) {
	req, err := http.NewRequest("GET", "/tranfers", nil)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	mockUsecase := mocks.NewMockUseCase(t)
	mockUsecase.On(testutils.FunctionName(t, ports.TransferUseCase.GetTransfer), mock.Anything).Return(nil, httperr.RequestError{StatusCode: 500, Error: errors.New("error")})
	mockHandler := NewHandler(mockUsecase, mockUsecase, mockUsecase)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockHandler.GetTransfer)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Result().StatusCode, http.StatusInternalServerError)
}
