// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"github.com/panoratech/go-sdk/models/components"
)

type ListAccountingIncomeStatementRequest struct {
	// The connection token
	XConnectionToken string `header:"style=simple,explode=false,name=x-connection-token"`
	// Set to true to include data from the original software.
	RemoteData *bool `queryParam:"style=form,explode=true,name=remote_data"`
	// Set to get the number of records.
	Limit *float64 `default:"50" queryParam:"style=form,explode=true,name=limit"`
	// Set to get the number of records after this cursor.
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (l ListAccountingIncomeStatementRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAccountingIncomeStatementRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListAccountingIncomeStatementRequest) GetXConnectionToken() string {
	if o == nil {
		return ""
	}
	return o.XConnectionToken
}

func (o *ListAccountingIncomeStatementRequest) GetRemoteData() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *ListAccountingIncomeStatementRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListAccountingIncomeStatementRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type ListAccountingIncomeStatementResponseBody struct {
	PrevCursor string                                              `json:"prev_cursor"`
	NextCursor string                                              `json:"next_cursor"`
	Data       []components.UnifiedAccountingIncomestatementOutput `json:"data"`
}

func (o *ListAccountingIncomeStatementResponseBody) GetPrevCursor() string {
	if o == nil {
		return ""
	}
	return o.PrevCursor
}

func (o *ListAccountingIncomeStatementResponseBody) GetNextCursor() string {
	if o == nil {
		return ""
	}
	return o.NextCursor
}

func (o *ListAccountingIncomeStatementResponseBody) GetData() []components.UnifiedAccountingIncomestatementOutput {
	if o == nil {
		return []components.UnifiedAccountingIncomestatementOutput{}
	}
	return o.Data
}

type ListAccountingIncomeStatementResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *ListAccountingIncomeStatementResponseBody
}

func (o *ListAccountingIncomeStatementResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListAccountingIncomeStatementResponse) GetObject() *ListAccountingIncomeStatementResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
