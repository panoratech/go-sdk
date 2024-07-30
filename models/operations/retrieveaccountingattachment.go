// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type RetrieveAccountingAttachmentRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// id of the attachment you want to retrieve.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Set to true to include data from the original Accounting software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
}

func (o *RetrieveAccountingAttachmentRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *RetrieveAccountingAttachmentRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RetrieveAccountingAttachmentRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

type RetrieveAccountingAttachmentResponse struct {
	HTTPMeta                          components.HTTPMetadata `json:"-"`
	UnifiedAccountingAttachmentOutput *components.UnifiedAccountingAttachmentOutput
}

func (o *RetrieveAccountingAttachmentResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RetrieveAccountingAttachmentResponse) GetUnifiedAccountingAttachmentOutput() *components.UnifiedAccountingAttachmentOutput {
	if o == nil {
		return nil
	}
	return o.UnifiedAccountingAttachmentOutput
}