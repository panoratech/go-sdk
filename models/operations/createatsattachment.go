// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type CreateAtsAttachmentRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original Ats software.
	RemoteData                *bool                                `queryParam:"style=form,explode=true,name=remote_data"`
	UnifiedAtsAttachmentInput components.UnifiedAtsAttachmentInput `request:"mediaType=application/json"`
}

func (o *CreateAtsAttachmentRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *CreateAtsAttachmentRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *CreateAtsAttachmentRequest) GetUnifiedAtsAttachmentInput() components.UnifiedAtsAttachmentInput {
	if o == nil {
		return components.UnifiedAtsAttachmentInput{}
	}
	return o.UnifiedAtsAttachmentInput
}

type CreateAtsAttachmentResponse struct {
	HTTPMeta                   components.HTTPMetadata `json:"-"`
	UnifiedAtsAttachmentOutput *components.UnifiedAtsAttachmentOutput
}

func (o *CreateAtsAttachmentResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateAtsAttachmentResponse) GetUnifiedAtsAttachmentOutput() *components.UnifiedAtsAttachmentOutput {
	if o == nil {
		return nil
	}
	return o.UnifiedAtsAttachmentOutput
}
