package dtos

import "github.com/Junkes887/transfers-api/internal/domain/model"

type AccountInput struct {
	Name      string  `json:"name"`
	CPF       string  `json:"cpf"`
	Secret    string  `json:"secret"`
	Balance   float64 `json:"balance"`
	CreatedAt string  `json:"created_at"`
}

type AccountOutput struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	CPF       string  `json:"cpf"`
	Balance   float64 `json:"balance"`
	CreatedAt string  `json:"created_at"`
}

func AccountInputToAccountModel(dto *AccountInput) *model.AccountModel {
	return &model.AccountModel{
		Name:      dto.Name,
		CPF:       dto.CPF,
		Secret:    dto.Secret,
		Balance:   dto.Balance,
		CreatedAt: dto.CreatedAt,
	}
}

func AccountModelToAccountOutputList(models []*model.AccountModel) []*AccountOutput {
	var dtos []*AccountOutput

	for _, model := range models {
		dtos = append(dtos, AccountModelToAccountOutput(model))
	}

	return dtos
}

func AccountModelToAccountOutput(model *model.AccountModel) *AccountOutput {
	return &AccountOutput{
		ID:        model.ID,
		Name:      model.Name,
		CPF:       model.CPF,
		Balance:   model.Balance,
		CreatedAt: model.CreatedAt,
	}
}

type BalanceOutput struct {
	Balance float64 `json:"balance"`
}

func AccountModelToBalanceOutput(model *model.AccountModel) *BalanceOutput {
	return &BalanceOutput{
		Balance: model.Balance,
	}
}
