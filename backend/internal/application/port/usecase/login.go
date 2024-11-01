package usecase

import (
	"backend/internal/application/domain/entity"
	"backend/internal/application/domain/step"
	"backend/internal/application/port/in"
	"backend/internal/utils"
)

type login struct {
	steps []utils.StepFunc
}

func NewLoginUseCase() in.Login {
	return &login{
		steps: []utils.StepFunc{
			step.GetUserByEmailStep,
			step.ValidateLoginStep,
			step.GetProfileByUserIDStep,
			step.GenerateJWTStep,
		},
	}
}

func (c *login) Execute(intention *entity.LoginIntention) error {
	err := utils.Run(c.steps, intention)
	if err != nil {
		return err
	}
	return nil
}
