package main

import (
	"database/sql"
	"time"

	"github.com/cezarcruz/gym-address/address"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	configDatabase()

	e := echo.New()

	e.GET("/address/:zipcode", address.GetAddress)

	e.Logger.Fatal(e.Start((":8081")))
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
