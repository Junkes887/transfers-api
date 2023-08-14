package dtos

import "github.com/Junkes887/transfers-api/internal/domain/model"

type LoginDtoInput struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}

type LoginDtoOutput struct {
	Token string `json:"token"`
}

func LoginDtoInputToLoginModel(dto *LoginDtoInput) *model.LoginModel {
	return &model.LoginModel{
		CPF:    dto.CPF,
		Secret: dto.Secret,
	}
}
