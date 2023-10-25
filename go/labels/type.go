package labels

type Label struct {
	name      string
	value     any
}

func NewLabel(name string, value any) Label {
	return Label{name: name, value: value}
}

func name[T any](name string) func(T) Label {
	return func(value T) Label {
		return NewLabel(name, value)
	}
}

func (l Label) Name() string {
	return l.name
}

func (l Label) Value() any {
	return l.value
}
