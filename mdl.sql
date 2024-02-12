CREATE TABLE address.address_cache (
	zipcode varchar(9) NOT NULL,
	street varchar(100) NULL,
	complement varchar(100) NULL,
	neighborhood varchar(100) NULL,
	city varchar(100) NULL,
	state varchar(2) NULL,
	CONSTRAINT address_cache_pk PRIMARY KEY (zipcode)
)
