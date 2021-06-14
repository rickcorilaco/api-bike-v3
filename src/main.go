package main

import (
	"os"
	"sync"

	"github.com/rickcorilaco/api-bike-v3/src/core/service"
	"github.com/rickcorilaco/api-bike-v3/src/presentation"
	"github.com/rickcorilaco/api-bike-v3/src/repository"
)

var (
	apiHost = os.Getenv("API_HOST")
	apiPort = os.Getenv("API_PORT")

	dbHost     = os.Getenv("DATABASE_HOST")
	dbPort     = os.Getenv("DATABASE_PORT")
	dbUser     = os.Getenv("DATABASE_USER")
	dbPassword = os.Getenv("DATABASE_PASSWORD")
	dbName     = os.Getenv("DATABASE_NAME")
)

func main() {
	// repository
	repositoryConfig := repository.Config{
		Host:     dbHost,
		Port:     dbPort,
		User:     dbUser,
		Password: dbPassword,
		Name:     dbName,
	}

	repositories, err := repository.Start(repositoryConfig)
	if err != nil {
		panic(err)
	}

	// service
	serviceConfig := service.Config{
		Repositories: repositories,
	}

	services, err := service.Start(serviceConfig)
	if err != nil {
		panic(err)
	}

	// presentation
	presentationConfig := presentation.Config{
		ApiPort:  apiPort,
		Services: services,
	}

	err = presentation.Start(presentationConfig)
	if err != nil {
		panic(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
