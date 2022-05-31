package sam

type sliceIter[T any] struct {
	wrap  []T
	index int
}

func NewSliceIter[T any](slice []T) Iterator[*T] {
	return &sliceIter[T]{
		wrap:  slice,
		index: 0,
	}
}

func (this *sliceIter[T]) Next() (*T, bool) {
	if this.index < len(this.wrap) {
		ele := &this.wrap[this.index]
		this.index++
		return ele, true
	} else {
		return nil, false
	}
}
