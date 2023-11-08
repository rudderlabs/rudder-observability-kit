package labels

import "go.uber.org/zap"

type Label[T any] struct {
	name  string
	value T
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

// Expand can be used with the go-kit sugared logger
// e.g. logger.Infow("my message", labels.Expand(l1, l2, l3)...)
func Expand(labels ...Label[any]) []any {
	result := make([]any, 0, len(labels)*2)
	for _, label := range labels {
		result = append(result, label.name, label.value)
	}
	return result
}

// Zap can be used with the go-kit non-sugared logger
// e.g. logger.Infon("my message", l1.Zap(), l2.Zap(), l3.Zap())
func (l Label[T]) Zap() zap.Field {
	switch v := any(l.value).(type) {
	case string:
		return zap.String(l.name, v)
	case bool:
		return zap.Bool(l.name, v)
	case int:
		return zap.Int(l.name, v)
	case float32:
		return zap.Float32(l.name, v)
	case float64:
		return zap.Float64(l.name, v)
	case error:
		return zap.Error(v)
	case int32:
		return zap.Int32(l.name, v)
	case int64:
		return zap.Int64(l.name, v)
	case uint:
		return zap.Uint(l.name, v)
	case uint32:
		return zap.Uint32(l.name, v)
	case uint64:
		return zap.Uint64(l.name, v)
	case []byte:
		return zap.Binary(l.name, v)
	default:
		return zap.Any(l.name, v)
	}
}
