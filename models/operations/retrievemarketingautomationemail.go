// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type RetrieveMarketingautomationEmailRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// id of the email you want to retrieve.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Set to true to include data from the original Marketingautomation software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
}

func (o *RetrieveMarketingautomationEmailRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *RetrieveMarketingautomationEmailRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RetrieveMarketingautomationEmailRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

type RetrieveMarketingautomationEmailResponse struct {
	HTTPMeta                              components.HTTPMetadata `json:"-"`
	UnifiedMarketingautomationEmailOutput *components.UnifiedMarketingautomationEmailOutput
}

func (o *RetrieveMarketingautomationEmailResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RetrieveMarketingautomationEmailResponse) GetUnifiedMarketingautomationEmailOutput() *components.UnifiedMarketingautomationEmailOutput {
	if o == nil {
		return nil
	}
	return o.UnifiedMarketingautomationEmailOutput
}
