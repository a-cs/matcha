package step

import (
	"backend/internal/application/domain/entity"
	"backend/internal/defines"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func HashPasswordStep(e interface{}) error {
	intention, ok := e.(*entity.CreateUserIntention)
	if !ok {
		return errors.New(defines.CannotHashPassword)
	}

	return hashPassword(intention)
}

func hashPassword(intention *entity.CreateUserIntention) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(intention.User.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	intention.HashedPassword = string(hashedPassword)

	return nil
}
