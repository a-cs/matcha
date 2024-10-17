package usecase

import (
	"backend/internal/application/domain/entity"
	"backend/internal/application/domain/step"
	"backend/internal/application/port/in"
	"backend/internal/utils"
)

type createUser struct {
}

func NewCreateUserUseCase() in.CreateUser {
	return &createUser{}
}

func getSteps() []utils.StepFunc {
	return []utils.StepFunc{
		step.HashPasswordStep,
		step.GenerateSlugIDStep,
		step.SaveUserOnDatabaseStep,
		step.SendEmailStep,
	}
}

func (c *createUser) Execute(intention *entity.CreateUserIntention) error {
	err := utils.Run(getSteps(), intention)
	if err != nil {
		return err
	}
	return nil
}
