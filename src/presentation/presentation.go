package presentation

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	bikeService "github.com/rickcorilaco/api-bike-v3/src/core/service/bike"
	bikePresentation "github.com/rickcorilaco/api-bike-v3/src/presentation/bike"
)

type Config struct {
	ApiPort  string
	Services []interface{}
}

func Start(config Config) (err error) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	for _, service := range config.Services {
		switch value := service.(type) {
		case *bikeService.BikeService:
			_, err = bikePresentation.NewAPI(e, value)
			if err != nil {
				return

			}
		default:
			err = fmt.Errorf("invalid service: %v", service)
		}
	}

	go func() {
		e.Logger.Fatal(e.Start(":" + config.ApiPort))
	}()

	return
}
