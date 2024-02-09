package address

import (
	"fmt"
	"net/http"

	"github.com/cezarcruz/gym-address/integration"
	"github.com/cezarcruz/gym-address/repository"
	"github.com/gin-gonic/gin"
)

func AddressRetrieve(c *gin.Context) {
	zipcode := c.Query("zipcode")

	a := repository.AddressByZipCode(zipcode)

	fmt.Println("reult", a)

	r := integration.GetAddress(zipcode)

	c.JSON(http.StatusOK, r)

}
