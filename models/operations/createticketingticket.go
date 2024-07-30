// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type CreateTicketingTicketRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original Ticketing software.
	RemoteData                  *bool                                  `queryParam:"style=form,explode=true,name=remote_data"`
	UnifiedTicketingTicketInput components.UnifiedTicketingTicketInput `request:"mediaType=application/json"`
}

func (o *CreateTicketingTicketRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *CreateTicketingTicketRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *CreateTicketingTicketRequest) GetUnifiedTicketingTicketInput() components.UnifiedTicketingTicketInput {
	if o == nil {
		return components.UnifiedTicketingTicketInput{}
	}
	return o.UnifiedTicketingTicketInput
}

type CreateTicketingTicketResponse struct {
	HTTPMeta                     components.HTTPMetadata `json:"-"`
	UnifiedTicketingTicketOutput *components.UnifiedTicketingTicketOutput
}

func (o *CreateTicketingTicketResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateTicketingTicketResponse) GetUnifiedTicketingTicketOutput() *components.UnifiedTicketingTicketOutput {
	if o == nil {
		return nil
	}
	return o.UnifiedTicketingTicketOutput
}