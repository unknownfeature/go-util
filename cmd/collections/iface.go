package collections

import "github.com/unknownfeature/go-util/cmd/funcs"

type Iterator[T any] interface {
	ForEachRemaining(funcs.Consumer[T])
	HasNext() bool
	Next() T
	Remove()
}
type Collection[T any] interface {
	Add(T) bool
	AddAll(Collection[T])
	Clear()
	Contains(any)
	ContainsAll(Collection[any])
	IsEmpty()
	Iterator() Iterator[T]
}
