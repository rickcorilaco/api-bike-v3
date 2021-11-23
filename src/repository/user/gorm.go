package bike

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"

	"github.com/rickcorilaco/api-bike-v3/src/core/domain"
)

var (
	ErrDuplicateRecord = errors.New("duplicate record")
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) (gormUserRepository *GormUserRepository, err error) {
	gormUserRepository = &GormUserRepository{db: db}
	return
}

func (ref *GormUserRepository) GetByUsername(username string) (result *domain.User, err error) {
	model := &User{}

	err = ref.db.Where("username = ?", username).First(&model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = nil
		}

		return
	}

	result = model.ToDomain()
	return
}

func (ref *GormUserRepository) Create(user domain.User) (result *domain.User, err error) {
	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}

	model := &User{}
	model.FromDomain(&user)

	err = ref.db.Create(model).Error
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr.Code == "23505" {
				err = ref.ErrDuplicateRecord()
			}
		}

		return
	}

	result = model.ToDomain()
	return
}

func (ref *GormUserRepository) ErrDuplicateRecord() (err error) {
	return ErrDuplicateRecord
}
