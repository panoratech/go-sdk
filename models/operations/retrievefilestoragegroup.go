// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type RetrieveFilestorageGroupRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// id of the permission you want to retrieve.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Set to true to include data from the original File Storage software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
}

func (o *RetrieveFilestorageGroupRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *RetrieveFilestorageGroupRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RetrieveFilestorageGroupRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

type RetrieveFilestorageGroupResponse struct {
	HTTPMeta                      components.HTTPMetadata `json:"-"`
	UnifiedFilestorageGroupOutput *components.UnifiedFilestorageGroupOutput
}

func (o *RetrieveFilestorageGroupResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RetrieveFilestorageGroupResponse) GetUnifiedFilestorageGroupOutput() *components.UnifiedFilestorageGroupOutput {
	if o == nil {
		return nil
	}
	return o.UnifiedFilestorageGroupOutput
}
