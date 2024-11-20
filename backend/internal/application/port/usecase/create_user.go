package usecase

import (
	"backend/internal/application/domain/entity"
	"backend/internal/application/domain/step"
	"backend/internal/application/port/in"
	"backend/internal/utils"
)

type createUser struct {
	steps []utils.StepFunc
}

func NewCreateUserUseCase() in.CreateUser {
	return &createUser{
		steps: []utils.StepFunc{
			step.HashPasswordStep,
			step.GenerateSlugIDStep,
			step.SaveUserOnDatabaseStep,
			step.GetUserByEmailStep,
			step.DeleteUserByEmailStep,
			step.CreateProfileOnDatabaseStep,
			step.DeleteUserByEmailStep,
			step.SendEmailStep,
			step.DeleteProfileByUserIDStep,
			step.DeleteUserByEmailStep,
		},
	}
}

func (c *createUser) Execute(intention *entity.CreateUserIntention) error {
	err := utils.Run(c.steps, intention)
	if err != nil {
		return err
	}
	return nil
}
