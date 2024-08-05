// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type CreateMarketingautomationCampaignRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original Marketingautomation software.
	RemoteData                              *bool                                              `queryParam:"style=form,explode=true,name=remote_data"`
	UnifiedMarketingautomationCampaignInput components.UnifiedMarketingautomationCampaignInput `request:"mediaType=application/json"`
}

func (o *CreateMarketingautomationCampaignRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *CreateMarketingautomationCampaignRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *CreateMarketingautomationCampaignRequest) GetUnifiedMarketingautomationCampaignInput() components.UnifiedMarketingautomationCampaignInput {
	if o == nil {
		return components.UnifiedMarketingautomationCampaignInput{}
	}
	return o.UnifiedMarketingautomationCampaignInput
}

type CreateMarketingautomationCampaignResponse struct {
	HTTPMeta                                 components.HTTPMetadata `json:"-"`
	UnifiedMarketingautomationCampaignOutput *components.UnifiedMarketingautomationCampaignOutput
}

func (o *CreateMarketingautomationCampaignResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateMarketingautomationCampaignResponse) GetUnifiedMarketingautomationCampaignOutput() *components.UnifiedMarketingautomationCampaignOutput {
	if o == nil {
		return nil
	}
	return o.UnifiedMarketingautomationCampaignOutput
}
