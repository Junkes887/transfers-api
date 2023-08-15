package web

import "github.com/Junkes887/transfers-api/internal/ports"

type Handler struct {
	AccountUseCase  ports.AccountUseCase
	LoginUseCase    ports.LoginUseCase
	TransferUseCase ports.TransferUseCase
}

func NewHandler(accountUseCase ports.AccountUseCase, loginUseCase ports.LoginUseCase, transferUseCase ports.TransferUseCase) *Handler {
	return &Handler{
		AccountUseCase:  accountUseCase,
		LoginUseCase:    loginUseCase,
		TransferUseCase: transferUseCase,
	}
}
