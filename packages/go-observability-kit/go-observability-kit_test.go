package go_observability_kit

import (
	"testing"
)

func TestGoObservabilityKit(t *testing.T) {
	result := GoObservabilityKit("works")
	if result != "GoObservabilityKit works" {
		t.Error("Expected GoObservabilityKit to append 'works'")
	}
}
