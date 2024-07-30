// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type RetrieveMarketingautomationEventRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// id of the event you want to retrieve.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Set to true to include data from the original Marketingautomation software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
}

func (o *RetrieveMarketingautomationEventRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *RetrieveMarketingautomationEventRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RetrieveMarketingautomationEventRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

type RetrieveMarketingautomationEventResponse struct {
	HTTPMeta                              components.HTTPMetadata `json:"-"`
	UnifiedMarketingautomationEventOutput *components.UnifiedMarketingautomationEventOutput
}

func (o *RetrieveMarketingautomationEventResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RetrieveMarketingautomationEventResponse) GetUnifiedMarketingautomationEventOutput() *components.UnifiedMarketingautomationEventOutput {
	if o == nil {
		return nil
	}
	return o.UnifiedMarketingautomationEventOutput
}