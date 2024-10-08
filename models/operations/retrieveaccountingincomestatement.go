// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type RetrieveAccountingIncomeStatementRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// id of the incomestatement you want to retrieve.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Set to true to include data from the original Accounting software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
}

func (o *RetrieveAccountingIncomeStatementRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *RetrieveAccountingIncomeStatementRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RetrieveAccountingIncomeStatementRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

type RetrieveAccountingIncomeStatementResponse struct {
	HTTPMeta                               components.HTTPMetadata `json:"-"`
	UnifiedAccountingIncomestatementOutput *components.UnifiedAccountingIncomestatementOutput
}

func (o *RetrieveAccountingIncomeStatementResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RetrieveAccountingIncomeStatementResponse) GetUnifiedAccountingIncomestatementOutput() *components.UnifiedAccountingIncomestatementOutput {
	if o == nil {
		return nil
	}
	return o.UnifiedAccountingIncomestatementOutput
}
