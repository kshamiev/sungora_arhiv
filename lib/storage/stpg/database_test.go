package stpg

import (
	"context"
	"database/sql"
	"sync"
	"testing"
	"time"

	"sungora/lib/storage"

	"github.com/volatiletech/null/v8"
)

type User struct {
	ID          int64       `db:"id" json:"id"`
	CreatedAt   time.Time   `db:"created_at" json:"created_at"`
	Login       string      `db:"login" json:"login"`
	Description null.String `db:"description" json:"description"`
}

// запросы вне транзакции
func TestPGQuery(t *testing.T) {
	var cfg = struct {
		Postgresql Config `yaml:"psql"`
	}{
		Postgresql: getConfig(),
	}
	if err := InitConnect(&cfg.Postgresql); err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	qu := Gist().Query(ctx)

	// GET one Object
	var res User
	if err := qu.Get(&res, SQL_USER, "testLogin"); err != nil && err != sql.ErrNoRows {
		t.Fatal(err)
	}
	t.Log("ID:", res.ID)

	// GET more Object
	// важно: в этом типе запроса ошибки sql.ErrNoRows не бывает
	// как для получение одного обьекта (что считается нормально)
	var resList []User
	login := []string{"zLVtPW2i", "gpscdIEk", "rV4VGiR9"}
	if err := qu.Select(&resList, SQL_USER_IN, "JaJOTZvl", login, "v3iwypkK"); err != nil {
		t.Fatal(err)
	}
	t.Log(len(resList))

	// GET more slice
	resSlice, err := qu.QuerySlice(SQL_USER, "testLogin")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(resSlice))

	// GET more map
	resMap, err := qu.QueryMap(SQL_USER, "testLogin")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(resMap))

	// INSERT
	arg := []interface{}{
		"Vasya Pupkin",
		"mama@mila.ramu",
	}
	id, err := qu.ExecInsert(SQL_USER_INSERT, arg...)
	if err != nil {
		t.Fatal(err)
	}

	// UPDATE valid DELETE
	arg = []interface{}{
		"Vanya Sidorov",
		"popcorn@popcorn.popcorn",
		id,
	}
	_, err = qu.Exec(SQL_USER_UPDATE, arg...)
	if err != nil {
		t.Fatal(err)
	}

	// UPSERT
	id = 1777777777
	arg = []interface{}{
		id,
		"1111111111",
		"1111111111",
	}
	if _, err = qu.Exec(SQL_USER_UPSERT, arg...); err != nil {
		t.Fatal(err)
	}
	arg = []interface{}{
		id,
		"2222222222222",
		"2222222222222",
	}
	if _, err = qu.Exec(SQL_USER_UPSERT, arg...); err != nil {
		t.Fatal(err)
	}

	// sample classic single query
	// _, err = Gist().Query(ctx).Exec("DELETE FROM Users WHERE id = $1", 34)
}

// запросы в транзакции
func TestPGQueryTx(t *testing.T) {
	var cfg = struct {
		Postgresql Config `yaml:"psql"`
	}{
		Postgresql: getConfig(),
	}
	if err := InitConnect(&cfg.Postgresql); err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	st := Gist()

	// все запросы в блоке выполняются в транзакции
	if err := st.QueryTx(ctx, func(qu storage.QueryTxEr) error {
		// SELECT запросы аналагичны выше описанным (TestPGQuery)

		// INSERT
		arg := []interface{}{
			"Vasya Pupkin",
			"mama@mila.ramu",
		}
		id, err := qu.ExecInsert(SQL_USER_INSERT, arg...)
		if err != nil {
			return err
		}

		// UPDATE valid DELETE
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

// нагрузочные тесты

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

	<-testInsertUpdate(t)
}

func testInsertUpdate(t *testing.T) chan bool {
	channelExit := make(chan bool)
	st := Gist()

	go func() {
		var wg sync.WaitGroup
		for i := 0; i < cntGo; i++ {
			wg.Add(1)
			go func() {
				for j := 0; j < cntIteration; j++ {
					if err := st.QueryTx(context.TODO(), func(qu storage.QueryTxEr) error {
						// INSERT
						arg := []interface{}{
							storage.GenString(8),
							storage.GenString(8),
						}
						id, err := qu.ExecInsert(SQL_USER_INSERT, arg...)
						if err != nil {
							return err
						}
						// UPDATE
						arg = []interface{}{
							storage.GenString(16),
							storage.GenString(16),
							id,
						}
						_, err = qu.Exec(SQL_USER_UPDATE, arg...)
						if err != nil {
							return err
						}
						// UPSERT
						id = 1999999999
						arg = []interface{}{
							id,
							storage.GenString(3),
							storage.GenString(3),
						}
						if _, err = qu.Exec(SQL_USER_UPSERT, arg...); err != nil {
							return err
						}
						arg = []interface{}{
							id,
							storage.GenString(16),
							storage.GenString(16),
						}
						if _, err = qu.Exec(SQL_USER_UPSERT, arg...); err != nil {
							return err
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

// go test -bench=. -benchmem -benchtime=1000000x
// Benchmark_sqlIn-8        1000000              2951 ns/op            1625 B/op         37 allocs/op
// Benchmark_sqlIn-8        1000000              1171 ns/op             439 B/op         11 allocs/op
func Benchmark_sqlIn(b *testing.B) {
	s := "SELECT * FROM table WHERE filed1 = $1 AND filed2 IN($2) AND filed3 = $3"
	// sql := "SELECT * FROM table WHERE filed1 = $1 AND filed2 = $2"
	in := []string{"val1", "val2", "val3", "val4", "val5", "val6", "val7", "val8", "val9"}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _, err := sqlIn(s, 23, in, "popcorn")
		if err != nil {
			b.Fatal(err)
		}
	}
}

func getConfig() Config {
	return Config{
		Postgres:     "",
		User:         "postgres",
		Pass:         "postgres",
		Host:         "localhost",
		Port:         5432,
		Dbname:       "test",
		Sslmode:      "disable",
		Blacklist:    []string{"test"},
		MaxIdleConns: 50,
		MaxOpenConns: 50,
		OcSQLTrace:   false,
	}
}
