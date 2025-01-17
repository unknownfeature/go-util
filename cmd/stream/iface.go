package stream

import "github.com/unknownfeature/go-util/cmd/funcs"

type Builder[T any] interface {
	Accept(T)
	Add(T) Builder[T]
	Build() Stream[T]
}

type Stream[T any] interface {
	AllMatch(funcs.Predicate[T]) bool
	AnyMatch(funcs.Predicate[T]) bool
	Count() uint64
	IsParallel() bool
	ForEach(funcs.Consumer[T])
	Reduce(T, funcs.BinaryOperator[T]) T
}
