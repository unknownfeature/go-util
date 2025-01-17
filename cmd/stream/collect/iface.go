package collect

import (
	"github.com/unknownfeature/go-util/cmd/collections"
	"github.com/unknownfeature/go-util/cmd/funcs"
)

type Characteristics int

const (
	Concurrent Characteristics = iota
	Unordered
	IdentityFinish
)

type Collector[T, A, R any] interface {
	Accumulator() funcs.BiConsumer[A, T]
	Combiner() funcs.BinaryOperator[A]
	Finisher() funcs.Function[A, R]
	Supplier() funcs.Supplier[A]
	Characteristics() collections.Set[Characteristics]
}
