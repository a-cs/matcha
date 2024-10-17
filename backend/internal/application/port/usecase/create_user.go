package usecase

import "backend/internal/application/port/in"

type createUser struct {
}

func NewCreateUserUseCase() in.CreateUser {
	return &createUser{}
}
