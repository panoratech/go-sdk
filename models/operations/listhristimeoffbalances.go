// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"github.com/panoratech/go-sdk/models/components"
)

type ListHrisTimeoffbalancesRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `default:"50" queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (l ListHrisTimeoffbalancesRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListHrisTimeoffbalancesRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListHrisTimeoffbalancesRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListHrisTimeoffbalancesRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListHrisTimeoffbalancesRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListHrisTimeoffbalancesRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListHrisTimeoffbalancesResponseBody struct {
	PrevCursor *string                                      `json:"prev_cursor"`
	NextCursor *string                                      `json:"next_cursor"`
	Data       []components.UnifiedHrisTimeoffbalanceOutput `json:"data"`
}

func (o *ListHrisTimeoffbalancesResponseBody) GetPrevCursor() *string {
	if o == nil {
		return nil
	}
	return o.PrevCursor
}

func (o *ListHrisTimeoffbalancesResponseBody) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *ListHrisTimeoffbalancesResponseBody) GetData() []components.UnifiedHrisTimeoffbalanceOutput {
	if o == nil {
		return []components.UnifiedHrisTimeoffbalanceOutput{}
	}
	return o.Data
}

type ListHrisTimeoffbalancesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListHrisTimeoffbalancesResponseBody

	Next func() (*ListHrisTimeoffbalancesResponse, error)
}

func (o *ListHrisTimeoffbalancesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListHrisTimeoffbalancesResponse) GetObject() *ListHrisTimeoffbalancesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
