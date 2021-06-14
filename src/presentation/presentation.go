package presentation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	bikeService "github.com/rickcorilaco/api-bike-v3/src/core/service/bike"
	bikePresentation "github.com/rickcorilaco/api-bike-v3/src/presentation/bike"
)

type Config struct {
	ApiPort  string
	Services []interface{}
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func Start(config Config) (err error) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Validator = &CustomValidator{validator: validator.New()}

	for _, service := range config.Services {
		switch value := service.(type) {
		case *bikeService.BikeService:
			_, err = bikePresentation.NewAPI(e, value)
			if err != nil {
				return

			}
		default:
			err = fmt.Errorf("invalid service: %v", service)
			break
		}
	}

	go func() {
		e.Logger.Fatal(e.Start(":" + config.ApiPort))
	}()

	return
}
