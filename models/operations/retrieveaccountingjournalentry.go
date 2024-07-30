// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type RetrieveAccountingJournalEntryRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// id of the journalentry you want to retrieve.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Set to true to include data from the original Accounting software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
}

func (o *RetrieveAccountingJournalEntryRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *RetrieveAccountingJournalEntryRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *RetrieveAccountingJournalEntryRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

type RetrieveAccountingJournalEntryResponse struct {
	HTTPMeta                            components.HTTPMetadata `json:"-"`
	UnifiedAccountingJournalentryOutput *components.UnifiedAccountingJournalentryOutput
}

func (o *RetrieveAccountingJournalEntryResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *RetrieveAccountingJournalEntryResponse) GetUnifiedAccountingJournalentryOutput() *components.UnifiedAccountingJournalentryOutput {
	if o == nil {
		return nil
	}
	return o.UnifiedAccountingJournalentryOutput
}