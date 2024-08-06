// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type ListAtsEeocsRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (o *ListAtsEeocsRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListAtsEeocsRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListAtsEeocsRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListAtsEeocsRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListAtsEeocsResponseBody struct {
	PrevCursor *string                            `json:"prev_cursor"`
	NextCursor *string                            `json:"next_cursor"`
	Data       []components.UnifiedAtsEeocsOutput `json:"data"`
}

func (o *ListAtsEeocsResponseBody) GetPrevCursor() *string {
	if o == nil {
		return nil
	}
	return o.PrevCursor
}

func (o *ListAtsEeocsResponseBody) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *ListAtsEeocsResponseBody) GetData() []components.UnifiedAtsEeocsOutput {
	if o == nil {
		return []components.UnifiedAtsEeocsOutput{}
	}
	return o.Data
}

type ListAtsEeocsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListAtsEeocsResponseBody

	Next func() (*ListAtsEeocsResponse, error)
}

func (o *ListAtsEeocsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListAtsEeocsResponse) GetObject() *ListAtsEeocsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
