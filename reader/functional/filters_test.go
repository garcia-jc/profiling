package functional

import (
	"io"
	"testing"
)

func BenchmarkFilters(b *testing.B) {
	items := ReadItems("../large.json")
	public, private := 0, 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Work(items, io.Discard)
	}
	_, _ = public, private
}
