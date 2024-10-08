// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/panoratech/go-sdk/internal/utils"
	"time"
)

// UnifiedHrisTimesheetEntryOutputFieldMappings - The custom field mappings of the object between the remote 3rd party & Panora
type UnifiedHrisTimesheetEntryOutputFieldMappings struct {
}

// UnifiedHrisTimesheetEntryOutputRemoteData - The remote data of the timesheet entry in the context of the 3rd Party
type UnifiedHrisTimesheetEntryOutputRemoteData struct {
}

type UnifiedHrisTimesheetEntryOutput struct {
	// The number of hours worked
	HoursWorked *float64 `json:"hours_worked,omitempty"`
	// The start time of the timesheet entry
	StartTime *time.Time `json:"start_time,omitempty"`
	// The end time of the timesheet entry
	EndTime *time.Time `json:"end_time,omitempty"`
	// The UUID of the associated employee
	EmployeeID *string `json:"employee_id,omitempty"`
	// Indicates if the timesheet entry was deleted in the remote system
	RemoteWasDeleted *bool `json:"remote_was_deleted,omitempty"`
	// The custom field mappings of the object between the remote 3rd party & Panora
	FieldMappings *UnifiedHrisTimesheetEntryOutputFieldMappings `json:"field_mappings,omitempty"`
	// The UUID of the timesheet entry record
	ID *string `json:"id,omitempty"`
	// The remote ID of the timesheet entry
	RemoteID *string `json:"remote_id,omitempty"`
	// The date when the timesheet entry was created in the remote system
	RemoteCreatedAt *time.Time `json:"remote_created_at,omitempty"`
	// The created date of the timesheet entry
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The last modified date of the timesheet entry
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// The remote data of the timesheet entry in the context of the 3rd Party
	RemoteData *UnifiedHrisTimesheetEntryOutputRemoteData `json:"remote_data,omitempty"`
}

func (u UnifiedHrisTimesheetEntryOutput) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UnifiedHrisTimesheetEntryOutput) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UnifiedHrisTimesheetEntryOutput) GetHoursWorked() *float64 {
	if o == nil {
		return nil
	}
	return o.HoursWorked
}

func (o *UnifiedHrisTimesheetEntryOutput) GetStartTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartTime
}

func (o *UnifiedHrisTimesheetEntryOutput) GetEndTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndTime
}

func (o *UnifiedHrisTimesheetEntryOutput) GetEmployeeID() *string {
	if o == nil {
		return nil
	}
	return o.EmployeeID
}

func (o *UnifiedHrisTimesheetEntryOutput) GetRemoteWasDeleted() *bool {
	if o == nil {
		return nil
	}
	return o.RemoteWasDeleted
}

func (o *UnifiedHrisTimesheetEntryOutput) GetFieldMappings() *UnifiedHrisTimesheetEntryOutputFieldMappings {
	if o == nil {
		return nil
	}
	return o.FieldMappings
}

func (o *UnifiedHrisTimesheetEntryOutput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UnifiedHrisTimesheetEntryOutput) GetRemoteID() *string {
	if o == nil {
		return nil
	}
	return o.RemoteID
}

func (o *UnifiedHrisTimesheetEntryOutput) GetRemoteCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.RemoteCreatedAt
}

func (o *UnifiedHrisTimesheetEntryOutput) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *UnifiedHrisTimesheetEntryOutput) GetModifiedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ModifiedAt
}

func (o *UnifiedHrisTimesheetEntryOutput) GetRemoteData() *UnifiedHrisTimesheetEntryOutputRemoteData {
	if o == nil {
		return nil
	}
	return o.RemoteData
}
