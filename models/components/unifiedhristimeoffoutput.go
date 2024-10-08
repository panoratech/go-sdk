// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"time"
)

// UnifiedHrisTimeoffOutputFieldMappings - The custom field mappings of the object between the remote 3rd party & Panora
type UnifiedHrisTimeoffOutputFieldMappings struct {
}

// UnifiedHrisTimeoffOutputRemoteData - The remote data of the time off in the context of the 3rd Party
type UnifiedHrisTimeoffOutputRemoteData struct {
}

type UnifiedHrisTimeoffOutput struct {
	// The UUID of the employee taking time off
	Employee *string `json:"employee,omitempty"`
	// The UUID of the approver for the time off request
	Approver *string `json:"approver,omitempty"`
	// The status of the time off request
	Status *string `json:"status,omitempty"`
	// A note from the employee about the time off request
	EmployeeNote *string `json:"employee_note,omitempty"`
	// The units used for the time off (e.g., Days, Hours)
	Units *string `json:"units,omitempty"`
	// The amount of time off requested
	Amount *float64 `json:"amount,omitempty"`
	// The type of time off request
	RequestType *string `json:"request_type,omitempty"`
	// The start time of the time off
	StartTime *time.Time `json:"start_time,omitempty"`
	// The end time of the time off
	EndTime *time.Time `json:"end_time,omitempty"`
	// The custom field mappings of the object between the remote 3rd party & Panora
	FieldMappings *UnifiedHrisTimeoffOutputFieldMappings `json:"field_mappings,omitempty"`
	// The UUID of the time off record
	ID *string `json:"id,omitempty"`
	// The remote ID of the time off in the context of the 3rd Party
	RemoteID *string `json:"remote_id,omitempty"`
	// The remote data of the time off in the context of the 3rd Party
	RemoteData *UnifiedHrisTimeoffOutputRemoteData `json:"remote_data,omitempty"`
	// The date when the time off was created in the 3rd party system
	RemoteCreatedAt *time.Time `json:"remote_created_at,omitempty"`
	// The created date of the time off record
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The last modified date of the time off record
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// Indicates if the time off was deleted in the remote system
	RemoteWasDeleted *bool `json:"remote_was_deleted,omitempty"`
}

func (u UnifiedHrisTimeoffOutput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UnifiedHrisTimeoffOutput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UnifiedHrisTimeoffOutput) GetEmployee() *string {
	if o == nil {
		return nil
	}
	return o.Employee
}

func (o *UnifiedHrisTimeoffOutput) GetApprover() *string {
	if o == nil {
		return nil
	}
	return o.Approver
}

func (o *UnifiedHrisTimeoffOutput) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *UnifiedHrisTimeoffOutput) GetEmployeeNote() *string {
	if o == nil {
		return nil
	}
	return o.EmployeeNote
}

func (o *UnifiedHrisTimeoffOutput) GetUnits() *string {
	if o == nil {
		return nil
	}
	return o.Units
}

func (o *UnifiedHrisTimeoffOutput) GetAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *UnifiedHrisTimeoffOutput) GetRequestType() *string {
	if o == nil {
		return nil
	}
	return o.RequestType
}

func (o *UnifiedHrisTimeoffOutput) GetStartTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartTime
}

func (o *UnifiedHrisTimeoffOutput) GetEndTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndTime
}

func (o *UnifiedHrisTimeoffOutput) GetFieldMappings() *UnifiedHrisTimeoffOutputFieldMappings {
	if o == nil {
		return nil
	}
	return o.FieldMappings
}

func (o *UnifiedHrisTimeoffOutput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UnifiedHrisTimeoffOutput) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *UnifiedHrisTimeoffOutput) GetRemoteData() *UnifiedHrisTimeoffOutputRemoteData {
	if o == nil {
		return nil
	}
	return o.RemoteData
}

func (o *UnifiedHrisTimeoffOutput) GetRemoteCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.RemoteCreatedAt
}

func (o *UnifiedHrisTimeoffOutput) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *UnifiedHrisTimeoffOutput) GetModifiedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedAt
}

func (o *UnifiedHrisTimeoffOutput) GetRemoteWasDeleted() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteWasDeleted
}
