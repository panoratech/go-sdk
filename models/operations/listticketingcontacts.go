// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"github.com/panoratech/go-sdk/models/components"
)

type ListTicketingContactsRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `default:"50" queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (l ListTicketingContactsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListTicketingContactsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListTicketingContactsRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListTicketingContactsRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListTicketingContactsRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListTicketingContactsRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListTicketingContactsResponseBody struct {
	PrevCursor *string                                    `json:"prev_cursor"`
	NextCursor *string                                    `json:"next_cursor"`
	Data       []components.UnifiedTicketingContactOutput `json:"data"`
}

func (o *ListTicketingContactsResponseBody) GetPrevCursor() *string {
	if o == nil {
		return nil
	}
	return o.PrevCursor
}

func (o *ListTicketingContactsResponseBody) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *ListTicketingContactsResponseBody) GetData() []components.UnifiedTicketingContactOutput {
	if o == nil {
		return []components.UnifiedTicketingContactOutput{}
	}
	return o.Data
}

type ListTicketingContactsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListTicketingContactsResponseBody
}

func (o *ListTicketingContactsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListTicketingContactsResponse) GetObject() *ListTicketingContactsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
