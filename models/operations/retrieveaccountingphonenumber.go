// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type RetrieveAccountingPhonenumberRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// id of the phonenumber you want to retrieve.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Set to true to include data from the original Accounting software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
}

func (o *RetrieveAccountingPhonenumberRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *RetrieveAccountingPhonenumberRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RetrieveAccountingPhonenumberRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

type RetrieveAccountingPhonenumberResponse struct {
	HTTPMeta                           components.HTTPMetadata `json:"-"`
	UnifiedAccountingPhonenumberOutput *components.UnifiedAccountingPhonenumberOutput
}

func (o *RetrieveAccountingPhonenumberResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RetrieveAccountingPhonenumberResponse) GetUnifiedAccountingPhonenumberOutput() *components.UnifiedAccountingPhonenumberOutput {
	if o == nil {
		return nil
	}
	return o.UnifiedAccountingPhonenumberOutput
}
