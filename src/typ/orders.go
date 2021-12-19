// nolint: lll // AFTER CODE GENERATED. DO NOT EDIT //
package typ

import (
	"time"

	"sungora/lib/null"
	"sungora/lib/typ"
)

//
type Orders struct {
	Id        typ.UUID  `json:"id" db:"id" example:"ca6f30f9-7207-4741-8dba-7f288edf1161"`           //
	UserId    typ.UUID  `json:"user_id" db:"user_id" example:"ca6f30f9-7207-4741-8dba-7f288edf1161"` //
	Number    int       `json:"number" db:"number"`                                                  //
	Status    string    `json:"status" db:"status"`                                                  //
	CreatedAt time.Time `json:"created_at" db:"created_at" example:"2006-01-02T15:04:05Z"`           //
	UpdatedAt time.Time `json:"updated_at" db:"updated_at" example:"2006-01-02T15:04:05Z"`           //
	DeletedAt null.Time `json:"deleted_at" db:"deleted_at" example:"2006-01-02T15:04:05Z"`           //
}

func (o *Orders) Select() (query string, args []interface{}) {
	const SQLOrdersSelect = "SELECT id, user_id, number, status, created_at, updated_at, deleted_at FROM public.orders WHERE id = $1"
	return SQLOrdersSelect, []interface{}{
		o.Id,
	}
}

func (o *Orders) Insert() (query string, args []interface{}) {
	const SQLOrdersInsert = "INSERT INTO public.orders (id, user_id, number, status, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	return SQLOrdersInsert, []interface{}{
		o.Id, o.UserId, o.Number, o.Status, o.CreatedAt, o.UpdatedAt, o.DeletedAt,
	}
}

func (o *Orders) Update() (query string, args []interface{}) {
	const SQLOrdersUpdate = "UPDATE public.orders SET id = $1, user_id = $2, number = $3, status = $4, created_at = $5, updated_at = $6, deleted_at = $7 WHERE id = $1"
	return SQLOrdersUpdate, []interface{}{
		o.Id, o.UserId, o.Number, o.Status, o.CreatedAt, o.UpdatedAt, o.DeletedAt,
	}
}

func (o *Orders) Upsert() (query string, args []interface{}) {
	const SQLOrdersUpsert = "INSERT INTO public.orders (id, user_id, number, status, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7) ON CONFLICT (id) DO UPDATE SET id = $1, user_id = $2, number = $3, status = $4, created_at = $5, updated_at = $6, deleted_at = $7"
	return SQLOrdersUpsert, []interface{}{
		o.Id, o.UserId, o.Number, o.Status, o.CreatedAt, o.UpdatedAt, o.DeletedAt,
	}
}

func (o *Orders) Delete() (query string, args []interface{}) {
	const SQLOrdersDelete = "DELETE FROM public.orders WHERE id = $1"
	return SQLOrdersDelete, []interface{}{
		o.Id,
	}
}

// BEFORE CODE GENERATED. DO NOT EDIT //
