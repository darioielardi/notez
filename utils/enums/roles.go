package enums

type Role string

const (
	User  Role = "user"
	Admin Role = "admin"
)

type Roles []Role

func (roles *Roles) Contains(role *Role) bool {
	for _, r := range *roles {
		if r == *role {
			return true
		}
	}
	return false
}
