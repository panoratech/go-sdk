// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type CreateTicketingCommentRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original Ticketing software.
	RemoteData                   *bool                                   `queryParam:"style=form,explode=true,name=remote_data"`
	UnifiedTicketingCommentInput components.UnifiedTicketingCommentInput `request:"mediaType=application/json"`
}

func (o *CreateTicketingCommentRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *CreateTicketingCommentRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *CreateTicketingCommentRequest) GetUnifiedTicketingCommentInput() components.UnifiedTicketingCommentInput {
	if o == nil {
		return components.UnifiedTicketingCommentInput{}
	}
	return o.UnifiedTicketingCommentInput
}

type CreateTicketingCommentResponse struct {
	HTTPMeta                      components.HTTPMetadata `json:"-"`
	UnifiedTicketingCommentOutput *components.UnifiedTicketingCommentOutput
}

func (o *CreateTicketingCommentResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateTicketingCommentResponse) GetUnifiedTicketingCommentOutput() *components.UnifiedTicketingCommentOutput {
	if o == nil {
		return nil
	}
	return o.UnifiedTicketingCommentOutput
}
