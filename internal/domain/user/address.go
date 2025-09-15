package user

import (
	"errors"
	"strings"

	"github.com/google/uuid"
)

var ErrInvalidAddress = errors.New("invalid address")

type AddressID = uuid.UUID

type Address struct {
	id       AddressID
	street   string
	number   string
	comp     string
	district string
	city     string
	state    string
	zip      string
	country  string
}

func NewAddress(street, number, comp, district, city, state, zip, country string) (Address, error) {
	trim := func(s string) string { return strings.TrimSpace(s) }
	street, number, city, state, zip, country = trim(street), trim(number), trim(city), trim(state), trim(zip), trim(country)
	if street == "" || city == "" || state == "" || country == "" {
		return Address{}, ErrInvalidAddress
	}
	return Address{
		id:       uuid.New(),
		street:   street,
		number:   number,
		comp:     trim(comp),
		district: trim(district),
		city:     city,
		state:    state,
		zip:      zip,
		country:  country,
	}, nil
}

func (a Address) ID() AddressID      { return a.id }
func (a Address) Street() string     { return a.street }
func (a Address) Number() string     { return a.number }
func (a Address) Complement() string { return a.comp }
func (a Address) District() string   { return a.district }
func (a Address) City() string       { return a.city }
func (a Address) State() string      { return a.state }
func (a Address) ZIP() string        { return a.zip }
func (a Address) Country() string    { return a.country }
