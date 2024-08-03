// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"time"
)

type UnifiedCrmTaskInput struct {
	// The subject of the task
	Subject *string `json:"subject"`
	// The content of the task
	Content *string `json:"content"`
	// The status of the task. Authorized values are PENDING, COMPLETED.
	Status *string `json:"status"`
	// The due date of the task
	DueDate *time.Time `json:"due_date,omitempty"`
	// The finished date of the task
	FinishedDate *time.Time `json:"finished_date,omitempty"`
	// The UUID of the user tied to the task
	UserID *string `json:"user_id,omitempty"`
	// The UUID of the company tied to the task
	CompanyID *string `json:"company_id,omitempty"`
	// The UUID of the deal tied to the task
	DealID *string `json:"deal_id,omitempty"`
	// The custom field mappings of the task between the remote 3rd party & Panora
	FieldMappings map[string]any `json:"field_mappings,omitempty"`
}

func (u UnifiedCrmTaskInput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UnifiedCrmTaskInput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UnifiedCrmTaskInput) GetSubject() *string {
	if o == nil {
		return nil
	}
	return o.Subject
}

func (o *UnifiedCrmTaskInput) GetContent() *string {
	if o == nil {
		return nil
	}
	return o.Content
}

func (o *UnifiedCrmTaskInput) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *UnifiedCrmTaskInput) GetDueDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.DueDate
}

func (o *UnifiedCrmTaskInput) GetFinishedDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.FinishedDate
}

func (o *UnifiedCrmTaskInput) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

func (o *UnifiedCrmTaskInput) GetCompanyID() *string {
	if o == nil {
		return nil
	}
	return o.CompanyID
}

func (o *UnifiedCrmTaskInput) GetDealID() *string {
	if o == nil {
		return nil
	}
	return o.DealID
}

func (o *UnifiedCrmTaskInput) GetFieldMappings() map[string]any {
	if o == nil {
		return nil
	}
	return o.FieldMappings
}
