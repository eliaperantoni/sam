package sam

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlice(t *testing.T) {
	slice := []int{1, 2, 3}
	assert.Equal(t, []*int{&slice[0], &slice[1], &slice[2]}, Collect(NewSliceIter(slice)))
}
