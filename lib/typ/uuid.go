package typ

import (
	"database/sql/driver"

	"github.com/google/uuid"
)

type UUID struct {
	uuid.UUID
}

func UUIDNew() UUID {
	return UUID{UUID: uuid.Must(uuid.NewRandom())}
}

func UUIDParse(s string) (UUID, error) {
	u, err := uuid.Parse(s)
	if err != nil {
		return UUID{}, err
	}
	return UUID{u}, nil
}

// return UUID{uuid.MustParse(s)}
func UUIDMustParse(s string) UUID {
	u, err := uuid.Parse(s)
	if err != nil {
		return UUID{}
	}
	return UUID{u}
}

func (u UUID) Value() (driver.Value, error) {
	if u.ID() == 0 {
		return nil, nil
	}
	return u.String(), nil
}

func (u UUID) Bytes() []byte {
	return u.UUID[:]
}

func (u UUID) IsNull() bool {
	return u == UUID{}
}

func (u UUID) IsNotNull() bool {
	return u != UUID{}
}

type UUIDS []UUID
