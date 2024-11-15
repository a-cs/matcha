package dto

import (
	"errors"
	"time"
)

type JwtObjDto struct {
	UserID   uint64 `json:"user_id"`
	Username string `json:"username"`
	Iat      int64  `json:"iat"`
	Exp      int64  `json:"exp"`
}

type JwtResponseDto struct {
	Token         string `json:"token"`
	ProfileStatus string `json:"profile_status"`
	UserID        uint64 `json:"user_id"`
}

func (j JwtObjDto) Valid() error {
	if j.UserID <= 0 || len(j.Username) == 0 || j.Iat <= 0 || j.Exp <= 0 ||
		j.Exp <= j.Iat || j.Exp <= time.Now().Unix() || j.Iat > time.Now().Unix() {
		return errors.New("invalid JWT object")
	}
	return nil
}
