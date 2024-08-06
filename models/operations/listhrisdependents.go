// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type ListHrisDependentsRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (o *ListHrisDependentsRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListHrisDependentsRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListHrisDependentsRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListHrisDependentsRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListHrisDependentsResponseBody struct {
	PrevCursor *string                                 `json:"prev_cursor"`
	NextCursor *string                                 `json:"next_cursor"`
	Data       []components.UnifiedHrisDependentOutput `json:"data"`
}

func (o *ListHrisDependentsResponseBody) GetPrevCursor() *string {
	if o == nil {
		return nil
	}
	return o.PrevCursor
}

func (o *ListHrisDependentsResponseBody) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *ListHrisDependentsResponseBody) GetData() []components.UnifiedHrisDependentOutput {
	if o == nil {
		return []components.UnifiedHrisDependentOutput{}
	}
	return o.Data
}

type ListHrisDependentsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListHrisDependentsResponseBody

	Next func() (*ListHrisDependentsResponse, error)
}

func (o *ListHrisDependentsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListHrisDependentsResponse) GetObject() *ListHrisDependentsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
