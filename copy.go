package sam

type copyIter[T any] struct {
	wrap Iterator[*T]
}

func Copy[T any](wrap Iterator[*T]) Iterator[T] {
	return &copyIter[T]{
		wrap: wrap,
	}
}

func (this *copyIter[T]) Next() (T, bool) {
	ele, ok := this.wrap.Next()
	if ok {
		return *ele, true
	} else {
		var empty T
		return empty, false
	}
}
