package sam

type concatIter[T any] struct {
	its []Iterator[T]
}

func Concat[T any](its ...Iterator[T]) Iterator[T] {
	return &concatIter[T]{
		its: its,
	}
}

func (this *concatIter[T]) Next() (T, bool) {
	if len(this.its) == 0 {
		var empty T
		return empty, false
	}

	ele, ok := this.its[0].Next()
	if ok {
		return ele, ok
	} else {
		this.its = this.its[1:]
		return this.Next()
	}
}
