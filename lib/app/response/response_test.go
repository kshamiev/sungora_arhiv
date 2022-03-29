package response

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// go test -bench=. -benchmem -benchtime=1000000x > new.txt
// Benchmark_generalHeaderSet-8     1000000              1163 ns/op             208 B/op         11 allocs/op
// Benchmark_generalHeaderSet-8     1000000               310.1 ns/op            69 B/op          4 allocs/op
func Benchmark_generalHeaderSet(b *testing.B) {
	r := httptest.NewRequest("GET", "/v1/leftpad/?str=test&len=50&chr=*", nil)
	w := httptest.NewRecorder()
	rw := New(r, w)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rw.generalHeaderSet("test.json", 23450, http.StatusOK)
	}
}
