// Package labels
// It contains labels for the common domain
// GENERATED CODE - DO NOT EDIT
package labels

import "time"

var (
	DestinationId   LabelFunc = func(v any) Label { return StringLabel("destinationId", v.(string)) }
	DestinationType LabelFunc = func(v any) Label { return StringLabel("destinationType", v.(string)) }
	SourceId        LabelFunc = func(v any) Label { return StringLabel("sourceId", v.(string)) }
	SourceType      LabelFunc = func(v any) Label { return StringLabel("sourceType", v.(string)) }
	WorkspaceId     LabelFunc = func(v any) Label { return StringLabel("workspaceId", v.(string)) }
	Namespace       LabelFunc = func(v any) Label { return StringLabel("namespace", v.(string)) }
	Error           LabelFunc = func(v any) Label { return StringLabel("error", v.(string)) }
)
