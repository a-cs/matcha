package in

import "backend/internal/application/domain/entity"

type GetProfile interface {
	Execute(profileIntention *entity.ProfileIntention) error
}
