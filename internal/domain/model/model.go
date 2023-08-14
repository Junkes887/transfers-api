package model

import (
	"time"

	"github.com/google/uuid"
)

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
		ID:        uuid.New().String(),
		Name:      name,
		CPF:       cpf,
		Secret:    secret,
		Balance:   balance,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
}

type LoginModel struct {
	CPF    string
	Secret string
}
