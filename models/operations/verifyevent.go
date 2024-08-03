// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type VerifyEventResponseBody struct {
	// Dynamic event payload
	Data map[string]any `json:"data,omitempty"`
}

func (o *VerifyEventResponseBody) GetData() map[string]any {
	if o == nil {
		return nil
	}
	return o.Data
}

type VerifyEventResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *VerifyEventResponseBody
}

func (o *VerifyEventResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *VerifyEventResponse) GetObject() *VerifyEventResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
