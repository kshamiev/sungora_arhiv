package typ

import (
	"database/sql/driver"
	"encoding/hex"

	"github.com/gofrs/uuid"
)

type UUID struct {
	uuid.UUID
}

func UUIDNew() UUID {
	return UUID{UUID: uuid.Must(uuid.NewV4())}
}

func UUIDParse(s string) (UUID, error) {
	u, err := uuid.FromString(s)
	if err != nil {
		return UUID{}, err
	}
	return UUID{u}, nil
}

func UUIDMustParse(s string) UUID {
	u, err := uuid.FromString(s)
	if err != nil {
		return UUID{}
	}
	return UUID{u}
}

func (u UUID) Value() (driver.Value, error) {
	if u.IsNull() {
		return nil, nil
	}
	return u.String(), nil
}

func (u UUID) Bytes() []byte {
	return u.UUID[:]
}

func (u UUID) StringShort() string {
	buf := make([]byte, 32)
	hex.Encode(buf, u.UUID[:])
	return string(buf)
}

func (u UUID) IsNull() bool {
	return u == UUID{}
}

func (u UUID) IsNotNull() bool {
	return u != UUID{}
}

type UUIDS []UUID
