// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"github.com/panoratech/go-sdk/models/components"
)

type ListHrisEmployeePayrollRunRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `default:"50" queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (l ListHrisEmployeePayrollRunRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListHrisEmployeePayrollRunRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListHrisEmployeePayrollRunRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListHrisEmployeePayrollRunRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListHrisEmployeePayrollRunRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListHrisEmployeePayrollRunRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListHrisEmployeePayrollRunResponseBody struct {
	PrevCursor *string                                          `json:"prev_cursor"`
	NextCursor *string                                          `json:"next_cursor"`
	Data       []components.UnifiedHrisEmployeepayrollrunOutput `json:"data"`
}

func (o *ListHrisEmployeePayrollRunResponseBody) GetPrevCursor() *string {
	if o == nil {
		return nil
	}
	return o.PrevCursor
}

func (o *ListHrisEmployeePayrollRunResponseBody) GetNextCursor() *string {
	if o == nil {
		return nil
	}
	return o.NextCursor
}

func (o *ListHrisEmployeePayrollRunResponseBody) GetData() []components.UnifiedHrisEmployeepayrollrunOutput {
	if o == nil {
		return []components.UnifiedHrisEmployeepayrollrunOutput{}
	}
	return o.Data
}

type ListHrisEmployeePayrollRunResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListHrisEmployeePayrollRunResponseBody

	Next func() (*ListHrisEmployeePayrollRunResponse, error)
}

func (o *ListHrisEmployeePayrollRunResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListHrisEmployeePayrollRunResponse) GetObject() *ListHrisEmployeePayrollRunResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
