package main

import (
	"database/sql"
	"time"

	"github.com/cezarcruz/gym-address/address"
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	configDatabase()

	r := gin.Default()

	r.GET("/address", address.AddressRetrieve)

	r.Run()
}

func configDatabase() {
	db, err := sql.Open("mysql", "root:root@/gym")

	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}
