package mathx

import (
	"fmt"
	"net/http"
	"testing"
)

func BenchmarkFib(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		generateLoad(b)
	}
}

func generateLoad(b *testing.B) {
	uri := fmt.Sprintf("http://localhost:6060/fib?n=%d", 35)
	r, err := http.DefaultClient.Get(uri)
	if err != nil {
		b.Log(err.Error())
		b.FailNow()
	}
	r.Body.Close()
}
