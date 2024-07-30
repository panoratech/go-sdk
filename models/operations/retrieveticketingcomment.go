// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type RetrieveTicketingCommentRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// id of the `comment` you want to retrive.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Set to true to include data from the original Ticketing software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
}

func (o *RetrieveTicketingCommentRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *RetrieveTicketingCommentRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RetrieveTicketingCommentRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

type RetrieveTicketingCommentResponseBody struct {
	PrevCursor string                                     `json:"prev_cursor"`
	NextCursor string                                     `json:"next_cursor"`
	Data       []components.UnifiedTicketingCommentOutput `json:"data"`
}

func (o *RetrieveTicketingCommentResponseBody) GetPrevCursor() string {
	if o == nil {
		return ""
	}
	return o.PrevCursor
}

func (o *RetrieveTicketingCommentResponseBody) GetNextCursor() string {
	if o == nil {
		return ""
	}
	return o.NextCursor
}

func (o *RetrieveTicketingCommentResponseBody) GetData() []components.UnifiedTicketingCommentOutput {
	if o == nil {
		return []components.UnifiedTicketingCommentOutput{}
	}
	return o.Data
}

type RetrieveTicketingCommentResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *RetrieveTicketingCommentResponseBody
}

func (o *RetrieveTicketingCommentResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RetrieveTicketingCommentResponse) GetObject() *RetrieveTicketingCommentResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
