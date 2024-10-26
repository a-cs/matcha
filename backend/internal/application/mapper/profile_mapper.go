package mapper

import (
	"backend/internal/adpter/dto"
	"backend/internal/application/domain/entity"
	"github.com/peteprogrammer/go-automapper"
)

func ToProfile(source *dto.ProfileDto) *entity.Profile {
	if source == nil {
		return nil
	}
	dest := entity.Profile{}
	automapper.MapLoose(source, &dest)
	return &dest
}
