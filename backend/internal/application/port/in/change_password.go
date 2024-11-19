package in

import "backend/internal/application/domain/entity"

type ChangePassword interface {
	Execute(passwordIntention *entity.PasswordIntention) error
}
