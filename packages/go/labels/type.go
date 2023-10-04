package labels

type Label struct {
	name  string
	value string
}

func NewLabel(name string, value string) Label {
	return Label{name: name, value: value}
}

func Name(name string) func(string) Label {
	return func(value string) Label {
		return NewLabel(name, value)
	}
}

func (l Label) Name() string {
	return l.name
}

func (l Label) Value() string {
	return l.value
}
