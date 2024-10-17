package in

import "backend/internal/application/domain/entity"

type CreateUser interface {
	Execute(createUserIntention *entity.CreateUserIntention) error
}
