package optional

type Optional[T any] struct {
	val *T
}

func (o Optional[T]) Get() T {
	if o.val == nil {
		panic("value not set")
	}
	return *o.val
}

func Of[T any](val T) Optional[T] {
	return Optional[T]{val: &val}
}

func OfNullable[T any](val *T) Optional[T] {
	return Optional[T]{val: val}
}
