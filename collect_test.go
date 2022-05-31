package sam

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCollect(t *testing.T) {
	slice := []int{1, 2, 3}
	assert.Equal(t, slice, Collect(Copy(NewSliceIter(slice))))
}
