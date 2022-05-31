package sam

import (
	"github.com/thoas/go-funk"
	"testing"
)

const n int = 1000000

func makeTestingSlice() []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = i
	}
	return slice
}

func BenchmarkSam(b *testing.B) {
	slice := makeTestingSlice()
	for i := 0; i < b.N; i++ {
		Collect(Map(Copy(NewSliceIter(slice)), func(ele int) int {
			return ele * 2
		}))
	}
}

func BenchmarkGoFunk(b *testing.B) {
	slice := makeTestingSlice()
	for i := 0; i < b.N; i++ {
		funk.Map(slice, func(ele int) int {
			return ele * 2
		})
	}
}
