package sam

type Iterator[T any] interface {
	Next() (T, bool)
}
