// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type CreateMarketingAutomationContactRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original Marketingautomation software.
	RemoteData                             *bool                                             `queryParam:"style=form,explode=true,name=remote_data"`
	UnifiedMarketingautomationContactInput components.UnifiedMarketingautomationContactInput `request:"mediaType=application/json"`
}

func (o *CreateMarketingAutomationContactRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *CreateMarketingAutomationContactRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *CreateMarketingAutomationContactRequest) GetUnifiedMarketingautomationContactInput() components.UnifiedMarketingautomationContactInput {
	if o == nil {
		return components.UnifiedMarketingautomationContactInput{}
	}
	return o.UnifiedMarketingautomationContactInput
}

type CreateMarketingAutomationContactResponse struct {
	HTTPMeta                                components.HTTPMetadata `json:"-"`
	UnifiedMarketingautomationContactOutput *components.UnifiedMarketingautomationContactOutput
}

func (o *CreateMarketingAutomationContactResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateMarketingAutomationContactResponse) GetUnifiedMarketingautomationContactOutput() *components.UnifiedMarketingautomationContactOutput {
	if o == nil {
		return nil
	}
	return o.UnifiedMarketingautomationContactOutput
}
