package typ

import "database/sql/driver"

type Role string

const (
	Role_ADMIN     Role = "ADMIN"
	Role_GUEST     Role = "INITIATOR"
	Role_MODERATOR Role = "MODERATOR"
)

var roles = map[Role]string{
	Role_ADMIN:     Role_ADMIN.String(),
	Role_GUEST:     Role_GUEST.String(),
	Role_MODERATOR: Role_MODERATOR.String(),
}

func (s Role) Enum() map[Role]string {
	return roles
}

func (s Role) Valid() bool {
	_, ok := roles[s]
	return ok
}

func (s Role) String() string {
	return string(s)
}

func (s Role) Value() (driver.Value, error) {
	if string(s) == "" {
		return nil, nil
	}
	return string(s), nil
}
