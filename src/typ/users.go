// nolint: lll // AFTER CODE GENERATED. DO NOT EDIT //
package typ

import (
	"time"

	"github.com/shopspring/decimal"

	"sungora/lib/null"
	"sungora/lib/uuid"
)

//
type Users struct {
	Id        uuid.UUID       `json:"id" db:"id" example:"ca6f30f9-7207-4741-8dba-7f288edf1161"` //
	Login     string          `json:"login" db:"login"`                                          //
	Email     string          `json:"email" db:"email"`                                          //
	Price     decimal.Decimal `json:"price" db:"price" example:"0.1"`                            //
	SummaOne  float32         `json:"summa_one" db:"summa_one" example:"0.1"`                    //
	SummaTwo  float64         `json:"summa_two" db:"summa_two" example:"0.1"`                    //
	Cnt2      int             `json:"cnt2" db:"cnt2"`                                            //
	Cnt4      int             `json:"cnt4" db:"cnt4"`                                            //
	Cnt8      int64           `json:"cnt8" db:"cnt8"`                                            //
	IsOnline  bool            `json:"is_online" db:"is_online"`                                  //
	Metrika   null.JSON       `json:"metrika" db:"metrika" swaggertype:"string"`                 //
	CreatedAt time.Time       `json:"created_at" db:"created_at" example:"2006-01-02T15:04:05Z"` //
	UpdatedAt time.Time       `json:"updated_at" db:"updated_at" example:"2006-01-02T15:04:05Z"` //
	DeletedAt null.Time       `json:"deleted_at" db:"deleted_at" example:"2006-01-02T15:04:05Z"` //
	Duration  time.Duration   `json:"duration" db:"duration" swaggertype:"string" example:"5m"`  //
}

func (o *Users) Select() (query string, args []interface{}) {
	const SQLUsersSelect = "SELECT id, login, email, price, summa_one, summa_two, cnt2, cnt4, cnt8, is_online, metrika, created_at, updated_at, deleted_at, duration FROM public.users WHERE id = $1"
	return SQLUsersSelect, []interface{}{
		o.Id,
	}
}

func (o *Users) Insert() (query string, args []interface{}) {
	const SQLUsersInsert = "INSERT INTO public.users (id, login, email, price, summa_one, summa_two, cnt2, cnt4, cnt8, is_online, metrika, created_at, updated_at, deleted_at, duration) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)"
	return SQLUsersInsert, []interface{}{
		o.Id, o.Login, o.Email, o.Price, o.SummaOne, o.SummaTwo, o.Cnt2, o.Cnt4, o.Cnt8, o.IsOnline, o.Metrika, o.CreatedAt, o.UpdatedAt, o.DeletedAt, o.Duration,
	}
}

func (o *Users) Update() (query string, args []interface{}) {
	const SQLUsersUpdate = "UPDATE public.users SET id = $1, login = $2, email = $3, price = $4, summa_one = $5, summa_two = $6, cnt2 = $7, cnt4 = $8, cnt8 = $9, is_online = $10, metrika = $11, created_at = $12, updated_at = $13, deleted_at = $14, duration = $15 WHERE id = $1"
	return SQLUsersUpdate, []interface{}{
		o.Id, o.Login, o.Email, o.Price, o.SummaOne, o.SummaTwo, o.Cnt2, o.Cnt4, o.Cnt8, o.IsOnline, o.Metrika, o.CreatedAt, o.UpdatedAt, o.DeletedAt, o.Duration,
	}
}

func (o *Users) Upsert() (query string, args []interface{}) {
	const SQLUsersUpsert = "INSERT INTO public.users (id, login, email, price, summa_one, summa_two, cnt2, cnt4, cnt8, is_online, metrika, created_at, updated_at, deleted_at, duration) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) ON CONFLICT (id) DO UPDATE SET id = $1, login = $2, email = $3, price = $4, summa_one = $5, summa_two = $6, cnt2 = $7, cnt4 = $8, cnt8 = $9, is_online = $10, metrika = $11, created_at = $12, updated_at = $13, deleted_at = $14, duration = $15"
	return SQLUsersUpsert, []interface{}{
		o.Id, o.Login, o.Email, o.Price, o.SummaOne, o.SummaTwo, o.Cnt2, o.Cnt4, o.Cnt8, o.IsOnline, o.Metrika, o.CreatedAt, o.UpdatedAt, o.DeletedAt, o.Duration,
	}
}

func (o *Users) Delete() (query string, args []interface{}) {
	const SQLUsersDelete = "DELETE FROM public.users WHERE id = $1"
	return SQLUsersDelete, []interface{}{
		o.Id,
	}
}

// BEFORE CODE GENERATED. DO NOT EDIT //
