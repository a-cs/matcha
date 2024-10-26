package usecase

import (
	"backend/internal/application/domain/entity"
	"backend/internal/application/domain/step"
	"backend/internal/application/port/in"
	"backend/internal/utils"
)

type confirmAccount struct {
	steps []utils.StepFunc
}

func NewConfirmAccountUseCase() in.ConfirmAccount {
	return &confirmAccount{
		steps: []utils.StepFunc{
			step.GetUserBySlugIDStep,
			step.UpdateUserOnDatabaseStep,
		},
	}
}

func (c *confirmAccount) Execute(intention *entity.ConfirmAccountIntention) error {
	err := utils.Run(c.steps, intention)
	if err != nil {
		return err
	}
	return nil
}
