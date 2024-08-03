// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"github.com/panoratech/go-sdk/models/components"
)

type ListAccountingTrackingCategorysRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `default:"50" queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (l ListAccountingTrackingCategorysRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAccountingTrackingCategorysRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListAccountingTrackingCategorysRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListAccountingTrackingCategorysRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListAccountingTrackingCategorysRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListAccountingTrackingCategorysRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListAccountingTrackingCategorysResponseBody struct {
	PrevCursor *string                                              `json:"prev_cursor"`
	NextCursor *string                                              `json:"next_cursor"`
	Data       []components.UnifiedAccountingTrackingcategoryOutput `json:"data"`
}

func (o *ListAccountingTrackingCategorysResponseBody) GetPrevCursor() *string {
	if o == nil {
		return nil
	}
	return o.PrevCursor
}

func (o *ListAccountingTrackingCategorysResponseBody) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *ListAccountingTrackingCategorysResponseBody) GetData() []components.UnifiedAccountingTrackingcategoryOutput {
	if o == nil {
		return []components.UnifiedAccountingTrackingcategoryOutput{}
	}
	return o.Data
}

type ListAccountingTrackingCategorysResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListAccountingTrackingCategorysResponseBody
}

func (o *ListAccountingTrackingCategorysResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListAccountingTrackingCategorysResponse) GetObject() *ListAccountingTrackingCategorysResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
