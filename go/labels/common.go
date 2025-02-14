// Package labels
// It contains labels for the common domain
// GENERATED CODE - DO NOT EDIT
package labels

import (
	log "github.com/rudderlabs/rudder-go-kit/logger"
)

var (
	DestinationID    = func(v string) log.Field { return log.NewStringField("destinationId", v) }
	DestinationType  = func(v string) log.Field { return log.NewStringField("destinationType", v) }
	SourceID         = func(v string) log.Field { return log.NewStringField("sourceId", v) }
	SourceType       = func(v string) log.Field { return log.NewStringField("sourceType", v) }
	WorkspaceID      = func(v string) log.Field { return log.NewStringField("workspaceId", v) }
	TransformationID = func(v string) log.Field { return log.NewStringField("transformationId", v) }
	Namespace        = func(v string) log.Field { return log.NewStringField("namespace", v) }
	Error            = func(v error) log.Field { return log.NewErrorField(v) }
)
