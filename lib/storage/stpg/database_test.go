package stpg

import (
	"context"
	"database/sql"
	"sync"
	"testing"
	"time"

	"github.com/volatiletech/null/v8"
	"sungora/lib/errs"
	"sungora/lib/storage"
)

type User struct {
	ID          int64       `db:"id" json:"id"`
	CreatedAt   time.Time   `db:"created_at" json:"created_at"`
	Login       string      `db:"login" json:"login"`
	Description null.String `db:"description" json:"description"`
}

var cntGo = 90
var cntIteration = 100

func TestPG(t *testing.T) {
	var cfg = struct {
		Postgresql Config `yaml:"psql"`
	}{
		Postgresql: getConfig(),
	}
	if err := InitConnect(&cfg.Postgresql); err != nil {
		t.Fatal(err)
	}

	channelInsertUpdate := testInsertUpdate(t, Gist())

	var flag = 1
	for 0 < flag {
		select {
		case <-channelInsertUpdate:
			flag--
		}
	}
}

func testInsertUpdate(t *testing.T, st *Storage) chan bool {
	var channelExit = make(chan bool)

	go func() {
		var wg sync.WaitGroup
		for i := 0; i < cntGo; i++ {
			wg.Add(1)
			go func() {
				for j := 0; j < cntIteration; j++ {
					if err := st.QueryTx(context.TODO(), func(qu storage.QueryTxEr) error {
						// INSERT
						arg := []interface{}{
							GenString(8),
							GenString(8),
						}
						id, err := qu.ExecInsert(SQL_USER_INSERT, arg...)
						if err != nil {
							return errs.NewBadRequest(err)
						}
						// UPDATE
						arg = []interface{}{
							GenString(16),
							GenString(16),
							id,
						}
						_, err = qu.Exec(SQL_USER_UPDATE, arg...)
						if err != nil {
							return errs.NewBadRequest(err)
						}
						// UPSERT
						id = 1999999999
						arg = []interface{}{
							id,
							GenString(3),
							GenString(3),
						}
						if _, err = qu.Exec(SQL_USER_UPSERT, arg...); err != nil {
							return errs.NewBadRequest(err)
						}
						arg = []interface{}{
							id,
							GenString(16),
							GenString(16),
						}
						if _, err = qu.Exec(SQL_USER_UPSERT, arg...); err != nil {
							return errs.NewBadRequest(err)
						}
						return nil
					}); err != nil {
						t.Log(err)
					}
				}
				wg.Done()
			}()
		}
		t.Logf("InsertUpdate: Запущено %d паралельных программ по %d итераций в каждой\n", cntGo, cntIteration)
		wg.Wait()
		t.Logf(`InsertUpdate: Done`)
		channelExit <- true
	}()

	return channelExit
}

func TestPGQuery(t *testing.T) {
	var cfg = struct {
		Postgresql Config `yaml:"psql"`
	}{
		Postgresql: getConfig(),
	}
	if err := InitConnect(&cfg.Postgresql); err != nil {
		t.Fatal(err)
	}
	st := Gist()

	// GET Object SLICE 1
	var resList []User
	login := []string{"zLVtPW2i", "gpscdIEk", "rV4VGiR9"}
	if err := st.Query(context.TODO()).Select(&resList, SQL_USER_IN, "JaJOTZvl", login, "v3iwypkK"); err != nil {
		t.Fatal(err)
	}
	t.Log(len(resList))

	// GET Object SLICE 2
	var resListLimit []User
	if err := st.Query(context.TODO()).Select(&resListLimit, SQL_USER_LIMIT); err != nil {
		t.Fatal(err)
	}
	t.Log(len(resListLimit))

	// GET Object SLICE 3
	var resListUsers []User
	if err := st.Query(context.TODO()).Select(&resListUsers, SQL_USER, "testLogin"); err != nil {
		t.Fatal(err)
	}
	t.Log(len(resListUsers))

	// GET Object ONE
	var res User
	if err := st.Query(context.TODO()).Get(&res, SQL_USER, "testLogin"); err != nil && err != sql.ErrNoRows {
		t.Fatal(err)
	}
	t.Log("ID:", res.ID)

	// GET SLICE
	resSlice, err := st.Query(context.TODO()).QuerySlice(SQL_USER, "testLogin")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(resSlice))

	// GET MAP
	resMap, err := st.Query(context.TODO()).QueryMap(SQL_USER, "testLogin")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(resMap))

	// //// Exec TX
	if err := st.QueryTx(context.TODO(), func(qu storage.QueryTxEr) error {

		// INSERT
		arg := []interface{}{
			"Vasya Pupkin",
			"mama@mila.ramu",
		}
		id, err := qu.ExecInsert(SQL_USER_INSERT, arg...)
		if err != nil {
			return err
		}

		// UPDATE
		arg = []interface{}{
			"Vanya Sidorov",
			"popcorn@popcorn.popcorn",
			id,
		}
		_, err = qu.Exec(SQL_USER_UPDATE, arg...)
		if err != nil {
			return err
		}

		// UPSERT
		id = 1888888888
		arg = []interface{}{
			id,
			"1111111111",
			"1111111111",
		}
		if _, err = qu.Exec(SQL_USER_UPSERT, arg...); err != nil {
			return err
		}
		arg = []interface{}{
			id,
			"2222222222222",
			"2222222222222",
		}
		if _, err = qu.Exec(SQL_USER_UPSERT, arg...); err != nil {
			return err
		}

		return nil

	}); err != nil {
		t.Fatal(err)
	}
}
