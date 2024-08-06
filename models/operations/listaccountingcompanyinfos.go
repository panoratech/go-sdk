// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type ListAccountingCompanyInfosRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (o *ListAccountingCompanyInfosRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListAccountingCompanyInfosRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListAccountingCompanyInfosRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListAccountingCompanyInfosRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListAccountingCompanyInfosResponseBody struct {
	PrevCursor *string                                         `json:"prev_cursor"`
	NextCursor *string                                         `json:"next_cursor"`
	Data       []components.UnifiedAccountingCompanyinfoOutput `json:"data"`
}

func (o *ListAccountingCompanyInfosResponseBody) GetPrevCursor() *string {
	if o == nil {
		return nil
	}
	return o.PrevCursor
}

func (o *ListAccountingCompanyInfosResponseBody) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *ListAccountingCompanyInfosResponseBody) GetData() []components.UnifiedAccountingCompanyinfoOutput {
	if o == nil {
		return []components.UnifiedAccountingCompanyinfoOutput{}
	}
	return o.Data
}

type ListAccountingCompanyInfosResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListAccountingCompanyInfosResponseBody

	Next func() (*ListAccountingCompanyInfosResponse, error)
}

func (o *ListAccountingCompanyInfosResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListAccountingCompanyInfosResponse) GetObject() *ListAccountingCompanyInfosResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
