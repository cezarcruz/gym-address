package repository

import (
	"database/sql"
	"fmt"

	d "github.com/cezarcruz/gym-address/domain"
)

func AddressByZipCode(zipcode string) d.Address {

	db, err := sql.Open("mysql", "root:root@/gym")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	var address d.Address

	row := db.QueryRow("SELECT a.city FROM address a WHERE zipcode = ?", zipcode)

	row.Scan(&address.City)

	fmt.Sprintln(address)

	return address
}
