package sam

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestMap(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	assert.Equal(t, []float64{2, 4, 8, 16}, Collect(Map(Copy(NewSliceIter(slice)), func(ele int) float64 {
		return math.Pow(2, float64(ele))
	})))
}
