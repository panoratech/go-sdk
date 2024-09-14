// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type CreateHrisTimesheetentryRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original Hris software.
	RemoteData                     *bool                                     `queryParam:"style=form,explode=true,name=remote_data"`
	UnifiedHrisTimesheetEntryInput components.UnifiedHrisTimesheetEntryInput `request:"mediaType=application/json"`
}

func (o *CreateHrisTimesheetentryRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *CreateHrisTimesheetentryRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *CreateHrisTimesheetentryRequest) GetUnifiedHrisTimesheetEntryInput() components.UnifiedHrisTimesheetEntryInput {
	if o == nil {
		return components.UnifiedHrisTimesheetEntryInput{}
	}
	return o.UnifiedHrisTimesheetEntryInput
}

type CreateHrisTimesheetentryResponse struct {
	HTTPMeta                        components.HTTPMetadata `json:"-"`
	UnifiedHrisTimesheetEntryOutput *components.UnifiedHrisTimesheetEntryOutput
}

func (o *CreateHrisTimesheetentryResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateHrisTimesheetentryResponse) GetUnifiedHrisTimesheetEntryOutput() *components.UnifiedHrisTimesheetEntryOutput {
	if o == nil {
		return nil
	}
	return o.UnifiedHrisTimesheetEntryOutput
}