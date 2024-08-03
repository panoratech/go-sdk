// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"github.com/panoratech/go-sdk/models/components"
)

type ListTicketingCollectionsRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `default:"50" queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (l ListTicketingCollectionsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListTicketingCollectionsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListTicketingCollectionsRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListTicketingCollectionsRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListTicketingCollectionsRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListTicketingCollectionsRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListTicketingCollectionsResponseBody struct {
	PrevCursor *string                                       `json:"prev_cursor"`
	NextCursor *string                                       `json:"next_cursor"`
	Data       []components.UnifiedTicketingCollectionOutput `json:"data"`
}

func (o *ListTicketingCollectionsResponseBody) GetPrevCursor() *string {
	if o == nil {
		return nil
	}
	return o.PrevCursor
}

func (o *ListTicketingCollectionsResponseBody) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *ListTicketingCollectionsResponseBody) GetData() []components.UnifiedTicketingCollectionOutput {
	if o == nil {
		return []components.UnifiedTicketingCollectionOutput{}
	}
	return o.Data
}

type ListTicketingCollectionsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListTicketingCollectionsResponseBody
}

func (o *ListTicketingCollectionsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListTicketingCollectionsResponse) GetObject() *ListTicketingCollectionsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
