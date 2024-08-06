// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type ListAccountingItemRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (o *ListAccountingItemRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListAccountingItemRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListAccountingItemRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListAccountingItemRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListAccountingItemResponseBody struct {
	PrevCursor *string                                  `json:"prev_cursor"`
	NextCursor *string                                  `json:"next_cursor"`
	Data       []components.UnifiedAccountingItemOutput `json:"data"`
}

func (o *ListAccountingItemResponseBody) GetPrevCursor() *string {
	if o == nil {
		return nil
	}
	return o.PrevCursor
}

func (o *ListAccountingItemResponseBody) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *ListAccountingItemResponseBody) GetData() []components.UnifiedAccountingItemOutput {
	if o == nil {
		return []components.UnifiedAccountingItemOutput{}
	}
	return o.Data
}

type ListAccountingItemResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListAccountingItemResponseBody

	Next func() (*ListAccountingItemResponse, error)
}

func (o *ListAccountingItemResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListAccountingItemResponse) GetObject() *ListAccountingItemResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
