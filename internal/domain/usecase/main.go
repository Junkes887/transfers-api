package usecase

import "github.com/Junkes887/transfers-api/internal/ports"

type UseCase struct {
	AccountRepository ports.AccountRepository
}

func NewUseCase(accountRepository ports.AccountRepository) *UseCase {
	return &UseCase{
		AccountRepository: accountRepository,
	}
}
