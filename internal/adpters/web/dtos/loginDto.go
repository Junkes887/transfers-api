package dtos

import "github.com/Junkes887/transfers-api/internal/domain/model"

type LoginInput struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}

type LoginOutput struct {
	Token string `json:"token"`
}

func LoginInputToLoginModel(dto *LoginInput) *model.LoginModel {
	return &model.LoginModel{
		CPF:    dto.CPF,
		Secret: dto.Secret,
	}
}
