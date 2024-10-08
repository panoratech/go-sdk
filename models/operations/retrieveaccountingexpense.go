// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type RetrieveAccountingExpenseRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// id of the expense you want to retrieve.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Set to true to include data from the original Accounting software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
}

func (o *RetrieveAccountingExpenseRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *RetrieveAccountingExpenseRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RetrieveAccountingExpenseRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

type RetrieveAccountingExpenseResponse struct {
	HTTPMeta                       components.HTTPMetadata `json:"-"`
	UnifiedAccountingExpenseOutput *components.UnifiedAccountingExpenseOutput
}

func (o *RetrieveAccountingExpenseResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RetrieveAccountingExpenseResponse) GetUnifiedAccountingExpenseOutput() *components.UnifiedAccountingExpenseOutput {
	if o == nil {
		return nil
	}
	return o.UnifiedAccountingExpenseOutput
}
