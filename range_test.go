package sam

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRange(t *testing.T) {
	slice, i := []int{1, 2, 3}, 0
	for ele := range Range(Copy(NewSliceIter(slice))) {
		assert.Equal(t, slice[i], ele)
		i++
	}
}
