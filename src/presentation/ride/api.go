package ride

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"

	"github.com/rickcorilaco/api-bike-v3/src/core/ports"
	"github.com/rickcorilaco/api-bike-v3/src/core/value"
)

type API struct {
	e           *echo.Echo
	rideService ports.RideService
}

func NewAPI(e *echo.Echo, rideService ports.RideService) (ridePresentation *API, err error) {
	ridePresentation = &API{
		e:           e,
		rideService: rideService,
	}

	rideGroup := e.Group("rides")
	rideGroup.GET("", ridePresentation.List)
	rideGroup.GET("/:ride_id", ridePresentation.Get)
	rideGroup.POST("", ridePresentation.Create)
	rideGroup.DELETE("/:ride_id", ridePresentation.Delete)

	return
}

func (ref *API) List(c echo.Context) (err error) {
	filter := value.RideListFilter{}

	if err = c.Bind(&filter); err != nil {
		return
	}

	result, err := ref.rideService.List(filter)
	if err != nil {
		return
	}

	response := Rides{}
	response.FromDomain(result)
	return c.JSON(http.StatusOK, response)
}

func (ref *API) Get(c echo.Context) (err error) {
	rideID, err := uuid.Parse(c.Param("ride_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	result, err := ref.rideService.Get(rideID)
	if err != nil {
		return
	}

	if result == nil {
		return c.NoContent(http.StatusNotFound)
	}

	response := Ride{}
	response.FromDomain(*result)
	return c.JSON(http.StatusOK, response)
}

func (ref *API) Create(c echo.Context) (err error) {
	payload := Ride{}

	if err = c.Bind(&payload); err != nil {
		return
	}

	if err = c.Validate(payload); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	dom, err := payload.ToDomain()
	if err != nil {
		return
	}

	result, err := ref.rideService.Create(dom)
	if err != nil {
		return
	}

	payload.FromDomain(*result)

	return c.JSON(http.StatusCreated, payload)
}

func (ref *API) Delete(c echo.Context) (err error) {
	rideID, err := uuid.Parse(c.Param("ride_id"))
	if err != nil {
		return
	}

	ride := Ride{ID: rideID}

	dom, err := ride.ToDomain()
	if err != nil {
		return
	}

	result, err := ref.rideService.Delete(dom)
	if err != nil {
		return
	}

	if result == nil {
		return c.NoContent(http.StatusNotFound)
	}

	return c.NoContent(http.StatusNoContent)
}
