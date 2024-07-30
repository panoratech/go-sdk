// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type RetrieveCrmContactRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// id of the `contact` you want to retrive.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Set to true to include data from the original CRM software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
}

func (o *RetrieveCrmContactRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *RetrieveCrmContactRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RetrieveCrmContactRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

type RetrieveCrmContactResponse struct {
	HTTPMeta                components.HTTPMetadata `json:"-"`
	UnifiedCrmContactOutput *components.UnifiedCrmContactOutput
}

func (o *RetrieveCrmContactResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RetrieveCrmContactResponse) GetUnifiedCrmContactOutput() *components.UnifiedCrmContactOutput {
	if o == nil {
		return nil
	}
	return o.UnifiedCrmContactOutput
}
