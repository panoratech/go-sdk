// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type CreateFilestorageFolderRequest struct {
	// The connection token
	XConnectionToken              string                                   `header:"style=simple,explode=false,name=x-connection-token"`
	RemoteData                    bool                                     `queryParam:"style=form,explode=true,name=remote_data"`
	UnifiedFilestorageFolderInput components.UnifiedFilestorageFolderInput `request:"mediaType=application/json"`
}

func (o *CreateFilestorageFolderRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *CreateFilestorageFolderRequest) GetRemoteData() bool {
	if o == nil {
		return false
	}
	return o.RemoteData
}

func (o *CreateFilestorageFolderRequest) GetUnifiedFilestorageFolderInput() components.UnifiedFilestorageFolderInput {
	if o == nil {
		return components.UnifiedFilestorageFolderInput{}
	}
	return o.UnifiedFilestorageFolderInput
}

type CreateFilestorageFolderResponse struct {
	HTTPMeta                       components.HTTPMetadata `json:"-"`
	UnifiedFilestorageFolderOutput *components.UnifiedFilestorageFolderOutput
}

func (o *CreateFilestorageFolderResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateFilestorageFolderResponse) GetUnifiedFilestorageFolderOutput() *components.UnifiedFilestorageFolderOutput {
	if o == nil {
		return nil
	}
	return o.UnifiedFilestorageFolderOutput
}
