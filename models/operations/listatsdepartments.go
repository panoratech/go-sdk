// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"github.com/panoratech/go-sdk/models/components"
)

type ListAtsDepartmentsRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `default:"50" queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (l ListAtsDepartmentsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAtsDepartmentsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListAtsDepartmentsRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListAtsDepartmentsRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListAtsDepartmentsRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListAtsDepartmentsRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListAtsDepartmentsResponseBody struct {
	PrevCursor *string                                 `json:"prev_cursor"`
	NextCursor *string                                 `json:"next_cursor"`
	Data       []components.UnifiedAtsDepartmentOutput `json:"data"`
}

func (o *ListAtsDepartmentsResponseBody) GetPrevCursor() *string {
	if o == nil {
		return nil
	}
	return o.PrevCursor
}

func (o *ListAtsDepartmentsResponseBody) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *ListAtsDepartmentsResponseBody) GetData() []components.UnifiedAtsDepartmentOutput {
	if o == nil {
		return []components.UnifiedAtsDepartmentOutput{}
	}
	return o.Data
}

type ListAtsDepartmentsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListAtsDepartmentsResponseBody

	Next func() (*ListAtsDepartmentsResponse, error)
}

func (o *ListAtsDepartmentsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListAtsDepartmentsResponse) GetObject() *ListAtsDepartmentsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
