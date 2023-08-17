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

func TestCreateAccountOk(t *testing.T) {
	body := dtos.AccountInput{}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest("POST", "/accounts", bodyReader)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	account := &model.AccountModel{Name: "teste"}
	mockUsecase := mocks.NewMockUseCase(t)
	mockUsecase.On(testutils.FunctionName(t, ports.AccountUseCase.CreateAccount), mock.Anything).Return(account, nil)
	mockHandler := NewHandler(mockUsecase, mockUsecase, mockUsecase)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockHandler.CreateAccount)
	handler.ServeHTTP(rr, req)

	var output dtos.AccountOutput

	if err := json.NewDecoder(rr.Body).Decode(&output); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	assert.Equal(t, rr.Result().StatusCode, http.StatusOK)
	assert.Equal(t, output.Name, account.Name)
}

func TestCreateAccountErrorDecode(t *testing.T) {
	jsonBody, err := json.Marshal("{'secret': 110.3}")
	if err != nil {
		t.Fatal(err)
	}
	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest("POST", "/accounts", bodyReader)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	mockUsecase := mocks.NewMockUseCase(t)
	mockHandler := NewHandler(mockUsecase, mockUsecase, mockUsecase)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockHandler.CreateAccount)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Result().StatusCode, http.StatusInternalServerError)
}

func TestCreateAccountErrorUseCase(t *testing.T) {
	body := dtos.AccountInput{}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest("POST", "/accounts", bodyReader)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	mockUsecase := mocks.NewMockUseCase(t)
	mockUsecase.On(testutils.FunctionName(t, ports.AccountUseCase.CreateAccount), mock.Anything).Return(nil, httperr.RequestError{StatusCode: 500, Error: errors.New("error")})
	mockHandler := NewHandler(mockUsecase, mockUsecase, mockUsecase)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockHandler.CreateAccount)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Result().StatusCode, http.StatusInternalServerError)
}

func TestGetAllAccountOk(t *testing.T) {
	req, err := http.NewRequest("GET", "/accounts", nil)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	var list []*model.AccountModel
	account := &model.AccountModel{Name: "teste"}
	list = append(list, account)

	mockUsecase := mocks.NewMockUseCase(t)
	mockUsecase.On(testutils.FunctionName(t, ports.AccountUseCase.GetAllAccount)).Return(list, nil)
	mockHandler := NewHandler(mockUsecase, mockUsecase, mockUsecase)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockHandler.GetAllAccount)
	handler.ServeHTTP(rr, req)

	var output []dtos.AccountOutput

	if err := json.NewDecoder(rr.Body).Decode(&output); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	assert.Equal(t, rr.Result().StatusCode, http.StatusOK)
	assert.Equal(t, len(output), len(list))
}

func TestGetAllAccountErrorUseCase(t *testing.T) {
	req, err := http.NewRequest("GET", "/accounts", nil)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	mockUsecase := mocks.NewMockUseCase(t)
	mockUsecase.On(testutils.FunctionName(t, ports.AccountUseCase.GetAllAccount), mock.Anything).Return(nil, httperr.RequestError{StatusCode: 500, Error: errors.New("error")})
	mockHandler := NewHandler(mockUsecase, mockUsecase, mockUsecase)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockHandler.GetAllAccount)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Result().StatusCode, http.StatusInternalServerError)
}

func TestGetBalanceOk(t *testing.T) {
	req, err := http.NewRequest("GET", "/accounts/1/balance", nil)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	account := &model.AccountModel{Balance: float64(10)}

	mockUsecase := mocks.NewMockUseCase(t)
	mockUsecase.On(testutils.FunctionName(t, ports.AccountUseCase.GetAccount), mock.Anything).Return(account, nil)
	mockHandler := NewHandler(mockUsecase, mockUsecase, mockUsecase)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockHandler.GetBalance)
	handler.ServeHTTP(rr, req)

	var output dtos.BalanceOutput

	if err := json.NewDecoder(rr.Body).Decode(&output); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	assert.Equal(t, rr.Result().StatusCode, http.StatusOK)
	assert.Equal(t, output.Balance, account.Balance)
}

func TestGetBalanceErrorUseCase(t *testing.T) {
	body := dtos.AccountInput{}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest("GET", "/accounts/1/balance", bodyReader)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	mockUsecase := mocks.NewMockUseCase(t)
	mockUsecase.On(testutils.FunctionName(t, ports.AccountUseCase.GetAccount), mock.Anything).Return(nil, httperr.RequestError{StatusCode: 500, Error: errors.New("error")})
	mockHandler := NewHandler(mockUsecase, mockUsecase, mockUsecase)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockHandler.GetBalance)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Result().StatusCode, http.StatusInternalServerError)
}
