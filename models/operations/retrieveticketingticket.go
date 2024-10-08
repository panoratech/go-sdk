// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type RetrieveTicketingTicketRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// id of the `ticket` you want to retrive.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Set to true to include data from the original Ticketing software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
}

func (o *RetrieveTicketingTicketRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *RetrieveTicketingTicketRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RetrieveTicketingTicketRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

type RetrieveTicketingTicketResponse struct {
	HTTPMeta                     components.HTTPMetadata `json:"-"`
	UnifiedTicketingTicketOutput *components.UnifiedTicketingTicketOutput
}

func (o *RetrieveTicketingTicketResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RetrieveTicketingTicketResponse) GetUnifiedTicketingTicketOutput() *components.UnifiedTicketingTicketOutput {
	if o == nil {
		return nil
	}
	return o.UnifiedTicketingTicketOutput
}
