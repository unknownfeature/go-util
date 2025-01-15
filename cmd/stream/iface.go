package stream

import "github.com/unknownfeature/go-util/cmd/funcs"

type Builder[T any] interface {
	Accept(T)
	Add(T) Builder[T]
	Build() Stream[T]
}

type Collector[T, A, R any] interface {
	// characteristics
	Accumulator() funcs.BiConsumer[A, T]
	Combiner() funcs.BinaryOperator[A]
	Finisher() funcs.Function[A, R]
	Supplier() funcs.Supplier[A]
}

type Stream[T any] interface {
	AllMatch(predicate funcs.Predicate[T]) bool
	AnyMatch(predicate funcs.Predicate[T]) bool
	//Collect()
}

//func Collector[R, T any](funcs.Supplier[R], funcs.BiConsumer[R, T], funcs.BiConsumer[R, R]) {
//
//}
