package optional

type Optional[T any] struct {
	value *T
}

func Option[T any](data T) Optional[T] {
	return Optional[T]{value: &data}
}

func None[T any]() Optional[T] {
	return Optional[T]{value: nil}
}

func (o Optional[T]) Exists() bool {
	return o.value != nil
}

func (o Optional[T]) Get() T {
	if !o.Exists() {
		panic("calling Get() on empty Optional")
	}
	return *o.value
}

func (o Optional[T]) OrElse(other T) T {
	if o.Exists() {
		return *o.value
	} else {
		return other
	}
}
