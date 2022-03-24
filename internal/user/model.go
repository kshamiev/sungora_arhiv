package user

import (
	"context"
	"time"

	"sungora/lib/errs"
	"sungora/lib/jaeger"
	"sungora/lib/storage"
	"sungora/services/mdsungora"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Model struct {
	st storage.Face
}

func NewModel(st storage.Face) *Model {
	return &Model{st}
}

func (mm *Model) Load(ctx context.Context, id int64) (*mdsungora.User, error) {
	s := jaeger.NewSpan(ctx)
	s.StringAttribute("param1", "fantik")
	s.Int64Attribute("param2", 34)
	s.Float64Attribute("param3", 45.76)
	s.BoolAttribute("param4", true)
	defer s.End()

	us := &mdsungora.User{}

	// sqlx
	if err := mm.st.DB().GetContext(ctx, us, "SELECT * FROM users WHERE id = $1", id); err != nil {
		return nil, errs.New(err, ErrUserTwo, id)
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
