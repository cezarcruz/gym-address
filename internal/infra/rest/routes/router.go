package routes

import (
	"github.com/cezarcruz/gym-address/internal/infra/rest/controller"
	"github.com/labstack/echo/v4"
)

func Load(e *echo.Echo) {
	e.GET("/address/:zipcode", controller.GetAddressController)
}
