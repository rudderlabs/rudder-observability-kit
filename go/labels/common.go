// Package labels
// It contains labels for the common domain
// GENERATED CODE - DO NOT EDIT
package labels

import (
	log "github.com/rudderlabs/rudder-go-kit/logger"
)

var (
	DestinationId   = func(v string) log.Field { return log.NewStringField("destinationId", v) }
	DestinationType = func(v string) log.Field { return log.NewStringField("destinationType", v) }
	SourceId        = func(v string) log.Field { return log.NewStringField("sourceId", v) }
	SourceType      = func(v string) log.Field { return log.NewStringField("sourceType", v) }
	WorkspaceId     = func(v string) log.Field { return log.NewStringField("workspaceId", v) }
	Namespace       = func(v string) log.Field { return log.NewStringField("namespace", v) }
	Error           = func(v string) log.Field { return log.NewStringField("error", v) }
)
