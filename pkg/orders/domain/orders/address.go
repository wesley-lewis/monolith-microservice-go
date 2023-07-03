package orders

import (
	"errors"
)

type Address struct {
	name     string
	street   string
	city     string
	postCode string
	country  string
}

func NewAddress(name, street, city, postCode, country string) (Address, error) {
	if len(name) == 0 || len(street) == 0 || len(city) == 0 || len(postCode) == 0 || len(country) == 0 {
		return Address{}, errors.New("Empty Address Field")
	}

	return Address{
		name,
		street,
		city,
		postCode,
		country,
	}, nil
}

func (a Address) Street() string {
	return a.street
}

func (a Address) Name() string {
	return a.name
}

func (a Address) City() string {
	return a.city
}

func (a Address) PostCode() string {
	return a.postCode
}

func (a Address) Country() string {
	return a.country
}
