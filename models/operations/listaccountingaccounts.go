// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type ListAccountingAccountsRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (o *ListAccountingAccountsRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListAccountingAccountsRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListAccountingAccountsRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListAccountingAccountsRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListAccountingAccountsResponseBody struct {
	PrevCursor *string                                     `json:"prev_cursor"`
	NextCursor *string                                     `json:"next_cursor"`
	Data       []components.UnifiedAccountingAccountOutput `json:"data"`
}

func (o *ListAccountingAccountsResponseBody) GetPrevCursor() *string {
	if o == nil {
		return nil
	}
	return o.PrevCursor
}

func (o *ListAccountingAccountsResponseBody) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *ListAccountingAccountsResponseBody) GetData() []components.UnifiedAccountingAccountOutput {
	if o == nil {
		return []components.UnifiedAccountingAccountOutput{}
	}
	return o.Data
}

type ListAccountingAccountsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListAccountingAccountsResponseBody

	Next func() (*ListAccountingAccountsResponse, error)
}

func (o *ListAccountingAccountsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListAccountingAccountsResponse) GetObject() *ListAccountingAccountsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
