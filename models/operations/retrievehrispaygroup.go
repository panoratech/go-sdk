// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type RetrieveHrisPaygroupRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// id of the paygroup you want to retrieve.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Set to true to include data from the original Hris software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
}

func (o *RetrieveHrisPaygroupRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *RetrieveHrisPaygroupRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RetrieveHrisPaygroupRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

type RetrieveHrisPaygroupResponse struct {
	HTTPMeta                  components.HTTPMetadata `json:"-"`
	UnifiedHrisPaygroupOutput *components.UnifiedHrisPaygroupOutput
}

func (o *RetrieveHrisPaygroupResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RetrieveHrisPaygroupResponse) GetUnifiedHrisPaygroupOutput() *components.UnifiedHrisPaygroupOutput {
	if o == nil {
		return nil
	}
	return o.UnifiedHrisPaygroupOutput
}
