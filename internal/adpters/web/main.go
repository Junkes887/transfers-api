package web

import "github.com/Junkes887/transfers-api/internal/ports"

type Handler struct {
	AccountUseCase ports.AccountUseCase
}

func NewHandler(accountUseCase ports.AccountUseCase) *Handler {
	return &Handler{
		AccountUseCase: accountUseCase,
	}
}
