package user

import (
	"context"
	"time"

	"sungora/lib/app"
	"sungora/lib/errs"
	"sungora/lib/storage"
	"sungora/lib/storage/pgsql"
	"sungora/lib/typ"
	"sungora/services/mdsample"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Model struct {
	st *pgsql.Storage
}

func NewModel(st *pgsql.Storage) *Model {
	return &Model{st}
}

func (mm *Model) Load(ctx context.Context, id typ.UUID) (*mdsample.User, error) {
	s := app.NewSpan(ctx)
	s.StringAttribute("param1", "fantik")
	s.Int64Attribute("param2", 34)
	s.Float64Attribute("param3", 45.76)
	s.BoolAttribute("param4", true)
	defer s.End()

	us := &mdsample.User{}

	// sqlx
	if err := mm.st.DB().GetContext(ctx, us, "SELECT * FROM users WHERE id = $1", id); err != nil {
		return nil, errs.NewBadRequest(err)
	}

	// boiler
	if err := us.Reload(ctx, mm.st.DB()); err != nil {
		return nil, errs.NewBadRequest(err)
	}

	// custom from sql
	if err := mm.st.Query(ctx).Get(us, "SELECT * FROM users WHERE id = $1", id); err != nil {
		return nil, err
	}

	if err := mm.st.QueryTx(ctx, func(qu storage.QueryTxEr) error {
		if err := qu.Get(us, "SELECT * FROM users WHERE id = $1", id); err != nil {
			return errs.NewBadRequest(err)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	us.Duration = time.Hour + time.Minute*10 + time.Second*10
	if _, err := us.Update(ctx, mm.st.DB(), boil.Infer()); err != nil {
		return nil, errs.NewBadRequest(err)
	}

	return us, nil
}
