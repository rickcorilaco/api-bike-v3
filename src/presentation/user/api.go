package bike

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"

	"github.com/rickcorilaco/api-bike-v3/src/core/ports"
)

type API struct {
	e           *echo.Echo
	userService ports.UserService
}

func NewAPI(e *echo.Echo, userService ports.UserService) (userPresentation *API, err error) {
	userPresentation = &API{
		e:           e,
		userService: userService,
	}

	bikeGroup := e.Group("users")
	bikeGroup.POST("/login", userPresentation.Login)
	bikeGroup.POST("", userPresentation.Create)

	return
}

func (ref *API) Login(c echo.Context) (err error) {
	payload := struct{
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if err = c.Bind(&payload); err != nil {
		return
	}

	if err = c.Validate(payload); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	token, err := ref.userService.Login(payload.Username, payload.Password)
	if err != nil {
		if errors.Is(err, ref.userService.ErrInvalidUsernameOrPassword()) {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid username or password"})
		}

		return
	}

	return c.JSON(http.StatusOK, echo.Map{"token": token})
}

func (ref *API) Create(c echo.Context) (err error) {
	payload := User{}

	if err = c.Bind(&payload); err != nil {
		return
	}

	if err = c.Validate(payload); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	result, err := ref.userService.Create(payload.ToDomain())
	if err != nil {
		if errors.Is(err, ref.userService.ErrDuplicateRecord()) {
			return c.JSON(http.StatusConflict, echo.Map{"error": "user already exists"})
		}

		return
	}

	payload.FromDomain(*result)

	return c.JSON(http.StatusCreated, payload)
}