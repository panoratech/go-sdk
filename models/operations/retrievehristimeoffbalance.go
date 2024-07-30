// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type RetrieveHrisTimeoffbalanceRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// id of the timeoffbalance you want to retrieve.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Set to true to include data from the original Hris software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
}

func (o *RetrieveHrisTimeoffbalanceRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *RetrieveHrisTimeoffbalanceRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RetrieveHrisTimeoffbalanceRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

type RetrieveHrisTimeoffbalanceResponse struct {
	HTTPMeta                        components.HTTPMetadata `json:"-"`
	UnifiedHrisTimeoffbalanceOutput *components.UnifiedHrisTimeoffbalanceOutput
}

func (o *RetrieveHrisTimeoffbalanceResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RetrieveHrisTimeoffbalanceResponse) GetUnifiedHrisTimeoffbalanceOutput() *components.UnifiedHrisTimeoffbalanceOutput {
	if o == nil {
		return nil
	}
	return o.UnifiedHrisTimeoffbalanceOutput
}
