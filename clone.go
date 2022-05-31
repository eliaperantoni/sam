package sam

// Cloneable [T] is implemented by type T if it can clone itself, producing another instance of T
type Cloneable[T any] interface {
	Clone() T
}

type cloneIter[T Cloneable[T]] struct {
	wrap Iterator[T]
}

func Clone[T Cloneable[T]](wrap Iterator[T]) Iterator[T] {
	return &cloneIter[T]{
		wrap: wrap,
	}
}

func (this *cloneIter[T]) Next() (T, bool) {
	ele, ok := this.wrap.Next()
	if !ok {
		return ele, false
	}

	return ele.Clone(), true
}
