// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type RetrieveCollectionRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// id of the collection you want to retrieve.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Set to true to include data from the original Ticketing software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
}

func (o *RetrieveCollectionRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *RetrieveCollectionRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RetrieveCollectionRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

type RetrieveCollectionResponse struct {
	HTTPMeta                         components.HTTPMetadata `json:"-"`
	UnifiedTicketingCollectionOutput *components.UnifiedTicketingCollectionOutput
}

func (o *RetrieveCollectionResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RetrieveCollectionResponse) GetUnifiedTicketingCollectionOutput() *components.UnifiedTicketingCollectionOutput {
	if o == nil {
		return nil
	}
	return o.UnifiedTicketingCollectionOutput
}
