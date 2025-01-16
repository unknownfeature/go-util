package stream

import "github.com/unknownfeature/go-util/cmd/funcs"

type Builder[T any] interface {
	Accept(T)
	Add(T) Builder[T]
	Build() Stream[T]
}

type Stream[T any] interface {
	AllMatch(predicate funcs.Predicate[T]) bool
	AnyMatch(predicate funcs.Predicate[T]) bool
	Count() uint64
}
