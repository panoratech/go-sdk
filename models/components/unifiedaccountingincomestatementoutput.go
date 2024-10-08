// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"time"
)

// UnifiedAccountingIncomestatementOutputFieldMappings - The custom field mappings of the object between the remote 3rd party & Panora
type UnifiedAccountingIncomestatementOutputFieldMappings struct {
}

// UnifiedAccountingIncomestatementOutputRemoteData - The remote data of the income statement in the context of the 3rd Party
type UnifiedAccountingIncomestatementOutputRemoteData struct {
}

type UnifiedAccountingIncomestatementOutput struct {
	// The name of the income statement
	Name *string `json:"name,omitempty"`
	// The currency used in the income statement
	Currency *string `json:"currency,omitempty"`
	// The start date of the period covered by the income statement
	StartPeriod *time.Time `json:"start_period,omitempty"`
	// The end date of the period covered by the income statement
	EndPeriod *time.Time `json:"end_period,omitempty"`
	// The gross profit for the period
	GrossProfit *float64 `json:"gross_profit,omitempty"`
	// The net operating income for the period
	NetOperatingIncome *float64 `json:"net_operating_income,omitempty"`
	// The net income for the period
	NetIncome *float64 `json:"net_income,omitempty"`
	// The custom field mappings of the object between the remote 3rd party & Panora
	FieldMappings *UnifiedAccountingIncomestatementOutputFieldMappings `json:"field_mappings,omitempty"`
	// The UUID of the income statement record
	ID *string `json:"id,omitempty"`
	// The remote ID of the income statement in the context of the 3rd Party
	RemoteID *string `json:"remote_id,omitempty"`
	// The remote data of the income statement in the context of the 3rd Party
	RemoteData *UnifiedAccountingIncomestatementOutputRemoteData `json:"remote_data,omitempty"`
	// The created date of the income statement record
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The last modified date of the income statement record
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
}

func (u UnifiedAccountingIncomestatementOutput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UnifiedAccountingIncomestatementOutput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UnifiedAccountingIncomestatementOutput) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UnifiedAccountingIncomestatementOutput) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *UnifiedAccountingIncomestatementOutput) GetStartPeriod() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartPeriod
}

func (o *UnifiedAccountingIncomestatementOutput) GetEndPeriod() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndPeriod
}

func (o *UnifiedAccountingIncomestatementOutput) GetGrossProfit() *float64 {
	if o == nil {
		return nil
	}
	return o.GrossProfit
}

func (o *UnifiedAccountingIncomestatementOutput) GetNetOperatingIncome() *float64 {
	if o == nil {
		return nil
	}
	return o.NetOperatingIncome
}

func (o *UnifiedAccountingIncomestatementOutput) GetNetIncome() *float64 {
	if o == nil {
		return nil
	}
	return o.NetIncome
}

func (o *UnifiedAccountingIncomestatementOutput) GetFieldMappings() *UnifiedAccountingIncomestatementOutputFieldMappings {
	if o == nil {
		return nil
	}
	return o.FieldMappings
}

func (o *UnifiedAccountingIncomestatementOutput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UnifiedAccountingIncomestatementOutput) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *UnifiedAccountingIncomestatementOutput) GetRemoteData() *UnifiedAccountingIncomestatementOutputRemoteData {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *UnifiedAccountingIncomestatementOutput) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *UnifiedAccountingIncomestatementOutput) GetModifiedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedAt
}
