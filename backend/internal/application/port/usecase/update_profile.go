package usecase

import (
	"backend/internal/application/domain/entity"
	"backend/internal/application/domain/step"
	"backend/internal/application/port/in"
	"backend/internal/utils"
)

type updateProfile struct {
	steps []utils.StepFunc
}

func NewUpdateProfileUseCase() in.UpdateProfile {
	return &updateProfile{
		steps: []utils.StepFunc{
			step.ValidateJwtStep,
			step.GetProfileByUserIDStep,
			step.BrokenAccessControlCheckStep,
			step.UpdateProfileOnDatabaseStep,
		},
	}
}

func (c *updateProfile) Execute(intention *entity.ProfileIntention) error {
	err := utils.Run(c.steps, intention)
	if err != nil {
		return err
	}
	return nil
}
