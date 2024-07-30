// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type RetrieveCrmNoteRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// id of the note you want to retrieve.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Set to true to include data from the original Crm software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
}

func (o *RetrieveCrmNoteRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *RetrieveCrmNoteRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RetrieveCrmNoteRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

type RetrieveCrmNoteResponse struct {
	HTTPMeta             components.HTTPMetadata `json:"-"`
	UnifiedCrmNoteOutput *components.UnifiedCrmNoteOutput
}

func (o *RetrieveCrmNoteResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RetrieveCrmNoteResponse) GetUnifiedCrmNoteOutput() *components.UnifiedCrmNoteOutput {
	if o == nil {
		return nil
	}
	return o.UnifiedCrmNoteOutput
}