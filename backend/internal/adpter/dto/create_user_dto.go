package dto

import "backend/internal/utils"

type CreateUserDto struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (c CreateUserDto) IsValid() bool {
	return !utils.IsSQLInjection(c.Email, c.Username, c.Password, c.FirstName, c.LastName) &&
		utils.IsValidEmail(c.Email) &&
		len(c.Username) > 0 &&
		len(c.Password) > 0 &&
		len(c.FirstName) > 0 &&
		len(c.LastName) > 0
}
