package repository

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/rickcorilaco/api-bike-v3/src/core/entity"
	bikeRepository "github.com/rickcorilaco/api-bike-v3/src/repository/bike"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func Start(config Config) (repositories []interface{}, err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.Host, config.User, config.Password, config.Name, config.Port)

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	err = gormDB.AutoMigrate(&entity.Bike{})
	if err != nil {
		return
	}

	db := gormDB

	bikeRepository, err := bikeRepository.New(db)
	if err != nil {
		return
	}

	repositories = append(repositories, bikeRepository)
	return
}
