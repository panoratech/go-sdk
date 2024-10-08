// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type CreateAccountingAccountRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original Accounting software.
	RemoteData                    *bool                                    `queryParam:"style=form,explode=true,name=remote_data"`
	UnifiedAccountingAccountInput components.UnifiedAccountingAccountInput `request:"mediaType=application/json"`
}

func (o *CreateAccountingAccountRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *CreateAccountingAccountRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *CreateAccountingAccountRequest) GetUnifiedAccountingAccountInput() components.UnifiedAccountingAccountInput {
	if o == nil {
		return components.UnifiedAccountingAccountInput{}
	}
	return o.UnifiedAccountingAccountInput
}

type CreateAccountingAccountResponse struct {
	HTTPMeta                       components.HTTPMetadata `json:"-"`
	UnifiedAccountingAccountOutput *components.UnifiedAccountingAccountOutput
}

func (o *CreateAccountingAccountResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateAccountingAccountResponse) GetUnifiedAccountingAccountOutput() *components.UnifiedAccountingAccountOutput {
	if o == nil {
		return nil
	}
	return o.UnifiedAccountingAccountOutput
}
