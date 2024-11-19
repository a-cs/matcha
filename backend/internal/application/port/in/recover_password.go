package in

import "backend/internal/application/domain/entity"

type RecoverPassword interface {
	Execute(passwordIntention *entity.PasswordIntention) error
}
