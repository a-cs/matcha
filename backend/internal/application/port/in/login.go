package in

import "backend/internal/application/domain/entity"

type Login interface {
	Execute(loginIntention *entity.LoginIntention) error
}
