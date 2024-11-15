package in

import "backend/internal/application/domain/entity"

type UpdateProfile interface {
	Execute(profileIntention *entity.ProfileIntention) error
}
