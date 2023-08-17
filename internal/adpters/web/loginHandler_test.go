package web

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Junkes887/transfers-api/internal/adpters/web/dtos"
	"github.com/Junkes887/transfers-api/internal/ports"
	"github.com/Junkes887/transfers-api/pkg/httperr"
	"github.com/Junkes887/transfers-api/tests/mocks"
	"github.com/sergicanet9/scv-go-tools/v3/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLoginOk(t *testing.T) {
	body := dtos.LoginInput{}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest("POST", "/login", bodyReader)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	mockUsecase := mocks.NewMockUseCase(t)
	mockUsecase.On(testutils.FunctionName(t, ports.LoginUseCase.Login), mock.Anything).Return("teste-token", nil)
	mockHandler := NewHandler(mockUsecase, mockUsecase, mockUsecase)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockHandler.Login)
	handler.ServeHTTP(rr, req)

	var output dtos.LoginOutput

	if err := json.NewDecoder(rr.Body).Decode(&output); err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	assert.Equal(t, rr.Result().StatusCode, http.StatusOK)
	assert.Equal(t, output.Token, "teste-token")
}

func TestLoginErrorDecode(t *testing.T) {
	jsonBody, err := json.Marshal("{'secret': 110.3}")
	if err != nil {
		t.Fatal(err)
	}
	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest("POST", "/login", bodyReader)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	mockUsecase := mocks.NewMockUseCase(t)
	mockHandler := NewHandler(mockUsecase, mockUsecase, mockUsecase)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockHandler.Login)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Result().StatusCode, http.StatusInternalServerError)
}

func TestLoginErrorUseCase(t *testing.T) {
	body := dtos.LoginInput{}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest("POST", "/login", bodyReader)

	if err != nil {
		t.Errorf("Error creating a new request: %v", err)
	}

	mockUsecase := mocks.NewMockUseCase(t)
	mockUsecase.On(testutils.FunctionName(t, ports.LoginUseCase.Login), mock.Anything).Return(nil, httperr.RequestError{StatusCode: 500, Error: errors.New("error")})
	mockHandler := NewHandler(mockUsecase, mockUsecase, mockUsecase)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(mockHandler.Login)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Result().StatusCode, http.StatusInternalServerError)
}
