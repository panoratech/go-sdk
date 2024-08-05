// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type ListAccountingTransactionRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (o *ListAccountingTransactionRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListAccountingTransactionRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListAccountingTransactionRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListAccountingTransactionRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListAccountingTransactionResponseBody struct {
	PrevCursor *string                                         `json:"prev_cursor"`
	NextCursor *string                                         `json:"next_cursor"`
	Data       []components.UnifiedAccountingTransactionOutput `json:"data"`
}

func (o *ListAccountingTransactionResponseBody) GetPrevCursor() *string {
	if o == nil {
		return nil
	}
	return o.PrevCursor
}

func (o *ListAccountingTransactionResponseBody) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *ListAccountingTransactionResponseBody) GetData() []components.UnifiedAccountingTransactionOutput {
	if o == nil {
		return []components.UnifiedAccountingTransactionOutput{}
	}
	return o.Data
}

type ListAccountingTransactionResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListAccountingTransactionResponseBody
}

func (o *ListAccountingTransactionResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListAccountingTransactionResponse) GetObject() *ListAccountingTransactionResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
