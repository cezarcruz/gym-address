package controller

import (
	"net/http"

	"github.com/cezarcruz/gym-address/internal/domain/usecase"
	"github.com/labstack/echo/v4"
)

func GetAddressController(c echo.Context) error {
	z := c.Param("zipcode")
	a := usecase.GetAndCacheAddress(z)
	return c.JSON(http.StatusOK, a)
}
