// Package labels
// It contains labels for the common domain
// GENERATED CODE - DO NOT EDIT
package labels

import (
	log "github.com/rudderlabs/rudder-go-kit/logger"
)

var (
	DestinationId   = func(v any) log.Field { return log.NewStringField("destinationId", v.(string)) }
	DestinationType = func(v any) log.Field { return log.NewStringField("destinationType", v.(string)) }
	SourceId        = func(v any) log.Field { return log.NewStringField("sourceId", v.(string)) }
	SourceType      = func(v any) log.Field { return log.NewStringField("sourceType", v.(string)) }
	WorkspaceId     = func(v any) log.Field { return log.NewStringField("workspaceId", v.(string)) }
	Namespace       = func(v any) log.Field { return log.NewStringField("namespace", v.(string)) }
	Error           = func(v any) log.Field { return log.NewStringField("error", v.(string)) }
)
