package step

import (
	"errors"
	"os"
	"time"

	"backend/internal/application/domain/entity"
	"backend/internal/defines"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWTStep(e interface{}) error {
	loginIntention, ok := e.(*entity.LoginIntention)
	if ok {
		exp := time.Now().Add(defines.JwtExpirationTime).Unix()
		res, err := generateJwt(loginIntention.User.ID, loginIntention.User.Username, loginIntention.ProfileStatus, exp)
		if err != nil {
			return err
		}
		loginIntention.JwtResponse = *res
		return nil
	}
	passwordIntention, ok := e.(*entity.PasswordIntention)
	if ok {
		exp := time.Now().Add(defines.RecoverPasswordJwtExpirationTime).Unix()
		res, err := generateJwt(passwordIntention.User.ID, passwordIntention.User.Username, defines.EmptyString, exp)
		if err != nil {
			return err
		}
		passwordIntention.User.RecoverPasswordSlugID = res.Token
		return nil
	}

	return errors.New(defines.CannotGenerateJwt)
}

func generateJwt(userID uint64, username, profileStatus string, exp int64) (*entity.JwtResponse, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, createJwtObj(userID, username, exp))
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return nil, errors.New(defines.CannotGenerateJwt)
	}

	return &entity.JwtResponse{
		Token:         tokenString,
		ProfileStatus: profileStatus,
		UserID:        userID,
	}, nil
}

func createJwtObj(userID uint64, username string, exp int64) jwt.MapClaims {
	return jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"iat":      time.Now().Unix(),
		"exp":      exp,
	}
}
