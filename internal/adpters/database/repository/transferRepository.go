package repository

import (
	"fmt"

	"github.com/Junkes887/transfers-api/internal/adpters/database/entity"
	"github.com/Junkes887/transfers-api/internal/domain/model"
)

func (r *Repository) CreateTransfer(model *model.TransferModel) error {
	entity := entity.TransferModelToTransferEntity(model)

	_, err := r.CFG.DB.Exec(
		"INSERT INTO TRANSFERS (ID, ACCOUNT_ORIGIN_ID, ACCOUNT_DESTINATION_ID, AMOUNT, CREATED_AT) VALUES(?,?,?,?,?)",
		entity.ID, entity.AccountOriginID, entity.AccountDestinationID, entity.Amount, entity.CreatedAt)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (r *Repository) GetTransfer(accountOriginID string) ([]*model.TransferModel, error) {
	rows, err := r.CFG.DB.Query("SELECT ID, ACCOUNT_ORIGIN_ID, ACCOUNT_DESTINATION_ID, AMOUNT, CREATED_AT FROM TRANSFERS WHERE ACCOUNT_ORIGIN_ID = ?", accountOriginID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var entities []*entity.TransferEntity
	for rows.Next() {
		var entity entity.TransferEntity
		err = rows.Scan(&entity.ID, &entity.AccountOriginID, &entity.AccountDestinationID, &entity.Amount, &entity.CreatedAt)
		if err != nil {
			return nil, err
		}

		entities = append(entities, &entity)
	}

	models := entity.TransferEntityToTransferModelList(entities)

	return models, nil
}
