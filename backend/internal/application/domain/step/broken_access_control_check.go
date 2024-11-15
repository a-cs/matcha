package step

import (
	"backend/internal/application/domain/entity"
	"backend/internal/defines"
	"errors"
)

func BrokenAccessControlCheckStep(e interface{}) error {
	profileIntention, ok := e.(*entity.ProfileIntention)
	if ok {
		return brokenAccessControlCheck(
			profileIntention.JwtObj.UserID, profileIntention.ProfileRequest.UserID, defines.EmptyString, defines.EmptyString)
	}

	return errors.New(defines.CannotCheckBrokenAccessControl)
}

func brokenAccessControlCheck(jwtUserID, requestUserID uint64, jwtUsername, requestUsername string) error {
	if jwtUserID > 0 && requestUserID > 0 {
		if !areUserIDsEqual(jwtUserID, requestUserID) {
			return errors.New(defines.UnauthorizedRequest)
		}
	}
	if jwtUsername != defines.EmptyString && requestUsername != defines.EmptyString {
		if !areUsernamesEqual(jwtUsername, requestUsername) {
			return errors.New(defines.UnauthorizedRequest)
		}
	}
	if jwtUserID == 0 && requestUserID == 0 && jwtUsername == defines.EmptyString && requestUsername == defines.EmptyString {
		return errors.New(defines.CannotValidateRequest)
	}
	return nil
}

func areUserIDsEqual(jwtUserID, requestUserID uint64) bool {
	return jwtUserID == requestUserID
}

func areUsernamesEqual(jwtUsername, requestUsername string) bool {
	return jwtUsername == requestUsername
}
