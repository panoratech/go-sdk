// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type CreateTicketingAttachmentRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original Ticketing software.
	RemoteData                      *bool                                      `queryParam:"style=form,explode=true,name=remote_data"`
	UnifiedTicketingAttachmentInput components.UnifiedTicketingAttachmentInput `request:"mediaType=application/json"`
}

func (o *CreateTicketingAttachmentRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *CreateTicketingAttachmentRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *CreateTicketingAttachmentRequest) GetUnifiedTicketingAttachmentInput() components.UnifiedTicketingAttachmentInput {
	if o == nil {
		return components.UnifiedTicketingAttachmentInput{}
	}
	return o.UnifiedTicketingAttachmentInput
}

type CreateTicketingAttachmentResponse struct {
	HTTPMeta                         components.HTTPMetadata `json:"-"`
	UnifiedTicketingAttachmentOutput *components.UnifiedTicketingAttachmentOutput
}

func (o *CreateTicketingAttachmentResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateTicketingAttachmentResponse) GetUnifiedTicketingAttachmentOutput() *components.UnifiedTicketingAttachmentOutput {
	if o == nil {
		return nil
	}
	return o.UnifiedTicketingAttachmentOutput
}