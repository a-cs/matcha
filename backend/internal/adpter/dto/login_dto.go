package dto

import "backend/internal/utils"

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c LoginDto) IsValid() bool {
	return !utils.IsSQLInjection(c.Email, c.Password) &&
		utils.IsValidEmail(c.Email) &&
		len(c.Password) > 0
}
