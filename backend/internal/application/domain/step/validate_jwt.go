package step

import (
	"backend/internal/adpter/dto"
	"backend/internal/application/domain/entity"
	"backend/internal/application/mapper"
	"backend/internal/defines"
	"errors"
	"github.com/golang-jwt/jwt"
	"os"
)

const (
	invalidJwtSignature = "invalid signature"
	invalidJwtToken     = "invalid token"
)

func ValidateJwtStep(e interface{}) error {
	profileIntention, ok := e.(*entity.ProfileIntention)
	if ok {
		res, err := validateJwt(profileIntention.JwtToken)
		if err != nil {
			return err
		}
		profileIntention.JwtObj = *res
		return nil
	}
	changePasswordIntention, ok := e.(*entity.PasswordIntention)
	if ok {
		res, err := validateJwt(changePasswordIntention.JwtToken)
		if err != nil {
			return err
		}
		changePasswordIntention.JwtObj = *res
		return nil
	}

	return errors.New(defines.CannotValidateJwt)
}

func validateJwt(jwtToken string) (*entity.JwtObj, error) {
	jwtKey := []byte(os.Getenv("JWT_SECRET"))
	claims := &dto.JwtObjDto{}

	token, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, errors.New(invalidJwtSignature)
		}
		return nil, errors.New(invalidJwtToken)
	}

	if !token.Valid {
		return nil, errors.New(invalidJwtToken)
	}

	return mapper.FromJwtObjDtoToJwtObj(claims), nil
}
