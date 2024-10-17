package mapper

import (
	"backend/internal/adpter/dto"
	"backend/internal/application/domain/entity"
	"github.com/peteprogrammer/go-automapper"
)

func ToCreateUser(source *dto.CreateUserDto) *entity.CreateUser {
	if source == nil {
		return nil
	}
	dest := entity.CreateUser{}
	automapper.MapLoose(source, &dest)
	return &dest
}

func ToUser(source *dto.UserDto) *entity.User {
	if source == nil {
		return nil
	}
	dest := entity.User{}
	automapper.MapLoose(source, &dest)
	return &dest
}
