package usecase

import (
	"backend/internal/application/domain/entity"
	"backend/internal/application/domain/step"
	"backend/internal/application/port/in"
	"backend/internal/utils"
)

type getProfile struct {
	steps []utils.StepFunc
}

func NewGetProfileUseCase() in.GetProfile {
	return &getProfile{
		steps: []utils.StepFunc{
			step.GetUserByUsernameStep,
			step.GetProfileByUserIDStep,
		},
	}
}

func (c *getProfile) Execute(intention *entity.ProfileIntention) error {
	err := utils.Run(c.steps, intention)
	if err != nil {
		return err
	}
	return nil
}
