package future

import (
	"github.com/unknownfeature/go-util/cmd/funcs"
	"sync"
)

type Future[T any] struct {
	task    func() T
	resChan chan T
	res     T
	done    bool
	lock    *sync.Mutex
}

func SupplyAsync[T any](task funcs.Supplier[T]) *Future[T] {

	f := &Future[T]{task: task, resChan: make(chan T), lock: &sync.Mutex{}}
	go func() {
		res := f.task()
		f.resChan <- res
	}()
	return f
}

func (f *Future[T]) Get() T {
	f.lock.Lock()
	defer f.lock.Unlock()
	if f.done {
		return f.res
	}
	f.res = <-f.resChan
	f.done = true
	return f.res
}
