// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/models/components"
)

type ListHrisEmployerBenefitsRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (o *ListHrisEmployerBenefitsRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListHrisEmployerBenefitsRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListHrisEmployerBenefitsRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListHrisEmployerBenefitsRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListHrisEmployerBenefitsResponseBody struct {
	PrevCursor *string                                       `json:"prev_cursor"`
	NextCursor *string                                       `json:"next_cursor"`
	Data       []components.UnifiedHrisEmployerbenefitOutput `json:"data"`
}

func (o *ListHrisEmployerBenefitsResponseBody) GetPrevCursor() *string {
	if o == nil {
		return nil
	}
	return o.PrevCursor
}

func (o *ListHrisEmployerBenefitsResponseBody) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *ListHrisEmployerBenefitsResponseBody) GetData() []components.UnifiedHrisEmployerbenefitOutput {
	if o == nil {
		return []components.UnifiedHrisEmployerbenefitOutput{}
	}
	return o.Data
}

type ListHrisEmployerBenefitsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListHrisEmployerBenefitsResponseBody

	Next func() (*ListHrisEmployerBenefitsResponse, error)
}

func (o *ListHrisEmployerBenefitsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListHrisEmployerBenefitsResponse) GetObject() *ListHrisEmployerBenefitsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
