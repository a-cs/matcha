package step

import (
	"backend/internal/application/domain/entity"
	"backend/internal/defines"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func HashPasswordStep(e interface{}) error {
	createUserIntention, ok := e.(*entity.CreateUserIntention)
	if ok {
		hashedPassword, err := hashPassword(createUserIntention.CreateUser.Password)
		if err != nil {
			return err
		}
		createUserIntention.HashedPassword = hashedPassword
		return nil
	}
	changePasswordIntention, ok := e.(*entity.PasswordIntention)
	if ok {
		hashedPassword, err := hashPassword(changePasswordIntention.NewPassword)
		if err != nil {
			return err
		}
		changePasswordIntention.HashedPassword = hashedPassword
		return nil
	}

	return errors.New(defines.CannotHashPassword)
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return defines.EmptyString, errors.New(defines.CannotHashPassword)
	}

	return string(hashedPassword), nil
}
