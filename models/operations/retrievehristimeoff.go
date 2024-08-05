// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type RetrieveHrisTimeoffRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// id of the time off you want to retrieve.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Set to true to include data from the original Hris software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
}

func (o *RetrieveHrisTimeoffRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *RetrieveHrisTimeoffRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RetrieveHrisTimeoffRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

type RetrieveHrisTimeoffResponse struct {
	HTTPMeta                 components.HTTPMetadata `json:"-"`
	UnifiedHrisTimeoffOutput *components.UnifiedHrisTimeoffOutput
}

func (o *RetrieveHrisTimeoffResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RetrieveHrisTimeoffResponse) GetUnifiedHrisTimeoffOutput() *components.UnifiedHrisTimeoffOutput {
	if o == nil {
		return nil
	}
	return o.UnifiedHrisTimeoffOutput
}
