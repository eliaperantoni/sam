package sam

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type strPtr struct {
	ptr *string
}

func (this strPtr) Clone() strPtr {
	cpy := *this.ptr
	return strPtr{
		ptr: &cpy,
	}
}

func allocStr(str string) *string {
	return &str
}

func TestClone(t *testing.T) {
	slice := []strPtr{
		{ptr: allocStr("This a test")},
	}

	copied := Collect(Copy(NewSliceIter(slice)))[0]
	cloned := Collect(Clone(Copy(NewSliceIter(slice))))[0]

	assert.True(t, slice[0].ptr == copied.ptr)
	assert.False(t, slice[0].ptr == cloned.ptr)
}
