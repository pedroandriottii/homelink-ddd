package user

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type User struct {
	id        uuid.UUID
	role      Role
	name      string
	phone     string
	email     Email
	password  string
	addresses map[AddressID]Address

	createdAt time.Time
	updatedAt time.Time
}

func NewUser(role Role, name, phone, password string, email Email) (*User, error) {
	now := time.Now().UTC()
	u := &User{
		id:        uuid.New(),
		role:      role,
		name:      name,
		phone:     strings.TrimSpace(phone),
		email:     email,
		password:  password,
		addresses: make(map[AddressID]Address),
		createdAt: now,
		updatedAt: now,
	}
	if err := u.Valid(); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) Valid() error {
	u.name = strings.TrimSpace(u.name)
	if !u.role.Valid() {
		return ErrInvalidRole
	}
	if u.name == "" {
		return ErrEmptyName
	}
	if strings.TrimSpace(u.password) == "" {
		return ErrEmptyName
	}
	return nil
}

func (u *User) ID() uuid.UUID        { return u.id }
func (u *User) Role() Role           { return u.role }
func (u *User) Name() string         { return u.name }
func (u *User) Email() Email         { return u.email }
func (u *User) Phone() string        { return u.phone }
func (u *User) CreatedAt() time.Time { return u.createdAt }
func (u *User) UpdatedAt() time.Time { return u.updatedAt }

func (u *User) ChangeEmail(e Email) {
	u.email = e
	u.touch()
}

func (u *User) ChangePhone(p string) {
	u.phone = strings.TrimSpace(p)
	u.touch()
}

func (u *User) AddAddress(a Address) {
	u.addresses[a.ID()] = a
	u.touch()
}

func (u *User) RemoveAddress(id AddressID) error {
	if _, ok := u.addresses[id]; !ok {
		return ErrAddressNotFound
	}
	delete(u.addresses, id)
	u.touch()
	return nil
}

func (u *User) Addresses() []Address {
	out := make([]Address, 0, len(u.addresses))
	for _, a := range u.addresses {
		out = append(out, a)
	}
	return out
}

func (u *User) touch() { u.updatedAt = time.Now().UTC() }
