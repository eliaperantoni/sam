package sam

type mapIter[T, U any] struct {
	wrap Iterator[T]
	fn   func(T) U
}

func Map[T, U any](wrap Iterator[T], fn func(T) U) Iterator[U] {
	return &mapIter[T, U]{
		wrap: wrap,
		fn:   fn,
	}
}

func (this *mapIter[T, U]) Next() (U, bool) {
	ele, ok := this.wrap.Next()
	if ok {
		return this.fn(ele), true
	} else {
		var empty U
		return empty, false
	}
}
