package sam

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCopy(t *testing.T) {
	slice := []int{1, 2, 3}
	assert.Equal(t, []int{1, 2, 3}, Collect(Copy(NewSliceIter(slice))))
}
