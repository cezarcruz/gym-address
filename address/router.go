package address

import (
	"net/http"

	"github.com/cezarcruz/gym-address/integration"
	"github.com/labstack/echo/v4"
)

func GetAddress(c echo.Context) error {
	zipcode := c.Param("zipcode")

	// 	a := repository.AddressByZipCode(zipcode)
	// 	fmt.Println("reult", a)

	r := integration.GetAddress(zipcode)

	return c.JSON(http.StatusOK, r)
}
