package bike

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/rickcorilaco/api-bike-v3/src/core/domain"
	"github.com/rickcorilaco/api-bike-v3/src/core/ports"
)

type Service struct {
	repository ports.UserRepository
	tokenKey   string
}

var (
	ErrInvalidUsernameOrPassword = errors.New("invalid username or password")
	ErrDuplicateRecord = errors.New("duplicate user")
)

func New(repository ports.UserRepository, tokenKey string) (userService ports.UserService, err error) {
	userService = &Service{
		repository: repository,
		tokenKey:   tokenKey,
	}
	return
}

func (ref *Service) Create(user domain.User) (result *domain.User, err error) {
	result, err = ref.repository.Create(user)
	if errors.Is(err, ref.repository.ErrDuplicateRecord()) {
		return nil, ref.ErrDuplicateRecord()
	}

	return
}

func (ref *Service) Login(username, password string) (token string, err error) {
	user, err := ref.getByUsername(username)
	if err != nil {
		return
	}

	if user == nil || user.Password != password {
		err = ref.ErrInvalidUsernameOrPassword()
		return
	}

	token, err = ref.generateToken(*user)
	return
}

func (ref *Service) ErrInvalidUsernameOrPassword() (err error) {
	return ErrInvalidUsernameOrPassword
}

func (ref *Service) ErrDuplicateRecord() (err error) {
	return ErrDuplicateRecord
}

func (ref *Service) getByUsername(username string) (user *domain.User, err error) {
	 return ref.repository.GetByUsername(username)
}

func (ref *Service) generateToken(user domain.User) (token string, err error) {
	claims := jwt.MapClaims{
		"username":  user.Username,
		"create_at": time.Now().Unix(),
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS512, claims).SignedString([]byte(ref.tokenKey))
}
