// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type RetrieveTicketingTagRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// id of the tag you want to retrieve.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Set to true to include data from the original Ticketing software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
}

func (o *RetrieveTicketingTagRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *RetrieveTicketingTagRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RetrieveTicketingTagRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

type RetrieveTicketingTagResponse struct {
	HTTPMeta                  components.HTTPMetadata `json:"-"`
	UnifiedTicketingTagOutput *components.UnifiedTicketingTagOutput
}

func (o *RetrieveTicketingTagResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RetrieveTicketingTagResponse) GetUnifiedTicketingTagOutput() *components.UnifiedTicketingTagOutput {
	if o == nil {
		return nil
	}
	return o.UnifiedTicketingTagOutput
}