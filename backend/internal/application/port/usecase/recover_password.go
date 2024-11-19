package usecase

import (
	"backend/internal/application/domain/entity"
	"backend/internal/application/domain/step"
	"backend/internal/application/port/in"
	"backend/internal/utils"
)

type recoverPassword struct {
	steps []utils.StepFunc
}

func NewRecoverPasswordUseCase() in.RecoverPassword {
	return &recoverPassword{
		steps: []utils.StepFunc{
			step.GetUserByEmailStep,
			step.GenerateJWTStep,
			step.UpdateUserOnDatabaseStep,
			step.SendEmailStep,
		},
	}
}

func (c *recoverPassword) Execute(intention *entity.PasswordIntention) error {
	err := utils.Run(c.steps, intention)
	if err != nil {
		return err
	}
	return nil
}
