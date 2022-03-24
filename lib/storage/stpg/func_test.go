package stpg

import (
	"crypto/rand"
	"io"
)

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

func genString(length int) string {
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
