// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type UpdatePullFrequencyResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *UpdatePullFrequencyResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
