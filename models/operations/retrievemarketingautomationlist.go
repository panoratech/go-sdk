// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type RetrieveMarketingautomationListRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// id of the list you want to retrieve.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Set to true to include data from the original Marketingautomation software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
}

func (o *RetrieveMarketingautomationListRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *RetrieveMarketingautomationListRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RetrieveMarketingautomationListRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

type RetrieveMarketingautomationListResponse struct {
	HTTPMeta                             components.HTTPMetadata `json:"-"`
	UnifiedMarketingautomationListOutput *components.UnifiedMarketingautomationListOutput
}

func (o *RetrieveMarketingautomationListResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RetrieveMarketingautomationListResponse) GetUnifiedMarketingautomationListOutput() *components.UnifiedMarketingautomationListOutput {
	if o == nil {
		return nil
	}
	return o.UnifiedMarketingautomationListOutput
}
