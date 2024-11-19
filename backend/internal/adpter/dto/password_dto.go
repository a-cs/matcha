package dto

import "backend/internal/utils"

type RecoverRequestDto struct {
	Email string `json:"email"`
}

type ChangePasswordDto struct {
	Token       string `json:"token"`
	NewPassword string `json:"new_password"`
}

func (r RecoverRequestDto) IsValid() bool {
	return len(r.Email) > 0 && !utils.IsSQLInjection(r.Email) && utils.IsValidEmail(r.Email)
}

func (c ChangePasswordDto) IsValid() bool {
	return len(c.Token) > 0 && len(c.NewPassword) > 0 && !utils.IsSQLInjection(c.NewPassword, c.Token)
}
