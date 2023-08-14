package entity

import (
	"github.com/Junkes887/transfers-api/internal/domain/model"
	"github.com/Junkes887/transfers-api/pkg/crypt"
)

type AccountEntity struct {
	ID        string
	Name      string
	CPF       string
	Secret    []byte
	Balance   float64
	CreatedAt string
}

func AccountEntityToModelList(entities []*AccountEntity) []*model.AccountModel {
	var models []*model.AccountModel

	for _, entity := range entities {
		models = append(models, AccountEntityToModel(entity))
	}

	return models
}

func AccountEntityToModel(entity *AccountEntity) *model.AccountModel {
	return &model.AccountModel{
		ID:        entity.ID,
		Name:      entity.Name,
		CPF:       entity.CPF,
		Secret:    crypt.Decrypt(entity.Secret),
		Balance:   entity.Balance,
		CreatedAt: entity.CreatedAt,
	}
}

func AccountModelToEntityList(models []*model.AccountModel) []*AccountEntity {
	var entities []*AccountEntity

	for _, model := range models {
		entities = append(entities, AccountModelToEntity(model))
	}

	return entities
}

func AccountModelToEntity(model *model.AccountModel) *AccountEntity {
	return &AccountEntity{
		ID:        model.ID,
		Name:      model.Name,
		CPF:       model.CPF,
		Secret:    crypt.Encrypt(model.Secret),
		Balance:   model.Balance,
		CreatedAt: model.CreatedAt,
	}
}
