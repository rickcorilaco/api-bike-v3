package repository

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	bikeRepository "github.com/rickcorilaco/api-bike-v3/src/repository/bike"
	rideRepository "github.com/rickcorilaco/api-bike-v3/src/repository/ride"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type Repositories struct {
	Bike bikeRepository.Repository
	Ride rideRepository.Repository
}

// Start initialize database connection and repositories
func Start(config Config) (repositories Repositories, err error) {
	db, err := getDB(config)
	if err != nil {
		return
	}

	repositories.Bike, err = bikeRepository.New(db)
	if err != nil {
		return
	}

	repositories.Ride, err = rideRepository.New(db)
	return
}

func getDB(config Config) (db interface{}, err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.Host, config.User, config.Password, config.Name, config.Port)

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	err = gormDB.AutoMigrate(&bikeRepository.Bike{}, &rideRepository.Ride{})
	if err != nil {
		return
	}

	db = gormDB
	return
}
