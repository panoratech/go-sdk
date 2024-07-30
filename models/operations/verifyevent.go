// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type VerifyEventResponse struct {
	HTTPMeta     components.HTTPMetadata `json:"-"`
	EventPayload *components.EventPayload
}

func (o *VerifyEventResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *VerifyEventResponse) GetEventPayload() *components.EventPayload {
	if o == nil {
		return nil
	}
	return o.EventPayload
}
