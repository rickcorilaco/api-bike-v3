package bike

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/labstack/echo"
	"github.com/rickcorilaco/api-bike-v3/src/core/ports"
	"github.com/rickcorilaco/api-bike-v3/src/core/value"
)

type API struct {
	e           *echo.Echo
	bikeService ports.BikeService
}

func NewAPI(e *echo.Echo, bikeService ports.BikeService) (bikePresentation *API, err error) {
	bikePresentation = &API{
		e:           e,
		bikeService: bikeService,
	}

	bikeGroup := e.Group("bikes")
	bikeGroup.GET("", bikePresentation.List)
	bikeGroup.GET("/:bike_id", bikePresentation.Get)
	bikeGroup.POST("", bikePresentation.Create)
	bikeGroup.DELETE("/:bike_id", bikePresentation.Delete)

	return
}

func (ref *API) List(c echo.Context) (err error) {
	filter := value.BikeListFilter{}

	if err = c.Bind(&filter); err != nil {
		return
	}

	result, err := ref.bikeService.List(filter)
	if err != nil {
		return
	}

	response := Bikes{}
	response.FromDomain(result)
	return c.JSON(http.StatusOK, response)
}

func (ref *API) Get(c echo.Context) (err error) {
	bikeID, err := uuid.Parse(c.Param("bike_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	result, err := ref.bikeService.Get(bikeID)
	if err != nil {
		return
	}

	if result == nil {
		return c.NoContent(http.StatusNotFound)
	}

	response := Bike{}
	response.FromDomain(*result)
	return c.JSON(http.StatusOK, response)
}

func (ref *API) Create(c echo.Context) (err error) {
	payload := Bike{}

	if err = c.Bind(&payload); err != nil {
		return
	}

	if err = c.Validate(payload); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	result, err := ref.bikeService.Create(payload.ToDomain())
	if err != nil {
		return
	}

	payload.FromDomain(*result)

	return c.JSON(http.StatusCreated, payload)
}

func (ref *API) Delete(c echo.Context) (err error) {
	bikeID, err := uuid.Parse(c.Param("bike_id"))
	if err != nil {
		return
	}

	bike := Bike{ID: bikeID}

	result, err := ref.bikeService.Delete(bike.ToDomain())
	if err != nil {
		return
	}

	if result == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.NoContent(http.StatusNoContent)
}
