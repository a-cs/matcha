package step

import (
	"backend/internal/application/domain/entity"
	"backend/internal/defines"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func ValidateLoginStep(e interface{}) error {
	loginIntention, ok := e.(*entity.LoginIntention)
	if ok {
		return validateLogin(loginIntention)
	}

	return errors.New(defines.CannotValidateLogin)
}

func validateLogin(intention *entity.LoginIntention) error {
	if err := bcrypt.CompareHashAndPassword([]byte(intention.User.Password), []byte(intention.Password)); err != nil {
		return errors.New(defines.InvalidPassword)
	}
	if intention.User.AccountStatus != defines.ActiveStatus {
		return errors.New(defines.AccountNotActiveYet)
	}

	return nil
}
