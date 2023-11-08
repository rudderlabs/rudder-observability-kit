package labels

import (
	"go.uber.org/zap"
	"time"
)

type FieldType uint8

const (
	UnknownType FieldType = iota
	StringType
	IntType
	BoolType
	FloatType
	TimeType
	DurationType
)

type Label struct {
	name      string
	fieldType FieldType

	unknown  any
	string   string
	int      int64
	bool     bool
	float    float64
	time     time.Time
	duration time.Duration
}

type LabelFunc func(any) Label

func NewLabel(name string, value any) Label {
	return Label{name: name, unknown: value}
}

func StringLabel(name, value string) Label {
	return Label{name: name, string: value, fieldType: StringType}
}

func IntLabel(name string, value int64) Label {
	return Label{name: name, int: value, fieldType: IntType}
}

func BoolLabel(name string, value bool) Label {
	return Label{name: name, bool: value, fieldType: BoolType}
}

func FloatLabel(name string, value float64) Label {
	return Label{name: name, float: value, fieldType: FloatType}
}

func TimeLabel(name string, value time.Time) Label {
	return Label{name: name, time: value, fieldType: TimeType}
}

func DurationLabel(name string, value time.Duration) Label {
	return Label{name: name, duration: value, fieldType: DurationType}
}

func (l Label) Name() string {
	return l.name
}

func (l Label) Value() any {
	switch l.fieldType {
	case StringType:
		return l.string
	case IntType:
		return l.int
	case BoolType:
		return l.bool
	case FloatType:
		return l.float
	case TimeType:
		return l.time
	case DurationType:
		return l.duration
	default:
		return l.unknown
	}
}

// Expand can be used with the go-kit sugared logger
// e.g. logger.Infow("my message", labels.Expand(l1, l2, l3)...)
func Expand(labels ...Label) []any {
	result := make([]any, 0, len(labels)*2)
	for _, label := range labels {
		result = append(result, label.name, label.Value())
	}
	return result
}

// Zap can be used with the go-kit non-sugared logger
// e.g. logger.Infon("my message", l1.Zap(), l2.Zap(), l3.Zap())
func (l Label) Zap() zap.Field {
	switch l.fieldType {
	case StringType:
		return zap.String(l.name, l.string)
	case IntType:
		return zap.Int64(l.name, l.int)
	case BoolType:
		return zap.Bool(l.name, l.bool)
	case FloatType:
		return zap.Float64(l.name, l.float)
	case TimeType:
		return zap.Time(l.name, l.time)
	case DurationType:
		return zap.Duration(l.name, l.duration)
	default:
		return zap.Any(l.name, l.unknown)
	}
}
