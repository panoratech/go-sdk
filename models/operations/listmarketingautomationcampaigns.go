// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type ListMarketingautomationCampaignsRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (o *ListMarketingautomationCampaignsRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListMarketingautomationCampaignsRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListMarketingautomationCampaignsRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListMarketingautomationCampaignsRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListMarketingautomationCampaignsResponseBody struct {
	PrevCursor *string                                               `json:"prev_cursor"`
	NextCursor *string                                               `json:"next_cursor"`
	Data       []components.UnifiedMarketingautomationCampaignOutput `json:"data"`
}

func (o *ListMarketingautomationCampaignsResponseBody) GetPrevCursor() *string {
	if o == nil {
		return nil
	}
	return o.PrevCursor
}

func (o *ListMarketingautomationCampaignsResponseBody) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *ListMarketingautomationCampaignsResponseBody) GetData() []components.UnifiedMarketingautomationCampaignOutput {
	if o == nil {
		return []components.UnifiedMarketingautomationCampaignOutput{}
	}
	return o.Data
}

type ListMarketingautomationCampaignsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListMarketingautomationCampaignsResponseBody

	Next func() (*ListMarketingautomationCampaignsResponse, error)
}

func (o *ListMarketingautomationCampaignsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListMarketingautomationCampaignsResponse) GetObject() *ListMarketingautomationCampaignsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
