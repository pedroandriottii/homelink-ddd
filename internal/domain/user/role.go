package user

type Role string

const (
	RoleMaster Role = "MASTER"
	RoleAdmin  Role = "ADMIN"
	RoleClient Role = "CLIENT"
)

func (r Role) Valid() bool {
	return r == RoleAdmin || r == RoleClient
}
