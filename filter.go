package sam

type filterIter[T any] struct {
	wrap Iterator[T]
	fn   func(T) bool
}

func Filter[T any](wrap Iterator[T], fn func(T) bool) Iterator[T] {
	return &filterIter[T]{
		wrap: wrap,
		fn:   fn,
	}
}

func (this *filterIter[T]) Next() (T, bool) {
	for {
		ele, ok := this.wrap.Next()
		if !ok {
			return ele, false
		}

		if this.fn(ele) {
			return ele, true
		}
	}
}
