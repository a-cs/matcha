package in

import "backend/internal/application/domain/entity"

type ConfirmAccount interface {
	Execute(confirmAccountIntention *entity.ConfirmAccountIntention) error
}
