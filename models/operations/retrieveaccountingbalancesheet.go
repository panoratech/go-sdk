// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type RetrieveAccountingBalanceSheetRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// id of the balancesheet you want to retrieve.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Set to true to include data from the original Accounting software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
}

func (o *RetrieveAccountingBalanceSheetRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *RetrieveAccountingBalanceSheetRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RetrieveAccountingBalanceSheetRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

type RetrieveAccountingBalanceSheetResponse struct {
	HTTPMeta                            components.HTTPMetadata `json:"-"`
	UnifiedAccountingBalancesheetOutput *components.UnifiedAccountingBalancesheetOutput
}

func (o *RetrieveAccountingBalanceSheetResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RetrieveAccountingBalanceSheetResponse) GetUnifiedAccountingBalancesheetOutput() *components.UnifiedAccountingBalancesheetOutput {
	if o == nil {
		return nil
	}
	return o.UnifiedAccountingBalancesheetOutput
}
