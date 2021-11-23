package bike

import (
	"errors"

	"gorm.io/gorm"

	"github.com/rickcorilaco/api-bike-v3/src/core/ports"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

func New(db interface{}) (userRepository ports.UserRepository, err error) {
	switch v := db.(type) {
	case *gorm.DB:
		userRepository, err = NewGormUserRepository(v)
	default:
		err = errors.New("invalid database")
	}

	return
}
