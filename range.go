package sam

// Range return a channel that ranges over the elements of the iterator and is closed at the end. The channel *must* be
// exhausted otherwise a goroutine is going to leak
func Range[T any](it Iterator[T]) <-chan T {
	c := make(chan T)
	go func() {
		for ele, ok := it.Next(); ok; ele, ok = it.Next() {
			c <- ele
		}
		close(c)
	}()
	return c
}
