// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type RetrieveMarketingautomationTemplateRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// id of the template you want to retrieve.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Set to true to include data from the original Marketingautomation software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
}

func (o *RetrieveMarketingautomationTemplateRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *RetrieveMarketingautomationTemplateRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RetrieveMarketingautomationTemplateRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

type RetrieveMarketingautomationTemplateResponse struct {
	HTTPMeta                                 components.HTTPMetadata `json:"-"`
	UnifiedMarketingautomationTemplateOutput *components.UnifiedMarketingautomationTemplateOutput
}

func (o *RetrieveMarketingautomationTemplateResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RetrieveMarketingautomationTemplateResponse) GetUnifiedMarketingautomationTemplateOutput() *components.UnifiedMarketingautomationTemplateOutput {
	if o == nil {
		return nil
	}
	return o.UnifiedMarketingautomationTemplateOutput
}