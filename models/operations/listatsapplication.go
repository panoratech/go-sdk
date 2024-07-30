// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"github.com/panoratech/go-sdk/models/components"
)

type ListAtsApplicationRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `default:"50" queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (l ListAtsApplicationRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAtsApplicationRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListAtsApplicationRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListAtsApplicationRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListAtsApplicationRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListAtsApplicationRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListAtsApplicationResponseBody struct {
	PrevCursor string                                   `json:"prev_cursor"`
	NextCursor string                                   `json:"next_cursor"`
	Data       []components.UnifiedAtsApplicationOutput `json:"data"`
}

func (o *ListAtsApplicationResponseBody) GetPrevCursor() string {
	if o == nil {
		return ""
	}
	return o.PrevCursor
}

func (o *ListAtsApplicationResponseBody) GetNextCursor() string {
	if o == nil {
		return ""
	}
	return o.NextCursor
}

func (o *ListAtsApplicationResponseBody) GetData() []components.UnifiedAtsApplicationOutput {
	if o == nil {
		return []components.UnifiedAtsApplicationOutput{}
	}
	return o.Data
}

type ListAtsApplicationResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListAtsApplicationResponseBody
}

func (o *ListAtsApplicationResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListAtsApplicationResponse) GetObject() *ListAtsApplicationResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}