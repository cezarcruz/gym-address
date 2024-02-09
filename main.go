package main

import (
	"github.com/cezarcruz/gym-address/address"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/address", address.AddressRetrieve)

	r.Run()
}
