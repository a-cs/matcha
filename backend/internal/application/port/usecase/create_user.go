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
			step.CreateProfileOnDatabaseStep,
			// TODO: step.DeleteUserIfErrorStep,
			step.SendEmailStep,
			// TODO: step.DeleteUserIfErrorStep,
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
