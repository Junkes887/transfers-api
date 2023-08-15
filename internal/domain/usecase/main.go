package usecase

import "github.com/Junkes887/transfers-api/internal/ports"

type UseCase struct {
	AccountRepository  ports.AccountRepository
	TransferRepository ports.TransferRepository
}

func NewUseCase(accountRepository ports.AccountRepository, transferRepository ports.TransferRepository) *UseCase {
	return &UseCase{
		AccountRepository:  accountRepository,
		TransferRepository: transferRepository,
	}
}
