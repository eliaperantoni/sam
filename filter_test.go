package sam

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilter(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.Equal(t, []int{2, 4, 6, 8, 10}, Collect(Filter(Copy(NewSliceIter(slice)), func(ele int) bool {
		return ele%2 == 0
	})))
}
