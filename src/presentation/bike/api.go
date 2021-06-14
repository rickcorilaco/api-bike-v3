package bike

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	bikeService "github.com/rickcorilaco/api-bike-v3/src/core/service/bike"
)

type BikeAPI struct {
	e           *echo.Echo
	bikeService *bikeService.BikeService
}

func NewAPI(e *echo.Echo, bikeService *bikeService.BikeService) (bikePresentation *BikeAPI, err error) {
	bikePresentation = &BikeAPI{
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

func (ref *BikeAPI) List(c echo.Context) (err error) {
	result, err := ref.bikeService.List()
	if err != nil {
		return
	}

	response := Bikes{}
	response.FromDomain(result)
	return c.JSON(http.StatusOK, response)
}

func (ref *BikeAPI) Get(c echo.Context) (err error) {
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

func (ref *BikeAPI) Create(c echo.Context) (err error) {
	payload := Bike{}

	if err = c.Bind(&payload); err != nil {
		return
	}

	// if err = c.Validate(payload); err != nil {
	// 	return
	// }

	result, err := ref.bikeService.Create(payload.ToDomain())
	if err != nil {
		return
	}

	payload.FromDomain(*result)

	return c.JSON(http.StatusCreated, payload)
}

func (ref *BikeAPI) Delete(c echo.Context) (err error) {
	bikeID, err := uuid.Parse(c.Param("bike_id"))
	if err != nil {
		return
	}

	bike := Bike{ID: bikeID}

	err = ref.bikeService.Delete(bike.ToDomain())
	if err != nil {
		return
	}

	return c.NoContent(http.StatusOK)
}
