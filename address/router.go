package address

import (
	"net/http"

	"github.com/cezarcruz/gym-address/integration"
	"github.com/gin-gonic/gin"
)

func AddressRetrieve(c *gin.Context) {
	zipcode := c.Query("zipcode")

	r := integration.GetAddress(zipcode)

	c.JSON(http.StatusOK, r)

}
