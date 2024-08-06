// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type ListAtsAttachmentRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (o *ListAtsAttachmentRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListAtsAttachmentRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListAtsAttachmentRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListAtsAttachmentRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListAtsAttachmentResponseBody struct {
	PrevCursor *string                                 `json:"prev_cursor"`
	NextCursor *string                                 `json:"next_cursor"`
	Data       []components.UnifiedAtsAttachmentOutput `json:"data"`
}

func (o *ListAtsAttachmentResponseBody) GetPrevCursor() *string {
	if o == nil {
		return nil
	}
	return o.PrevCursor
}

func (o *ListAtsAttachmentResponseBody) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *ListAtsAttachmentResponseBody) GetData() []components.UnifiedAtsAttachmentOutput {
	if o == nil {
		return []components.UnifiedAtsAttachmentOutput{}
	}
	return o.Data
}

type ListAtsAttachmentResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListAtsAttachmentResponseBody

	Next func() (*ListAtsAttachmentResponse, error)
}

func (o *ListAtsAttachmentResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListAtsAttachmentResponse) GetObject() *ListAtsAttachmentResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
