package stpg

import (
	"testing"
)

// go test -bench=. -benchmem -benchtime=1000000x
// Benchmark_sqlIn-8        1000000              2951 ns/op            1625 B/op         37 allocs/op
func Benchmark_sqlIn(b *testing.B) {
	sql := "SELECT * FROM table WHERE filed1 = $1 AND filed2 IN($2) AND filed3 = $3"
	in := []string{"val1", "val2", "val3", "val4", "val5", "val6", "val7", "val8", "val9"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, err := sqlIn(sql, 23, in, "popcorn")
		if err != nil {
			b.Fatal(err)
		}
	}
}
