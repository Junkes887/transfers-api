package entity

import (
	"github.com/Junkes887/transfers-api/internal/domain/model"
)

type TransferEntity struct {
	ID                   string
	AccountOriginID      string
	AccountDestinationID string
	Amount               float64
	CreatedAt            string
}

func TransferModelToTransferEntity(model *model.TransferModel) *TransferEntity {
	return &TransferEntity{
		ID:                   model.ID,
		AccountOriginID:      model.AccountOriginID,
		AccountDestinationID: model.AccountDestinationID,
		Amount:               model.Amount,
		CreatedAt:            model.CreatedAt,
	}
}

func TransferEntityToTransferModelList(entities []*TransferEntity) []*model.TransferModel {
	var models []*model.TransferModel

	for _, entity := range entities {
		models = append(models, TransferEntityToTransferModel(entity))
	}

	return models
}

func TransferEntityToTransferModel(entity *TransferEntity) *model.TransferModel {
	return &model.TransferModel{
		ID:                   entity.ID,
		AccountOriginID:      entity.AccountOriginID,
		AccountDestinationID: entity.AccountDestinationID,
		Amount:               entity.Amount,
		CreatedAt:            entity.CreatedAt,
	}
}
