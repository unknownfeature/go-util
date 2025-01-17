package collections

import (
	"github.com/unknownfeature/go-util/cmd/funcs"
	"github.com/unknownfeature/go-util/cmd/stream"
)

type Iterator[T any] interface {
	ForEachRemaining(funcs.Consumer[T])
	HasNext() bool
	Next() T
	Remove()
}
type Collection[T any] interface {
	Add(T) bool
	AddAll(Collection[T]) bool
	Remove(any) bool
	RemoveAll(Collection[T]) bool
	RetainAll(Collection[T]) bool
	Clear()
	Contains(any) bool
	ContainsAll(Collection[any])
	IsEmpty() bool
	Iterator() Iterator[T]
	Size() int
	ToSlice(T[]) []T
	Stream() stream.Stream[T]
	// splititerator todo
}

type Set[T any] interface {
	Collection[T]
}
