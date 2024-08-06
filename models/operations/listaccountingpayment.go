// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type ListAccountingPaymentRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (o *ListAccountingPaymentRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListAccountingPaymentRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListAccountingPaymentRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListAccountingPaymentRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListAccountingPaymentResponseBody struct {
	PrevCursor *string                                     `json:"prev_cursor"`
	NextCursor *string                                     `json:"next_cursor"`
	Data       []components.UnifiedAccountingPaymentOutput `json:"data"`
}

func (o *ListAccountingPaymentResponseBody) GetPrevCursor() *string {
	if o == nil {
		return nil
	}
	return o.PrevCursor
}

func (o *ListAccountingPaymentResponseBody) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *ListAccountingPaymentResponseBody) GetData() []components.UnifiedAccountingPaymentOutput {
	if o == nil {
		return []components.UnifiedAccountingPaymentOutput{}
	}
	return o.Data
}

type ListAccountingPaymentResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListAccountingPaymentResponseBody

	Next func() (*ListAccountingPaymentResponse, error)
}

func (o *ListAccountingPaymentResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListAccountingPaymentResponse) GetObject() *ListAccountingPaymentResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
