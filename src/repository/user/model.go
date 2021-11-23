package bike

import (
	"github.com/google/uuid"

	"github.com/rickcorilaco/api-bike-v3/src/core/domain"
)

type User struct {
	ID       uuid.UUID `gorm:"id"`
	Name     string    `gorm:"name"`
	Username string    `gorm:"username;unique"`
	Password string    `gorm:"password"`
}

func (ref *User) FromDomain(dom *domain.User) {
	if dom == nil {
		return
	}

	ref.ID = dom.ID
	ref.Name = dom.Name
	ref.Username = dom.Username
	ref.Password = dom.Password
}

func (ref *User) ToDomain() (dom *domain.User) {
	dom = &domain.User{
		ID:    ref.ID,
		Name: ref.Name,
		Username: ref.Username,
		Password: ref.Password,
	}

	return
}

type Users []User

func (ref *Users) FromDomain(dom *domain.Users) {
	for _, domUser := range *dom {
		user := User{}
		user.FromDomain(&domUser)
		*ref = append(*ref, user)
	}
}

func (ref *Users) ToDomain() (dom *domain.Users) {
	dom = &domain.Users{}

	for _, refUser := range *ref {
		*dom = append(*dom, *refUser.ToDomain())
	}

	return
}
