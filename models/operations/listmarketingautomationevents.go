// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type ListMarketingAutomationEventsRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (o *ListMarketingAutomationEventsRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListMarketingAutomationEventsRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListMarketingAutomationEventsRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListMarketingAutomationEventsRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListMarketingAutomationEventsResponseBody struct {
	PrevCursor *string                                            `json:"prev_cursor"`
	NextCursor *string                                            `json:"next_cursor"`
	Data       []components.UnifiedMarketingautomationEventOutput `json:"data"`
}

func (o *ListMarketingAutomationEventsResponseBody) GetPrevCursor() *string {
	if o == nil {
		return nil
	}
	return o.PrevCursor
}

func (o *ListMarketingAutomationEventsResponseBody) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *ListMarketingAutomationEventsResponseBody) GetData() []components.UnifiedMarketingautomationEventOutput {
	if o == nil {
		return []components.UnifiedMarketingautomationEventOutput{}
	}
	return o.Data
}

type ListMarketingAutomationEventsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListMarketingAutomationEventsResponseBody

	Next func() (*ListMarketingAutomationEventsResponse, error)
}

func (o *ListMarketingAutomationEventsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListMarketingAutomationEventsResponse) GetObject() *ListMarketingAutomationEventsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
