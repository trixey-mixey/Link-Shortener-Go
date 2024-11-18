package di

import "go/projcet-Adv/internal/users"

type IStatRepositrory interface {
	AddClick(linkId uint)
}

type IUserRepository interface {
	Create(user *users.User) (*users.User, error)
	FindByEmail(email string) (*users.User, error)
}
