// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type CreateFilestorageFileRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original Accounting software.
	RemoteData                  *bool                                  `queryParam:"style=form,explode=true,name=remote_data"`
	UnifiedFilestorageFileInput components.UnifiedFilestorageFileInput `request:"mediaType=application/json"`
}

func (o *CreateFilestorageFileRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *CreateFilestorageFileRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *CreateFilestorageFileRequest) GetUnifiedFilestorageFileInput() components.UnifiedFilestorageFileInput {
	if o == nil {
		return components.UnifiedFilestorageFileInput{}
	}
	return o.UnifiedFilestorageFileInput
}

type CreateFilestorageFileResponse struct {
	HTTPMeta                     components.HTTPMetadata `json:"-"`
	UnifiedFilestorageFileOutput *components.UnifiedFilestorageFileOutput
}

func (o *CreateFilestorageFileResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateFilestorageFileResponse) GetUnifiedFilestorageFileOutput() *components.UnifiedFilestorageFileOutput {
	if o == nil {
		return nil
	}
	return o.UnifiedFilestorageFileOutput
}
