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
		res, err := generateJwt(loginIntention.User.ID, loginIntention.User.Username, loginIntention.ProfileStatus)
		if err != nil {
			return err
		}
		loginIntention.JwtResponse = *res
		return nil
	}

	return errors.New(defines.CannotGenerateSlugID)
}

func generateJwt(userID uint64, username, profileStatus string) (*entity.JwtResponse, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, createJwtObj(userID, username))
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	return &entity.JwtResponse{
		Token:         tokenString,
		ProfileStatus: profileStatus,
		UserID:        userID,
	}, nil
}

func createJwtObj(userID uint64, username string) jwt.MapClaims {
	return jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(defines.JwtExpirationTime).Unix(),
	}
}
