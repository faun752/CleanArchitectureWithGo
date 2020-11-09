package usecase

import "CleanArchitectureWithGo/src/domain"

// UserRepository is
type UserRepository interface {
	Store(domain.User)
	Select() []domain.User
	Delete(id string)
}
