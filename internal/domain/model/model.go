package model

import (
	"time"

	"github.com/google/uuid"
)

var FORMAT_DATE_TIME = "2006-01-02 15:04:05"

type AccountModel struct {
	ID        string
	Name      string
	CPF       string
	Secret    string
	Balance   float64
	CreatedAt string
}

func NewAccountModel(name string, cpf string, secret string, balance float64) *AccountModel {
	return &AccountModel{
		ID:        uuid.NewString(),
		Name:      name,
		CPF:       cpf,
		Secret:    secret,
		Balance:   balance,
		CreatedAt: time.Now().Format(FORMAT_DATE_TIME),
	}
}

type LoginModel struct {
	CPF    string
	Secret string
}

type TransferModel struct {
	ID                   string
	AccountOriginID      string
	AccountDestinationID string
	Amount               float64
	CreatedAt            string
}

func NewTransferModel(accountOriginID string, accountDestinationID string, amount float64) *TransferModel {
	return &TransferModel{
		ID:                   uuid.NewString(),
		AccountOriginID:      accountOriginID,
		AccountDestinationID: accountDestinationID,
		Amount:               amount,
		CreatedAt:            time.Now().Format(FORMAT_DATE_TIME),
	}
}
