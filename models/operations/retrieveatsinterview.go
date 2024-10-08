// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type RetrieveAtsInterviewRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// id of the interview you want to retrieve.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Set to true to include data from the original Ats software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
}

func (o *RetrieveAtsInterviewRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *RetrieveAtsInterviewRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RetrieveAtsInterviewRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

type RetrieveAtsInterviewResponse struct {
	HTTPMeta                  components.HTTPMetadata `json:"-"`
	UnifiedAtsInterviewOutput *components.UnifiedAtsInterviewOutput
}

func (o *RetrieveAtsInterviewResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RetrieveAtsInterviewResponse) GetUnifiedAtsInterviewOutput() *components.UnifiedAtsInterviewOutput {
	if o == nil {
		return nil
	}
	return o.UnifiedAtsInterviewOutput
}
