package collect

import "github.com/unknownfeature/go-util/cmd/stream"

func Collect[T, A, R any](collector Collector[T, A, R], stream stream.Stream[T]) R {
	container := collector.Supplier()()
	accumulator := collector.Accumulator()
	stream.ForEach(func(t T) {
		accumulator(container, t)
	})
	return collector.Finisher()(container)
	// todo this is just high level impl, a lot is needed to be added (like parallelism)
}
