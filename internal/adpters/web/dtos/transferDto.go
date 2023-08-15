package dtos

import "github.com/Junkes887/transfers-api/internal/domain/model"

type TransferInput struct {
	AccountDestinationID string  `json:"account_destination_id"`
	Amount               float64 `json:"amount"`
}

type TransferOutput struct {
	ID                   string  `json:"id"`
	AccountOriginID      string  `json:"account_origin_id"`
	AccountDestinationID string  `json:"account_destination_id"`
	Amount               float64 `json:"amount"`
	CreatedAt            string  `json:"created_at"`
}

func TransferInputToTransferModel(dto *TransferInput) *model.TransferModel {
	return &model.TransferModel{
		AccountDestinationID: dto.AccountDestinationID,
		Amount:               dto.Amount,
	}
}

func TransferModelToTransferOutputList(models []*model.TransferModel) []*TransferOutput {
	var list []*TransferOutput

	for _, model := range models {
		list = append(list, TransferModelToTransferOutput(model))
	}

	return list
}

func TransferModelToTransferOutput(model *model.TransferModel) *TransferOutput {
	return &TransferOutput{
		ID:                   model.ID,
		AccountOriginID:      model.AccountOriginID,
		AccountDestinationID: model.AccountDestinationID,
		Amount:               model.Amount,
		CreatedAt:            model.CreatedAt,
	}
}
