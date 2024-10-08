// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type RetrieveAtsAttachmentRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// id of the attachment you want to retrieve.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Set to true to include data from the original Ats software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
}

func (o *RetrieveAtsAttachmentRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *RetrieveAtsAttachmentRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RetrieveAtsAttachmentRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

type RetrieveAtsAttachmentResponse struct {
	HTTPMeta                   components.HTTPMetadata `json:"-"`
	UnifiedAtsAttachmentOutput *components.UnifiedAtsAttachmentOutput
}

func (o *RetrieveAtsAttachmentResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RetrieveAtsAttachmentResponse) GetUnifiedAtsAttachmentOutput() *components.UnifiedAtsAttachmentOutput {
	if o == nil {
		return nil
	}
	return o.UnifiedAtsAttachmentOutput
}
