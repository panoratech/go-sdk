// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type HelloResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns a greeting message
	Res *string
}

func (o *HelloResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *HelloResponse) GetRes() *string {
	if o == nil {
		return nil
	}
	return o.Res
}
