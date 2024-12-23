// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type ListFilestorageGroupRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (o *ListFilestorageGroupRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListFilestorageGroupRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListFilestorageGroupRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListFilestorageGroupRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListFilestorageGroupResponseBody struct {
	PrevCursor *string                                    `json:"prev_cursor"`
	NextCursor *string                                    `json:"next_cursor"`
	Data       []components.UnifiedFilestorageGroupOutput `json:"data"`
}

func (o *ListFilestorageGroupResponseBody) GetPrevCursor() *string {
	if o == nil {
		return nil
	}
	return o.PrevCursor
}

func (o *ListFilestorageGroupResponseBody) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *ListFilestorageGroupResponseBody) GetData() []components.UnifiedFilestorageGroupOutput {
	if o == nil {
		return []components.UnifiedFilestorageGroupOutput{}
	}
	return o.Data
}

type ListFilestorageGroupResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListFilestorageGroupResponseBody

	Next func() (*ListFilestorageGroupResponse, error)
}

func (o *ListFilestorageGroupResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListFilestorageGroupResponse) GetObject() *ListFilestorageGroupResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
