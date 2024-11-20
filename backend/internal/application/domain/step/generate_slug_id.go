package step

import (
	"backend/internal/application/domain/entity"
	"backend/internal/defines"
	"errors"
	"github.com/google/uuid"
)

func GenerateSlugIDStep(e interface{}) error {
	intention, ok := e.(*entity.CreateUserIntention)
	if !ok {
		return errors.New(defines.CannotGenerateSlugID)
	}

	return generateUUID(intention)
}

func generateUUID(intention *entity.CreateUserIntention) error {
	u, err := uuid.NewUUID()
	if err != nil {
		return errors.New(defines.CannotGenerateSlugID)
	}
	intention.SlugID = u.String()
	return nil
}
