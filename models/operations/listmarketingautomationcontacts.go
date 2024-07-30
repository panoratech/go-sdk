// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"github.com/panoratech/go-sdk/models/components"
)

type ListMarketingAutomationContactsRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `default:"50" queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (l ListMarketingAutomationContactsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListMarketingAutomationContactsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListMarketingAutomationContactsRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListMarketingAutomationContactsRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListMarketingAutomationContactsRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListMarketingAutomationContactsRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListMarketingAutomationContactsResponseBody struct {
	PrevCursor string                                               `json:"prev_cursor"`
	NextCursor string                                               `json:"next_cursor"`
	Data       []components.UnifiedMarketingautomationContactOutput `json:"data"`
}

func (o *ListMarketingAutomationContactsResponseBody) GetPrevCursor() string {
	if o == nil {
		return ""
	}
	return o.PrevCursor
}

func (o *ListMarketingAutomationContactsResponseBody) GetNextCursor() string {
	if o == nil {
		return ""
	}
	return o.NextCursor
}

func (o *ListMarketingAutomationContactsResponseBody) GetData() []components.UnifiedMarketingautomationContactOutput {
	if o == nil {
		return []components.UnifiedMarketingautomationContactOutput{}
	}
	return o.Data
}

type ListMarketingAutomationContactsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListMarketingAutomationContactsResponseBody
}

func (o *ListMarketingAutomationContactsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListMarketingAutomationContactsResponse) GetObject() *ListMarketingAutomationContactsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
