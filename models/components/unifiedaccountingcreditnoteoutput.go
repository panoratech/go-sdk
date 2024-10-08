// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"time"
)

// UnifiedAccountingCreditnoteOutputFieldMappings - The custom field mappings of the object between the remote 3rd party & Panora
type UnifiedAccountingCreditnoteOutputFieldMappings struct {
}

// UnifiedAccountingCreditnoteOutputRemoteData - The remote data of the credit note in the context of the 3rd Party
type UnifiedAccountingCreditnoteOutputRemoteData struct {
}

type UnifiedAccountingCreditnoteOutput struct {
	// The date of the credit note transaction
	TransactionDate *time.Time `json:"transaction_date,omitempty"`
	// The status of the credit note
	Status *string `json:"status,omitempty"`
	// The number of the credit note
	Number *string `json:"number,omitempty"`
	// The UUID of the associated contact
	ContactID *string `json:"contact_id,omitempty"`
	// The UUID of the associated company
	CompanyID *string `json:"company_id,omitempty"`
	// The exchange rate applied to the credit note
	ExchangeRate *string `json:"exchange_rate,omitempty"`
	// The total amount of the credit note
	TotalAmount *float64 `json:"total_amount,omitempty"`
	// The remaining credit on the credit note
	RemainingCredit *float64 `json:"remaining_credit,omitempty"`
	// The UUIDs of the tracking categories associated with the credit note
	TrackingCategories []string `json:"tracking_categories,omitempty"`
	// The currency of the credit note
	Currency *string `json:"currency,omitempty"`
	// The payments associated with the credit note
	Payments []string `json:"payments,omitempty"`
	// The applied payments associated with the credit note
	AppliedPayments []string `json:"applied_payments,omitempty"`
	// The UUID of the associated accounting period
	AccountingPeriodID *string `json:"accounting_period_id,omitempty"`
	// The custom field mappings of the object between the remote 3rd party & Panora
	FieldMappings *UnifiedAccountingCreditnoteOutputFieldMappings `json:"field_mappings,omitempty"`
	// The UUID of the credit note record
	ID *string `json:"id,omitempty"`
	// The remote ID of the credit note in the context of the 3rd Party
	RemoteID *string `json:"remote_id,omitempty"`
	// The remote data of the credit note in the context of the 3rd Party
	RemoteData *UnifiedAccountingCreditnoteOutputRemoteData `json:"remote_data,omitempty"`
	// The date when the credit note was created in the remote system
	RemoteCreatedAt *time.Time `json:"remote_created_at,omitempty"`
	// The date when the credit note was last updated in the remote system
	RemoteUpdatedAt *time.Time `json:"remote_updated_at,omitempty"`
	// The created date of the credit note record
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The last modified date of the credit note record
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
}

func (u UnifiedAccountingCreditnoteOutput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UnifiedAccountingCreditnoteOutput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UnifiedAccountingCreditnoteOutput) GetTransactionDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.TransactionDate
}

func (o *UnifiedAccountingCreditnoteOutput) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *UnifiedAccountingCreditnoteOutput) GetNumber() *string {
	if o == nil {
		return nil
	}
	return o.Number
}

func (o *UnifiedAccountingCreditnoteOutput) GetContactID() *string {
	if o == nil {
		return nil
	}
	return o.ContactID
}

func (o *UnifiedAccountingCreditnoteOutput) GetCompanyID() *string {
	if o == nil {
		return nil
	}
	return o.CompanyID
}

func (o *UnifiedAccountingCreditnoteOutput) GetExchangeRate() *string {
	if o == nil {
		return nil
	}
	return o.ExchangeRate
}

func (o *UnifiedAccountingCreditnoteOutput) GetTotalAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}

func (o *UnifiedAccountingCreditnoteOutput) GetRemainingCredit() *float64 {
	if o == nil {
		return nil
	}
	return o.RemainingCredit
}

func (o *UnifiedAccountingCreditnoteOutput) GetTrackingCategories() []string {
	if o == nil {
		return nil
	}
	return o.TrackingCategories
}

func (o *UnifiedAccountingCreditnoteOutput) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *UnifiedAccountingCreditnoteOutput) GetPayments() []string {
	if o == nil {
		return nil
	}
	return o.Payments
}

func (o *UnifiedAccountingCreditnoteOutput) GetAppliedPayments() []string {
	if o == nil {
		return nil
	}
	return o.AppliedPayments
}

func (o *UnifiedAccountingCreditnoteOutput) GetAccountingPeriodID() *string {
	if o == nil {
		return nil
	}
	return o.AccountingPeriodID
}

func (o *UnifiedAccountingCreditnoteOutput) GetFieldMappings() *UnifiedAccountingCreditnoteOutputFieldMappings {
	if o == nil {
		return nil
	}
	return o.FieldMappings
}

func (o *UnifiedAccountingCreditnoteOutput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UnifiedAccountingCreditnoteOutput) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *UnifiedAccountingCreditnoteOutput) GetRemoteData() *UnifiedAccountingCreditnoteOutputRemoteData {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *UnifiedAccountingCreditnoteOutput) GetRemoteCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.RemoteCreatedAt
}

func (o *UnifiedAccountingCreditnoteOutput) GetRemoteUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.RemoteUpdatedAt
}

func (o *UnifiedAccountingCreditnoteOutput) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *UnifiedAccountingCreditnoteOutput) GetModifiedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedAt
}
