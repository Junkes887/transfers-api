package repository

import (
	"fmt"

	"github.com/Junkes887/transfers-api/internal/adpters/database/entity"
	"github.com/Junkes887/transfers-api/internal/domain/model"
)

func (r *Repository) CreateAccount(model *model.AccountModel) error {
	entity := entity.AccountModelToAccountEntity(model)

	_, err := r.CFG.DB.Exec(
		"INSERT INTO ACCOUNTS (ID, NAME, CPF, SECRET, BALANCE, CREATED_AT) VALUES(?,?,?,?,?,?)",
		entity.ID, entity.Name, entity.CPF, entity.Secret, entity.Balance, entity.CreatedAt)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (r *Repository) GetAllAccount() ([]*model.AccountModel, error) {
	rows, err := r.CFG.DB.Query("SELECT ID, NAME, CPF, SECRET, BALANCE, CREATED_AT FROM ACCOUNTS")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var entities []*entity.AccountEntity
	for rows.Next() {
		var entity entity.AccountEntity
		err = rows.Scan(&entity.ID, &entity.Name, &entity.CPF, &entity.Secret, &entity.Balance, &entity.CreatedAt)
		if err != nil {
			return nil, err
		}

		entities = append(entities, &entity)
	}

	models := entity.AccountEntityToAccountModelList(entities)

	return models, nil
}

func (r *Repository) GetAccount(id string) (*model.AccountModel, error) {
	rows, err := r.CFG.DB.Query("SELECT ID, NAME, CPF, SECRET, BALANCE, CREATED_AT FROM ACCOUNTS WHERE ID = ?", id)
	model := &model.AccountModel{}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		e := &entity.AccountEntity{}
		err = rows.Scan(&e.ID, &e.Name, &e.CPF, &e.Secret, &e.Balance, &e.CreatedAt)
		if err != nil {
			return nil, err
		}

		model = entity.AccountEntityToAccountModel(e)
	}

	return model, nil
}

func (r *Repository) GetAccountByCpf(cpf string) (*model.AccountModel, error) {
	rows, err := r.CFG.DB.Query("SELECT ID, NAME, CPF, SECRET, BALANCE, CREATED_AT FROM ACCOUNTS WHERE CPF = ?", cpf)
	model := &model.AccountModel{}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		e := &entity.AccountEntity{}
		err = rows.Scan(&e.ID, &e.Name, &e.CPF, &e.Secret, &e.Balance, &e.CreatedAt)
		if err != nil {
			return nil, err
		}

		model = entity.AccountEntityToAccountModel(e)
	}

	return model, nil
}

func (r *Repository) UpdateAccount(id string, balance float64) error {
	_, err := r.CFG.DB.Exec("UPDATE ACCOUNTS SET BALANCE = ? WHERE ID = ?", balance, id)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
