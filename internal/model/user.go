package model

import (
	"context"
	"time"

	"sample/internal/body"
	"sample/lib/errs"
	"sample/lib/jaeger"
	"sample/lib/storage"
	"sample/lib/storage/stpg"
	"sample/services/mdsample"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type User struct {
	st storage.Face
}

func NewUser() *User {
	return &User{st: stpg.Gist()}
}

func (mm *User) Load(ctx context.Context, id int64) (*mdsample.User, error) {
	s := jaeger.NewSpan(ctx)
	s.StringAttribute("param1", "fantik")
	s.Int64Attribute("param2", 34)
	s.Float64Attribute("param3", 45.76)
	s.BoolAttribute("param4", true)
	defer s.End()

	us := &mdsample.User{}

	// sqlx
	if err := mm.st.DB().GetContext(ctx, us, "SELECT * FROM users WHERE id = $1", id); err != nil {
		return nil, errs.New(err, body.ErrUserTwo, id)
	}

	// boiler
	if err := us.Reload(ctx, mm.st.DB()); err != nil {
		return nil, errs.New(err)
	}

	// custom from sql
	if err := mm.st.Query(ctx).Get(us, "SELECT * FROM users WHERE id = $1", id); err != nil {
		return nil, err
	}

	// custom from sql (transaction)
	if err := mm.st.QueryTx(ctx, func(qu storage.QueryTxEr) error {
		if err := qu.Get(us, "SELECT * FROM users WHERE id = $1", id); err != nil {
			return errs.New(err)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	us.Duration = time.Hour + time.Minute*10 + time.Second*10
	if _, err := us.Update(ctx, mm.st.DB(), boil.Infer()); err != nil {
		return nil, errs.New(err)
	}

	return us, nil
}
