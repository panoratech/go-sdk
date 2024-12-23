// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type ListMarketingautomationTemplatesRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (o *ListMarketingautomationTemplatesRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListMarketingautomationTemplatesRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListMarketingautomationTemplatesRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListMarketingautomationTemplatesRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListMarketingautomationTemplatesResponseBody struct {
	PrevCursor *string                                               `json:"prev_cursor"`
	NextCursor *string                                               `json:"next_cursor"`
	Data       []components.UnifiedMarketingautomationTemplateOutput `json:"data"`
}

func (o *ListMarketingautomationTemplatesResponseBody) GetPrevCursor() *string {
	if o == nil {
		return nil
	}
	return o.PrevCursor
}

func (o *ListMarketingautomationTemplatesResponseBody) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *ListMarketingautomationTemplatesResponseBody) GetData() []components.UnifiedMarketingautomationTemplateOutput {
	if o == nil {
		return []components.UnifiedMarketingautomationTemplateOutput{}
	}
	return o.Data
}

type ListMarketingautomationTemplatesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListMarketingautomationTemplatesResponseBody

	Next func() (*ListMarketingautomationTemplatesResponse, error)
}

func (o *ListMarketingautomationTemplatesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListMarketingautomationTemplatesResponse) GetObject() *ListMarketingautomationTemplatesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
