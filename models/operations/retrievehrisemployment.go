// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type RetrieveHrisEmploymentRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// id of the employment you want to retrieve.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Set to true to include data from the original Hris software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
}

func (o *RetrieveHrisEmploymentRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *RetrieveHrisEmploymentRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RetrieveHrisEmploymentRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

type RetrieveHrisEmploymentResponse struct {
	HTTPMeta                    components.HTTPMetadata `json:"-"`
	UnifiedHrisEmploymentOutput *components.UnifiedHrisEmploymentOutput
}

func (o *RetrieveHrisEmploymentResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RetrieveHrisEmploymentResponse) GetUnifiedHrisEmploymentOutput() *components.UnifiedHrisEmploymentOutput {
	if o == nil {
		return nil
	}
	return o.UnifiedHrisEmploymentOutput
}
