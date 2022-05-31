package sam

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcat(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, Collect(Copy(Concat(
		NewSliceIter([]int{1, 2}),
		NewSliceIter([]int{3, 4}),
		NewSliceIter([]int{5, 6}),
	))))
}
