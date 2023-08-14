package web

import "github.com/Junkes887/transfers-api/internal/ports"

type Handler struct {
	AccountUseCase ports.AccountUseCase
	LoginUseCase   ports.LoginUseCase
}

func NewHandler(accountUseCase ports.AccountUseCase, loginUseCase ports.LoginUseCase) *Handler {
	return &Handler{
		AccountUseCase: accountUseCase,
		LoginUseCase:   loginUseCase,
	}
}
