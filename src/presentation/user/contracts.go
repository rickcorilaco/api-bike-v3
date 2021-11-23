package bike

import (
	"github.com/google/uuid"

	"github.com/rickcorilaco/api-bike-v3/src/core/domain"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name" validate:"required"`
	Username string    `json:"username" validate:"required"`
	Password string    `json:"password" validate:"required"`
}

func (ref *User) FromDomain(dom domain.User) {
	ref.ID = dom.ID
	ref.Name = dom.Name
	ref.Username = dom.Username
	ref.Password = dom.Password
}

func (ref *User) ToDomain() (dom domain.User) {
	dom = domain.User{
		ID:       ref.ID,
		Name:     ref.Name,
		Username: ref.Username,
		Password: ref.Password,
	}

	return
}

type Users []User

func (ref *Users) FromDomain(dom *domain.Users) {
	for _, domUser := range *dom {
		user := User{}
		user.FromDomain(domUser)
		*ref = append(*ref, user)
	}

	return
}

func (ref *Users) ToDomain() (dom *domain.Users) {
	dom = &domain.Users{}

	for _, refUser := range *ref {
		*dom = append(*dom, refUser.ToDomain())
	}

	return
}
