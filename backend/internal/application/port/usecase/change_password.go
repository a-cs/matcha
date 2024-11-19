package usecase

import (
	"backend/internal/application/domain/entity"
	"backend/internal/application/domain/step"
	"backend/internal/application/port/in"
	"backend/internal/utils"
)

type changePassword struct {
	steps []utils.StepFunc
}

func NewChangePasswordUseCase() in.ChangePassword {
	return &changePassword{
		steps: []utils.StepFunc{
			step.ValidateJwtStep,
			step.GetUserByUsernameStep,
			step.HashPasswordStep,
			step.UpdateUserPasswordStep,
		},
	}
}

func (c *changePassword) Execute(intention *entity.PasswordIntention) error {
	err := utils.Run(c.steps, intention)
	if err != nil {
		return err
	}
	return nil
}
