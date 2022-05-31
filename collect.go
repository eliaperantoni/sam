package sam

func Collect[T any](it Iterator[T]) []T {
	res := make([]T, 0)
	for ele, ok := it.Next(); ok; ele, ok = it.Next() {
		res = append(res, ele)
	}
	return res
}
