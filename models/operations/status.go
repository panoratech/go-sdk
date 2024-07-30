// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type StatusRequest struct {
	Vertical string `pathParam:"style=simple,explode=false,name=vertical"`
}

func (o *StatusRequest) GetVertical() string {
	if o == nil {
		return ""
	}
	return o.Vertical
}

type StatusResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *StatusResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}