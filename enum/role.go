package enum

import "database/sql/driver"

type Role string

const (
	Role_DEVELOP   Role = "DEVELOP"
	Role_ADMIN     Role = "ADMIN"
	Role_GUEST     Role = "GUEST"
	Role_MODERATOR Role = "MODERATOR"
)

var RoleName = map[string]Role{
	Role_DEVELOP.String():   Role_DEVELOP,
	Role_ADMIN.String():     Role_ADMIN,
	Role_GUEST.String():     Role_GUEST,
	Role_MODERATOR.String(): Role_MODERATOR,
}

var RoleValue = map[Role]string{
	Role_DEVELOP:   Role_DEVELOP.String(),
	Role_ADMIN:     Role_ADMIN.String(),
	Role_GUEST:     Role_GUEST.String(),
	Role_MODERATOR: Role_MODERATOR.String(),
}

func (s Role) Valid() bool {
	_, ok := RoleValue[s]
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
