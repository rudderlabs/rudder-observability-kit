package labels

type Label[T any] struct {
	name      string
	value     T
}

func NewLabel[T any](name string, value T) Label[T] {
	return Label[T]{name: name, value: value}
}

func name[T any](name string) func(T) Label[T] {
	return func(value T) Label[T] {
		return NewLabel[T](name, value)
	}
}

func (l Label[T]) Name() string {
	return l.name
}

func (l Label[T]) Value() T {
	return l.value
}
