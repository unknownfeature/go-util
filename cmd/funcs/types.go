package funcs

type Supplier[T any] func() T
type Consumer[T any] func(T)
type Function[T, K any] func(T) K
type BiFunction[T, K, L any] func(T, K) L
type Predicate[T any] Function[T, bool]
type BiPredicate[T, K any] BiFunction[T, K, bool]
type BiConsumer[T, K any] func(T, K)
type Comparator[T any] BiFunction[T, T, int]
type BinaryOperator[T any] BiFunction[T, T, T]

func MaxBy[T any](comparator Comparator[T]) BinaryOperator[T] {
	return func(t1 T, t2 T) T {
		compRes := comparator(t1, t2)
		if compRes >= 0 {
			return t1
		}
		return t2
	}
}

func MinBy[T any](comparator Comparator[T]) BinaryOperator[T] {
	return func(t1 T, t2 T) T {
		compRes := comparator(t1, t2)
		if compRes < 0 {
			return t1
		}
		return t2
	}
}
