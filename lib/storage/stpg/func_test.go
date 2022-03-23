package stpg

import (
	"crypto/rand"
	"io"
	"testing"
)

// go test -bench=. -benchmem -benchtime=1000000x
// Benchmark_sqlIn-8        1000000              2951 ns/op            1625 B/op         37 allocs/op
// Benchmark_sqlIn-8        1000000              1171 ns/op             439 B/op         11 allocs/op
func Benchmark_sqlIn(b *testing.B) {
	sql := "SELECT * FROM table WHERE filed1 = $1 AND filed2 IN($2) AND filed3 = $3"
	// sql := "SELECT * FROM table WHERE filed1 = $1 AND filed2 = $2"
	in := []string{"val1", "val2", "val3", "val4", "val5", "val6", "val7", "val8", "val9"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, err := sqlIn(sql, 23, in, "popcorn")
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

const (
	num     = "0123456789"
	strdown = "abcdefghijklmnopqrstuvwxyz"
	strup   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func GenString(length int) string {
	return randChar(length, []byte(strdown+strup+num))
}

func randChar(length int, chars []byte) string {
	pword := make([]byte, length)
	data := make([]byte, length+(length/4)) // storage for random bytes.
	clen := byte(len(chars))
	maxrb := byte(256 - (256 % len(chars)))
	i := 0

	for {
		if _, err := io.ReadFull(rand.Reader, data); err != nil {
			panic(err)
		}

		for _, c := range data {
			if c >= maxrb {
				continue
			}

			pword[i] = chars[c%clen]
			i++

			if i == length {
				return string(pword)
			}
		}
	}
}
