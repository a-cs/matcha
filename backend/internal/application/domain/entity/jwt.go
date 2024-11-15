package entity

import (
	"errors"
	"time"
)

type JwtObj struct {
	UserID   uint64
	Username string
	Iat      int64
	Exp      int64
}

func (j JwtObj) Valid() error {
	if j.UserID <= 0 || len(j.Username) == 0 || j.Iat <= 0 || j.Exp <= 0 ||
		j.Exp <= j.Iat || j.Exp <= time.Now().Unix() || j.Iat > time.Now().Unix() {
		return errors.New("invalid JWT object")
	}
	return nil
}

type JwtResponse struct {
	Token         string
	ProfileStatus string
	UserID        uint64
}
