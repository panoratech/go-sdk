// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type ListAtsRejectReasonsRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (o *ListAtsRejectReasonsRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListAtsRejectReasonsRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListAtsRejectReasonsRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListAtsRejectReasonsRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListAtsRejectReasonsResponseBody struct {
	PrevCursor *string                                   `json:"prev_cursor"`
	NextCursor *string                                   `json:"next_cursor"`
	Data       []components.UnifiedAtsRejectreasonOutput `json:"data"`
}

func (o *ListAtsRejectReasonsResponseBody) GetPrevCursor() *string {
	if o == nil {
		return nil
	}
	return o.PrevCursor
}

func (o *ListAtsRejectReasonsResponseBody) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *ListAtsRejectReasonsResponseBody) GetData() []components.UnifiedAtsRejectreasonOutput {
	if o == nil {
		return []components.UnifiedAtsRejectreasonOutput{}
	}
	return o.Data
}

type ListAtsRejectReasonsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListAtsRejectReasonsResponseBody
}

func (o *ListAtsRejectReasonsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListAtsRejectReasonsResponse) GetObject() *ListAtsRejectReasonsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
