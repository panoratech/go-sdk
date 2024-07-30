// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type RetrieveCrmCompanyRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// id of the company you want to retrieve.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Set to true to include data from the original Crm software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
}

func (o *RetrieveCrmCompanyRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *RetrieveCrmCompanyRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RetrieveCrmCompanyRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

type RetrieveCrmCompanyResponse struct {
	HTTPMeta                components.HTTPMetadata `json:"-"`
	UnifiedCrmCompanyOutput *components.UnifiedCrmCompanyOutput
}

func (o *RetrieveCrmCompanyResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RetrieveCrmCompanyResponse) GetUnifiedCrmCompanyOutput() *components.UnifiedCrmCompanyOutput {
	if o == nil {
		return nil
	}
	return o.UnifiedCrmCompanyOutput
}
