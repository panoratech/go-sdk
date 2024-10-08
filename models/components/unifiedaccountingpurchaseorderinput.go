// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"time"
)

// UnifiedAccountingPurchaseorderInputFieldMappings - The custom field mappings of the object between the remote 3rd party & Panora
type UnifiedAccountingPurchaseorderInputFieldMappings struct {
}

type UnifiedAccountingPurchaseorderInput struct {
	// The status of the purchase order
	Status *string `json:"status,omitempty"`
	// The issue date of the purchase order
	IssueDate *time.Time `json:"issue_date,omitempty"`
	// The purchase order number
	PurchaseOrderNumber *string `json:"purchase_order_number,omitempty"`
	// The delivery date for the purchase order
	DeliveryDate *time.Time `json:"delivery_date,omitempty"`
	// The UUID of the delivery address
	DeliveryAddress *string `json:"delivery_address,omitempty"`
	// The UUID of the customer
	Customer *string `json:"customer,omitempty"`
	// The UUID of the vendor
	Vendor *string `json:"vendor,omitempty"`
	// A memo or note for the purchase order
	Memo *string `json:"memo,omitempty"`
	// The UUID of the company
	CompanyID *string `json:"company_id,omitempty"`
	// The total amount of the purchase order in cents
	TotalAmount *float64 `json:"total_amount,omitempty"`
	// The currency of the purchase order
	Currency *string `json:"currency,omitempty"`
	// The exchange rate applied to the purchase order
	ExchangeRate *string `json:"exchange_rate,omitempty"`
	// The UUIDs of the tracking categories associated with the purchase order
	TrackingCategories []string `json:"tracking_categories,omitempty"`
	// The UUID of the associated accounting period
	AccountingPeriodID *string `json:"accounting_period_id,omitempty"`
	// The line items associated with this purchase order
	LineItems []LineItem `json:"line_items,omitempty"`
	// The custom field mappings of the object between the remote 3rd party & Panora
	FieldMappings *UnifiedAccountingPurchaseorderInputFieldMappings `json:"field_mappings,omitempty"`
}

func (u UnifiedAccountingPurchaseorderInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UnifiedAccountingPurchaseorderInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UnifiedAccountingPurchaseorderInput) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *UnifiedAccountingPurchaseorderInput) GetIssueDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.IssueDate
}

func (o *UnifiedAccountingPurchaseorderInput) GetPurchaseOrderNumber() *string {
	if o == nil {
		return nil
	}
	return o.PurchaseOrderNumber
}

func (o *UnifiedAccountingPurchaseorderInput) GetDeliveryDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeliveryDate
}

func (o *UnifiedAccountingPurchaseorderInput) GetDeliveryAddress() *string {
	if o == nil {
		return nil
	}
	return o.DeliveryAddress
}

func (o *UnifiedAccountingPurchaseorderInput) GetCustomer() *string {
	if o == nil {
		return nil
	}
	return o.Customer
}

func (o *UnifiedAccountingPurchaseorderInput) GetVendor() *string {
	if o == nil {
		return nil
	}
	return o.Vendor
}

func (o *UnifiedAccountingPurchaseorderInput) GetMemo() *string {
	if o == nil {
		return nil
	}
	return o.Memo
}

func (o *UnifiedAccountingPurchaseorderInput) GetCompanyID() *string {
	if o == nil {
		return nil
	}
	return o.CompanyID
}

func (o *UnifiedAccountingPurchaseorderInput) GetTotalAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}

func (o *UnifiedAccountingPurchaseorderInput) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *UnifiedAccountingPurchaseorderInput) GetExchangeRate() *string {
	if o == nil {
		return nil
	}
	return o.ExchangeRate
}

func (o *UnifiedAccountingPurchaseorderInput) GetTrackingCategories() []string {
	if o == nil {
		return nil
	}
	return o.TrackingCategories
}

func (o *UnifiedAccountingPurchaseorderInput) GetAccountingPeriodID() *string {
	if o == nil {
		return nil
	}
	return o.AccountingPeriodID
}

func (o *UnifiedAccountingPurchaseorderInput) GetLineItems() []LineItem {
	if o == nil {
		return nil
	}
	return o.LineItems
}

func (o *UnifiedAccountingPurchaseorderInput) GetFieldMappings() *UnifiedAccountingPurchaseorderInputFieldMappings {
	if o == nil {
		return nil
	}
	return o.FieldMappings
}
