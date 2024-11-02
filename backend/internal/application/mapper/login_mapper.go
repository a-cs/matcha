package mapper

import (
	"backend/internal/adpter/dto"
	"backend/internal/application/domain/entity"
	"github.com/peteprogrammer/go-automapper"
)

func ToJwtResponseDto(source *entity.JwtResponse) *dto.JwtResponseDto {
	if source == nil {
		return nil
	}
	dest := dto.JwtResponseDto{}
	automapper.MapLoose(source, &dest)
	return &dest
}
