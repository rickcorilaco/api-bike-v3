package presentation

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/rickcorilaco/api-bike-v3/src/core/service"
	bikePresentation "github.com/rickcorilaco/api-bike-v3/src/presentation/bike"
	ridePresentation "github.com/rickcorilaco/api-bike-v3/src/presentation/ride"
)

type Config struct {
	ApiPort  string
	Services service.Services
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

	if config.Services.Bike != nil {
		if _, err = bikePresentation.NewAPI(e, config.Services.Bike); err != nil {
			return
		}

		e.Logger.Print("Bike API started!")
	}

	if config.Services.Ride != nil {
		if _, err = ridePresentation.NewAPI(e, config.Services.Ride); err != nil {
			return
		}

		e.Logger.Print("Ride API started!")
	}

	go func() {
		e.Logger.Fatal(e.Start(":" + config.ApiPort))
	}()

	return
}
