// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type RetrieveFilestorageFileRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// id of the file you want to retrieve.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Set to true to include data from the original File Storage software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
}

func (o *RetrieveFilestorageFileRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *RetrieveFilestorageFileRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RetrieveFilestorageFileRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

type RetrieveFilestorageFileResponse struct {
	HTTPMeta                     components.HTTPMetadata `json:"-"`
	UnifiedFilestorageFileOutput *components.UnifiedFilestorageFileOutput
}

func (o *RetrieveFilestorageFileResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RetrieveFilestorageFileResponse) GetUnifiedFilestorageFileOutput() *components.UnifiedFilestorageFileOutput {
	if o == nil {
		return nil
	}
	return o.UnifiedFilestorageFileOutput
}