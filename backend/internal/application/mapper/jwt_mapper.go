package mapper

import (
	"backend/internal/adpter/dto"
	"backend/internal/application/domain/entity"
	"github.com/peteprogrammer/go-automapper"
)

func FromJwtObjDtoToJwtObj(source *dto.JwtObjDto) *entity.JwtObj {
	if source == nil {
		return nil
	}
	dest := entity.JwtObj{}
	automapper.MapLoose(source, &dest)
	return &dest
}
