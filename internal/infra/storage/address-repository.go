package storage

import (
	m "github.com/cezarcruz/gym-address/internal/domain/model"
)

const selectAddressByZipCode = `
	SELECT
		zipcode,
		street,
		complement,
		neighborhood,
		city,
		state
	FROM address_cache WHERE zipcode = ?	
`

const insertAddressQuery = `
	INSERT INTO address_cache (
		zipcode,
		street,
		complement,
		neighborhood,
		city,
		state
	) VALUES (
		?,
		?,
		?,
		?,
		?,
		?
	);
`

func GetAddressBy(zipcode string) *m.Address {
	var a m.Address

	row := Db.QueryRow(selectAddressByZipCode, zipcode).Scan(&a.Zipcode, &a.Street, &a.Complement, &a.Neighborhood, &a.City, &a.State)

	if row != nil {
		return nil
	}

	return &a
}

func SaveAddress(a *m.Address) {

	if Db == nil {
		panic("fudeu")
	}

	_, err := Db.Exec(
		insertAddressQuery,
		a.Zipcode,
		a.Street,
		a.Complement,
		a.Neighborhood,
		a.City,
		a.State,
	)

	if err != nil {
		panic("i morreu")
	}
}
