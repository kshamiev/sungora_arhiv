// nolint: lll // AFTER CODE GENERATED. DO NOT EDIT //
package typ

import "sungora/lib/uuid"

//
type Roles struct {
	Id          uuid.UUID `json:"id" db:"id" example:"ca6f30f9-7207-4741-8dba-7f288edf1161"` //
	Code        string    `json:"code" db:"code"`                                            //
	Description string    `json:"description" db:"description"`                              //
}

func (o *Roles) Select() (query string, args []interface{}) {
	const SQLRolesSelect = "SELECT id, code, description FROM public.roles WHERE id = $1"
	return SQLRolesSelect, []interface{}{
		o.Id,
	}
}

func (o *Roles) Insert() (query string, args []interface{}) {
	const SQLRolesInsert = "INSERT INTO public.roles (id, code, description) VALUES ($1, $2, $3)"
	return SQLRolesInsert, []interface{}{
		o.Id, o.Code, o.Description,
	}
}

func (o *Roles) Update() (query string, args []interface{}) {
	const SQLRolesUpdate = "UPDATE public.roles SET id = $1, code = $2, description = $3 WHERE id = $1"
	return SQLRolesUpdate, []interface{}{
		o.Id, o.Code, o.Description,
	}
}

func (o *Roles) Upsert() (query string, args []interface{}) {
	const SQLRolesUpsert = "INSERT INTO public.roles (id, code, description) VALUES ($1, $2, $3) ON CONFLICT (id) DO UPDATE SET id = $1, code = $2, description = $3"
	return SQLRolesUpsert, []interface{}{
		o.Id, o.Code, o.Description,
	}
}

func (o *Roles) Delete() (query string, args []interface{}) {
	const SQLRolesDelete = "DELETE FROM public.roles WHERE id = $1"
	return SQLRolesDelete, []interface{}{
		o.Id,
	}
}

// BEFORE CODE GENERATED. DO NOT EDIT //
