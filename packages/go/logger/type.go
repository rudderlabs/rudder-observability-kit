package logger

import (
	"github.com/rudderlabs/rudder-observability-kit/packages/go/labels"
)

type Logger interface {
	Debug(string, ...labels.Label)
	Info(string, ...labels.Label)
	Warn(string, ...labels.Label)
	Error(string, ...labels.Label)
	Fatal(string, ...labels.Label)
	GetLabels() []labels.Label
	CreateLogger(...labels.Label) Logger
}
