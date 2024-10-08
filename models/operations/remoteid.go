// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type RemoteIDRequest struct {
	RemoteID string `queryParam:"style=form,explode=true,name=remoteId"`
}

func (o *RemoteIDRequest) GetRemoteID() string {
	if o == nil {
		return ""
	}
	return o.RemoteID
}

type RemoteIDResponse struct {
	HTTPMeta           components.HTTPMetadata `json:"-"`
	LinkedUserResponse *components.LinkedUserResponse
}

func (o *RemoteIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RemoteIDResponse) GetLinkedUserResponse() *components.LinkedUserResponse {
	if o == nil {
		return nil
	}
	return o.LinkedUserResponse
}
