// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type ListLinkedUsersResponse struct {
	HTTPMeta            components.HTTPMetadata `json:"-"`
	LinkedUserResponses []components.LinkedUserResponse
}

func (o *ListLinkedUsersResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListLinkedUsersResponse) GetLinkedUserResponses() []components.LinkedUserResponse {
	if o == nil {
		return nil
	}
	return o.LinkedUserResponses
}
