package model

import (
	"context"

	"sungora/lib/app"
	"sungora/lib/errs"
	"sungora/lib/storage"
	"sungora/lib/storage/pgsql"
	"sungora/lib/typ"
	"sungora/types/mdsun"
)

type User struct {
	st *pgsql.Storage
}

func NewUser(st *pgsql.Storage) *User {
	return &User{st}
}

func (u *User) Load(ctx context.Context, id typ.UUID) (*mdsun.User, error) {
	s := app.NewSpan(ctx)
	s.StringAttribute("param1", "fantik")
	s.Int64Attribute("param2", 34)
	s.Float64Attribute("param3", 45.76)
	s.BoolAttribute("param4", true)
	defer s.End()

	us := &mdsun.User{}

	// sqlx
	if err := u.st.DB().GetContext(ctx, us, "SELECT * FROM users WHERE id = $1", id); err != nil {
		return nil, errs.NewBadRequest(err)
	}

	// boiler
	if err := us.Reload(ctx, u.st.DB()); err != nil {
		return nil, errs.NewBadRequest(err)
	}

	// custom from sql
	if err := u.st.Query(ctx).Get(us, "SELECT * FROM users WHERE id = $1", id); err != nil {
		return nil, err
	}

	if err := u.st.QueryTx(ctx, func(qu storage.QueryTxEr) error {
		if err := qu.Get(us, "SELECT * FROM users WHERE id = $1", id); err != nil {
			return errs.NewBadRequest(err)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return us, nil
}
