// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type CreateHrisEmployeeRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original Hris software.
	RemoteData               *bool                               `queryParam:"style=form,explode=true,name=remote_data"`
	UnifiedHrisEmployeeInput components.UnifiedHrisEmployeeInput `request:"mediaType=application/json"`
}

func (o *CreateHrisEmployeeRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *CreateHrisEmployeeRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *CreateHrisEmployeeRequest) GetUnifiedHrisEmployeeInput() components.UnifiedHrisEmployeeInput {
	if o == nil {
		return components.UnifiedHrisEmployeeInput{}
	}
	return o.UnifiedHrisEmployeeInput
}

type CreateHrisEmployeeResponse struct {
	HTTPMeta                  components.HTTPMetadata `json:"-"`
	UnifiedHrisEmployeeOutput *components.UnifiedHrisEmployeeOutput
}

func (o *CreateHrisEmployeeResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateHrisEmployeeResponse) GetUnifiedHrisEmployeeOutput() *components.UnifiedHrisEmployeeOutput {
	if o == nil {
		return nil
	}
	return o.UnifiedHrisEmployeeOutput
}