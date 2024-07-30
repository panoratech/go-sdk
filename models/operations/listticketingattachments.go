// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"github.com/panoratech/go-sdk/models/components"
)

type ListTicketingAttachmentsRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `default:"50" queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (l ListTicketingAttachmentsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListTicketingAttachmentsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListTicketingAttachmentsRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListTicketingAttachmentsRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListTicketingAttachmentsRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListTicketingAttachmentsRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListTicketingAttachmentsResponseBody struct {
	PrevCursor string                                        `json:"prev_cursor"`
	NextCursor string                                        `json:"next_cursor"`
	Data       []components.UnifiedTicketingAttachmentOutput `json:"data"`
}

func (o *ListTicketingAttachmentsResponseBody) GetPrevCursor() string {
	if o == nil {
		return ""
	}
	return o.PrevCursor
}

func (o *ListTicketingAttachmentsResponseBody) GetNextCursor() string {
	if o == nil {
		return ""
	}
	return o.NextCursor
}

func (o *ListTicketingAttachmentsResponseBody) GetData() []components.UnifiedTicketingAttachmentOutput {
	if o == nil {
		return []components.UnifiedTicketingAttachmentOutput{}
	}
	return o.Data
}

type ListTicketingAttachmentsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListTicketingAttachmentsResponseBody
}

func (o *ListTicketingAttachmentsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListTicketingAttachmentsResponse) GetObject() *ListTicketingAttachmentsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
