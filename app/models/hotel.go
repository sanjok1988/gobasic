package models

import (
	"github.com/revel/revel"
)

type Hotel struct {
	HotelId				int
	Name, Address		string
	City, State, Zip	string
	Country				string
	Price				int
}

func (hotel *Hotel) Validate(v *revel.Validation) {
	v.Check(hotel.Name,
		revel.Required{},
		revel.MaxSize{50},
	)
}