package collect

import (
	"github.com/unknownfeature/go-util/cmd/funcs"
)

type Collector[T, A, R any] interface {
	Accumulator() funcs.BiConsumer[A, T]
	Combiner() funcs.BinaryOperator[A]
	Finisher() funcs.Function[A, R]
	Supplier() funcs.Supplier[A]
}
