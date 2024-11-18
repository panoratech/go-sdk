// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type ListAccountingCashflowStatementRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (o *ListAccountingCashflowStatementRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListAccountingCashflowStatementRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListAccountingCashflowStatementRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListAccountingCashflowStatementRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListAccountingCashflowStatementResponseBody struct {
	PrevCursor *string                                               `json:"prev_cursor"`
	NextCursor *string                                               `json:"next_cursor"`
	Data       []components.UnifiedAccountingCashflowstatementOutput `json:"data"`
}

func (o *ListAccountingCashflowStatementResponseBody) GetPrevCursor() *string {
	if o == nil {
		return nil
	}
	return o.PrevCursor
}

func (o *ListAccountingCashflowStatementResponseBody) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *ListAccountingCashflowStatementResponseBody) GetData() []components.UnifiedAccountingCashflowstatementOutput {
	if o == nil {
		return []components.UnifiedAccountingCashflowstatementOutput{}
	}
	return o.Data
}

type ListAccountingCashflowStatementResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListAccountingCashflowStatementResponseBody

	Next func() (*ListAccountingCashflowStatementResponse, error)
}

func (o *ListAccountingCashflowStatementResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListAccountingCashflowStatementResponse) GetObject() *ListAccountingCashflowStatementResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
